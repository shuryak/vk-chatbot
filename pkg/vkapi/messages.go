package vkapi

func (vkapi *VKAPI) MessagesSend(params Params) (response int, err error) {
	reqParams := Params{
		"user_ids": "",
		"peer_ids": "",
	}

	err = vkapi.RequestUnmarshal("messages.send", &response, params, reqParams)

	return
}
