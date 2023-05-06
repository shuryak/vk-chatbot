package params

import "github.com/shuryak/vk-chatbot/pkg/vkapi"

type UsersGetBuilder struct {
	vkapi.Params
}

func NewUsersGetBuilder() *UsersGetBuilder {
	return &UsersGetBuilder{vkapi.Params{}}
}

func (b *UsersGetBuilder) UserIDs(v []string) *UsersGetBuilder {
	b.Params["user_ids"] = v
	return b
}

func (b *UsersGetBuilder) Fields(v []string) *UsersGetBuilder {
	b.Params["fields"] = v
	return b
}
