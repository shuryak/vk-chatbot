package transport

import (
	"fmt"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/doc"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"reflect"
)

// https://vk.com/dev/callback_api?f=1.7.%20HTTP-заголовки
const (
	XRetryCounter = "X-Retry-Counter"
	RetryAfter    = "Retry-After"
)

const UserAgent = "InternBot/" + doc.Version

type RequestMetadata int

const (
	HTTPClientKey RequestMetadata = iota
	UserAgentKey
	GroupIDKey
	EventIDKey
	LongPollTsKey
	CallbackRetryCounterKey
	CallbackRetryAfterKey
	CallbackRemove
	EventVersionKey
)

func fmtReflectValue(value reflect.Value, depth int) string {
	switch f := value; value.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return fmtBool(f.Bool())
	case reflect.Array, reflect.Slice:
		s := ""

		for i := 0; i < f.Len(); i++ {
			if i > 0 {
				s += ","
			}

			s += FmtValue(f.Index(i).Interface(), depth)
		}

		return s
	case reflect.Ptr:
		// pointer to array or slice or struct? ok at top level
		// but not embedded (avoid loops)
		if depth == 0 && f.Pointer() != 0 {
			switch a := f.Elem(); a.Kind() {
			case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
				return FmtValue(a.Interface(), depth+1)
			}
		}
	}

	return fmt.Sprint(value)
}

func FmtValue(value interface{}, depth int) string {
	if value == nil {
		return ""
	}

	switch f := value.(type) {
	case bool:
		return fmtBool(f)
	case objects.Attachment:
		return f.ToAttachment()
	case objects.JSONObject:
		return f.ToJSON()
	case reflect.Value:
		return fmtReflectValue(f, depth)
	}

	return fmtReflectValue(reflect.ValueOf(value), depth)
}

func fmtBool(value bool) string {
	if value {
		return "1"
	}
	return "0"
}
