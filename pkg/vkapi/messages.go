package vkapi

import "github.com/shuryak/vk-chatbot/pkg/vkapi/objects"

type MessageSendPeerIDsResponse []struct {
	PeerID                int           `json:"peer_id"`
	MessageID             int           `json:"message_id"`
	ConversationMessageID int           `json:"conversation_message_id"`
	Error                 objects.Error `json:"error"`
}

func (vkapi *VKAPI) MessagesSend(params Params) (response MessageSendPeerIDsResponse, err error) {
	reqParams := Params{
		"user_ids": "",
		"peer_ids": "",
	}

	err = vkapi.RequestUnmarshal("messages.send", &response, params, reqParams)

	return
}
