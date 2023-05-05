package vkapi

import (
	"bytes"
	"encoding/json"
	"github.com/shuryak/vk-chatbot/pkg/vkapi/objects"
	"io"
	"mime/multipart"
)

type PhotosGetMessagesUploadServerResponse struct {
	AlbumID   int    `json:"album_id"`
	UploadURL string `json:"upload_url"`
	UserID    int    `json:"user_id,omitempty"`
	GroupID   int    `json:"group_id,omitempty"`
}

func (vkapi *VKAPI) PhotosGetMessagesUploadServer(params Params) (response PhotosGetMessagesUploadServerResponse, err error) {
	err = vkapi.RequestUnmarshal("photos.getMessagesUploadServer", &response, params)
	return
}

func (vkapi *VKAPI) UploadFile(url string, file io.Reader, fieldname, filename string) (bodyContent []byte, err error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(fieldname, filename)
	if err != nil {
		vkapi.l.Error("VKAPI - UploadFile - writer.CreateFormFile: %v", err)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		vkapi.l.Error("VKAPI - UploadFile - io.Copy: %v", err)
		return
	}

	contentType := writer.FormDataContentType()
	_ = writer.Close()

	resp, err := vkapi.Client.Post(url, contentType, body)
	if err != nil {
		vkapi.l.Error("VKAPI - UploadFile - vkapi.Client.Post: %v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			vkapi.l.Error("VKAPI - UploadFile - Body.Close: %v", err)
		}
	}(resp.Body)

	bodyContent, err = io.ReadAll(resp.Body)

	return
}

type PhotosSaveMessagesPhotoResponse []objects.Photo

type PhotosMessageUploadResponse struct {
	Hash   string `json:"hash"`
	Photo  string `json:"photo"`
	Server int    `json:"server"`
}

func (vkapi *VKAPI) UploadMessagesPhoto(peerID int, file io.Reader) (response PhotosSaveMessagesPhotoResponse, err error) {
	uploadServer, err := vkapi.PhotosGetMessagesUploadServer(Params{
		"peer_id": peerID,
	})
	if err != nil {
		vkapi.l.Error("VKAPI - UploadMessagesPhoto - vkapi.PhotosGetMessagesUploadServer: %v", err)
		return
	}

	bodyContent, err := vkapi.UploadFile(uploadServer.UploadURL, file, "photo", "photo.jpeg")
	if err != nil {
		vkapi.l.Error("VKAPI - UploadMessagesPhoto - vkapi.UploadFile: %v", err)
		return
	}

	var handler PhotosMessageUploadResponse

	err = json.Unmarshal(bodyContent, &handler)
	if err != nil {
		vkapi.l.Error("VKAPI - UploadMessagesPhoto - json.Unmarshal: %v", err)
		return
	}

	response, err = vkapi.PhotosSaveMessagesPhoto(Params{
		"server": handler.Server,
		"photo":  handler.Photo,
		"hash":   handler.Hash,
	})

	return
}

func (vkapi *VKAPI) PhotosSaveMessagesPhoto(params Params) (response PhotosSaveMessagesPhotoResponse, err error) {
	err = vkapi.RequestUnmarshal("photos.saveMessagesPhoto", &response, params)
	return
}
