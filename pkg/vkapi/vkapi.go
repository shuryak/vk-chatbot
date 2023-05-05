package vkapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/doc"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/transport"
	"io"
	"mime"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
)

type VKAPI struct {
	accessTokens []string
	lastToken    uint32
	Version      string
	MethodURL    string
	Client       *http.Client
	UserAgent    string
	Handler      func(method string, params ...Params) (Response, error)
	// TODO: limits

	mux sync.Mutex
}

func NewVKAPI(tokens ...string) *VKAPI {
	vk := VKAPI{
		accessTokens: tokens,
		Version:      doc.Version,
		MethodURL:    doc.MethodURL,
		Client:       http.DefaultClient,
		UserAgent:    transport.UserAgent,
	}
	vk.Handler = vk.DefaultHandler
	return &vk
}

type RawMessage []byte

func (m *RawMessage) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return *m, nil
}

func (m *RawMessage) UnmarshalJSON(data []byte) error {
	*m = append((*m)[0:0], data...)
	return nil
}

type Response struct {
	Response      RawMessage             `json:"response"`
	Error         objects.Error          `json:"error"`
	ExecuteErrors []objects.ExecuteError `json:"execute_errors"`
}

func (vkapi *VKAPI) getToken() string {
	i := atomic.AddUint32(&vkapi.lastToken, 1)
	return vkapi.accessTokens[(int(i)-1)%len(vkapi.accessTokens)] // TODO: explain
}

type Params map[string]interface{}

func (p Params) WithContext(ctx context.Context) Params {
	p[":context"] = ctx
	return p
}

func buildQuery(paramsSlice ...Params) (context.Context, url.Values) {
	query := url.Values{}
	ctx := context.Background()

	for _, params := range paramsSlice {
		for key, value := range params {
			switch key {
			case "access_token":
				continue
			case ":context":
				ctx = value.(context.Context)
			default:
				query.Set(key, transport.FmtValue(value, 0))
			}
		}
	}

	return ctx, query
}

func (vkapi *VKAPI) DefaultHandler(method string, params ...Params) (Response, error) {
	endpoint := vkapi.MethodURL + method
	ctx, query := buildQuery(params...)
	var response Response

	rawBody := bytes.NewBufferString(query.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, rawBody)
	if err != nil {
		return response, err
	}

	acceptEncdoing := "gzip"

	token := params[len(params)-1]["access_token"].(string)
	req.Header.Set("Authorization", "Bearer "+token)

	req.Header.Set("User-Agent", vkapi.UserAgent)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Accept-Encoding", acceptEncdoing)

	var reader io.Reader

	resp, err := vkapi.Client.Do(req)
	if err != nil {
		return response, err
	}

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		gzipReader, _ := gzip.NewReader(resp.Body)
		defer gzipReader.Close()

		reader = gzipReader
	default:
		reader = resp.Body
	}

	mediatype, _, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	switch mediatype {
	case "application/json":
		err = json.NewDecoder(reader).Decode(&response)
		if err != nil {
			_ = resp.Body.Close()
			return response, err
		}
	default:
		// TODO: handle
	}

	_ = resp.Body.Close()

	switch response.Error.Code {
	case objects.ErrorNoType:
		return response, nil
	case objects.ErrorTooManyRequests:
		// TODO: handle
	}

	return response, &response.Error
}

func (vkapi *VKAPI) Request(method string, params ...Params) ([]byte, error) {
	token := vkapi.getToken()

	reqParams := Params{
		"access_token": token,
		"v":            vkapi.Version,
	}

	params = append(params, reqParams)

	resp, err := vkapi.Handler(method, params...)

	return resp.Response, err
}

func (vkapi *VKAPI) RequestUnmarshal(method string, obj interface{}, params ...Params) error {
	rawResponse, err := vkapi.Request(method, params...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawResponse, &obj)

	return err
}
