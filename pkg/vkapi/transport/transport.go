package transport

import "github.com/shuryak/vk-chatbot/pkg/vkapi"

// https://vk.com/dev/callback_api?f=1.7.%20HTTP-заголовки
const (
	XRetryCounter = "X-Retry-Counter"
	RetryAfter    = "Retry-After"
)

const UserAgent = "InternBot/" + vkapi.Version

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
