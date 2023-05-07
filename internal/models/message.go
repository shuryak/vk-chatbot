package models

type Message struct {
	PeerID     int
	Text       string
	Attachment *Attachment
	Keyboard   *Keyboard
}

func NewMessage(peerID int, text string, attachment *Attachment, keyboard *Keyboard) *Message {
	return &Message{
		PeerID:     peerID,
		Text:       text,
		Attachment: attachment,
		Keyboard:   keyboard,
	}
}

func NewTextMessage(peerID int, text string) *Message {
	return &Message{
		PeerID: peerID,
		Text:   text,
	}
}
