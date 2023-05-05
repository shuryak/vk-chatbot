package params

import "github.com/shuryak/vk-chatbot/pkg/vkapi"

type MessagesSendBuilder struct {
	vkapi.Params
}

func NewMessagesSendBuilder() *MessagesSendBuilder {
	return &MessagesSendBuilder{vkapi.Params{}}
}

func (b *MessagesSendBuilder) PeerID(v int) *MessagesSendBuilder {
	b.Params["peer_id"] = v
	return b
}

func (b *MessagesSendBuilder) ChatID(v int) *MessagesSendBuilder {
	b.Params["chat_id"] = v
	return b
}

func (b *MessagesSendBuilder) PeerIDs(v []int) *MessagesSendBuilder {
	b.Params["peer_ids"] = v
	return b
}

func (b *MessagesSendBuilder) Message(v string) *MessagesSendBuilder {
	b.Params["message"] = v
	return b
}

func (b *MessagesSendBuilder) RandomID(v int) *MessagesSendBuilder {
	b.Params["random_id"] = v
	return b
}

func (b *MessagesSendBuilder) Keyboard(v interface{}) *MessagesSendBuilder {
	b.Params["keyboard"] = v
	return b
}

func (b *MessagesSendBuilder) Attachment(v string) *MessagesSendBuilder {
	b.Params["attachment"] = v
	return b
}
