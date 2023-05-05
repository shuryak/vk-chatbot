package callback

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shuryak/vk-chatbot/pkg/logger"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/events"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/transport"
	"net/http"
	"strconv"
	"time"
)

type Callback struct {
	ConfirmationKeys map[int]string // GroupID: ConfirmationKey
	SecretKeys       map[int]string // GroupID: SecretKey
	events.FuncList

	l logger.Interface
}

func NewCallback(l logger.Interface) *Callback {
	return &Callback{
		ConfirmationKeys: make(map[int]string),
		SecretKeys:       make(map[int]string),
		l:                l,
		FuncList:         *events.NewFuncList(l),
	}
}

func (cb *Callback) HandleFunc(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var e events.GroupEvent
	if err := decoder.Decode(&e); err != nil {
		cb.l.Error("Callback - HandleFunc - decoder.Decode: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	secretKey, ok := cb.SecretKeys[e.GroupID]

	if ok && e.Secret != secretKey {
		cb.l.Error("Callback - HandleFunc - secret key check")
		http.Error(w, "Bad Secret", http.StatusForbidden)
	}

	if e.Type == events.EventConfirmation {
		if _, ok := cb.ConfirmationKeys[e.GroupID]; ok {
			_, err := fmt.Fprintf(w, cb.ConfirmationKeys[e.GroupID])
			if err != nil {
				cb.l.Error("Callback - HandleFunc - confirmation key print: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}
		return
	}

	ctx := r.Context()

	// https://vk.com/dev/callback_api?f=1.7.%20HTTP-заголовки
	retryCounter, _ := strconv.Atoi(r.Header.Get(transport.XRetryCounter))
	ctx = context.WithValue(ctx, transport.CallbackRetryCounterKey, retryCounter)

	var (
		code   int
		date   time.Time
		remove bool
	)

	retryAfter := func(c int, d time.Time) {
		code = c
		date = d
	}
	ctx = context.WithValue(ctx, transport.CallbackRemove, retryAfter)

	// https://vk.com/dev/callback_api?f=1.2.%20Удаление сервера
	removeFunc := func() {
		remove = true
	}
	ctx = context.WithValue(ctx, transport.CallbackRemove, removeFunc)

	if err := cb.Handler(ctx, e); err != nil {
		cb.l.Error("Callback - HandleFunc - cb.Handler: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if remove {
		_, _ = w.Write([]byte("remove"))
		return
	}

	if code != 0 {
		w.Header().Set(transport.RetryAfter, date.Format(http.TimeFormat))
		http.Error(w, http.StatusText(code), code)

		return
	}

	_, _ = w.Write([]byte("ok"))
}
