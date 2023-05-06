package objects

import (
	"encoding/json"
	"fmt"
)

type ClientInfo struct {
	ButtonActions  []string `json:"button_actions"`
	Keyboard       bool     `json:"keyboard"`
	InlineKeyboard bool     `json:"inline_keyboard"`
	Carousel       bool     `json:"carousel"`
	LangId         int      `json:"lang_id"`
}

type Message struct {
	ID                    int                 `json:"id"`
	Date                  int                 `json:"date"`
	PeerID                int                 `json:"peer_id"`
	FromID                int                 `json:"from_id"`
	Text                  string              `json:"text"`
	RandomID              int                 `json:"random_id"`
	Attachments           []MessageAttachment `json:"attachments"`
	Important             bool                `json:"important"`
	Payload               string              `json:"payload"`
	Keyboard              MessagesKeyboard    `json:"keyboard"`
	FwdMessages           []Message           `json:"fwd_messages"`
	ReplyMessage          *Message            `json:"reply_message"`
	Action                MessageAction       `json:"action"`
	AdminAuthorID         int                 `json:"admin_author_id"`
	ConversationMessageID int                 `json:"conversation_message_id"`
	IsCropped             bool                `json:"is_cropped"`
	MembersCount          int                 `json:"members_count"`
	UpdateTime            int                 `json:"update_time"`
	WasListened           bool                `json:"was_listened"`
	PinnedAt              int                 `json:"pinned_at"`
	MessageTag            string              `json:"message_tag"`

	// TODO: full compliance with the docs (https://vk.com/dev/objects/message)
}

type MessageAction struct {
	Type     string             `json:"type"`
	MemberId int                `json:"member_id"`
	Text     string             `json:"text"`
	Email    string             `json:"email"`
	Photo    MessageActionPhoto `json:"photo"`
}

type MessageActionPhoto struct {
	Photo50  string `json:"photo_50"`
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`
}

type MessageAttachment struct {
	Photo   Photo   `json:"photo"`
	Sticker Sticker `json:"sticker"`
	// TODO: full compliance with the docs (https://vk.com/dev/objects/attachments_m)

	Type string `json:"type"`
}

type Photo struct {
	ID      int    `json:"id"` // Photo ID
	AlbumID int    `json:"album_id"`
	OwnerID int    `json:"owner_id"`
	UserId  int    `json:"user_id"`
	Text    string `json:"text"`
	Date    int    `json:"date"`
	// TODO: Sizes (https://vk.com/dev/objects/photo)

	Width  int `json:"width"`
	Height int `json:"height"`
}

func (photo Photo) ToAttachment() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

type Sticker struct {
	ProductID int `json:"product_id"`
	StickerID int `json:"sticker_id"`
	// TODO: Images, ImagesWithBackground (https://vk.com/dev/objects/sticker)

	AnimationURL string `json:"animation-url"`
	IsAllowed    bool   `json:"is_allowed"`
}

type MessagesKeyboard struct {
	OneTime bool                      `json:"one_time,omitempty"`
	Buttons [][]MessageKeyboardButton `json:"buttons"`
	Inline  bool                      `json:"inline,omitempty"`
}

func NewMessagesKeyboard(oneTime bool) *MessagesKeyboard {
	return &MessagesKeyboard{
		OneTime: oneTime,
		Buttons: [][]MessageKeyboardButton{},
	}
}

func (mg *MessagesKeyboard) ToJSON() string {
	b, _ := json.Marshal(mg)
	return string(b)
}

func NewMessagesKeyboardInline() *MessagesKeyboard {
	return &MessagesKeyboard{
		Buttons: [][]MessageKeyboardButton{},
		Inline:  true,
	}
}

func (mg *MessagesKeyboard) AddRow() *MessagesKeyboard {
	if len(mg.Buttons) == 0 {
		mg.Buttons = make([][]MessageKeyboardButton, 1)
	} else {
		row := make([]MessageKeyboardButton, 0)
		mg.Buttons = append(mg.Buttons, row)
	}

	return mg
}

func (mg *MessagesKeyboard) AddTextButton(label string, payload interface{}, color string) *MessagesKeyboard {
	b, err := json.Marshal(payload)
	if err != nil {
		panic(err) // TODO: properly handle error
	}

	button := MessageKeyboardButton{
		Action: MessageKeyboardButtonAction{
			Type:    ButtonText,
			Label:   label,
			Payload: string(b),
		},
		Color: color,
	}

	lastRowIdx := len(mg.Buttons) - 1
	mg.Buttons[lastRowIdx] = append(mg.Buttons[lastRowIdx], button)

	return mg
}

type MessageKeyboardButton struct {
	Action MessageKeyboardButtonAction `json:"action"`
	Color  string                      `json:"color,omitempty"`
}

type MessageKeyboardButtonAction struct {
	Type    string `json:"type"`
	Label   string `json:"label,omitempty"`
	Payload string `json:"payload,omitempty"`
	Link    string `json:"link,omitempty"`
	AppID   int    `json:"app_id,omitempty"`
	OwnerID int    `json:"owner_id,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

type MessageNewObject struct {
	Message    Message    `json:"message"`
	ClientInfo ClientInfo `json:"client_info"`
}

// https://dev.vk.com/reference/objects/user
type User struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Deactivated     string `json:"deactivated"`
	IsClosed        string `json:"is_closed"`
	CanAccessClosed bool   `json:"can_access_closed"`
	About           string `json:"about"`
	Activities      string `json:"activities"`
	Bdate           string `json:"bdate"`
	Blacklisted     int    `json:"blacklisted"`
	BlacklistedByMe int    `json:"blacklisted_by_me"`
	Books           string `json:"books"`
	PhotoId         string `json:"photo_id"`
	//Relation        int    `json:"relation"`
	// TODO: ...
}
