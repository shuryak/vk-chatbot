package vkapi

import "github.com/shuryak/vk-chatbot/pkg/vkapi/objects"

type UsersGetResponse []objects.User

func (vkapi *VKAPI) UsersGet(params Params) (response UsersGetResponse, err error) {
	err = vkapi.RequestUnmarshal("users.get", &response, params)
	return
}
