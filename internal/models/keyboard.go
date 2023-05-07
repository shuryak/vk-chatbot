package models

type ButtonColor string

const (
	PrimaryColor   ButtonColor = "primary"
	SecondaryColor ButtonColor = "secondary"
	NegativeColor  ButtonColor = "negative"
	PositiveColor  ButtonColor = "positive"
)

type Button struct {
	Text    string
	Color   ButtonColor
	Payload Payload
}

type Keyboard struct {
	Inline  bool
	OneTime bool
	Buttons [][]Button
}

func NewInlineKeyboard() *Keyboard {
	return &Keyboard{Inline: true, Buttons: [][]Button{}}
}

func NewKeyboard(oneTime bool) *Keyboard {
	return &Keyboard{OneTime: true, Buttons: [][]Button{}}
}

func (k *Keyboard) AddRow() *Keyboard {
	if len(k.Buttons) == 0 {
		k.Buttons = make([][]Button, 1)
	} else {
		row := make([]Button, 0)
		k.Buttons = append(k.Buttons, row)
	}

	return k
}

func (k *Keyboard) AddButtonWithCommandOnly(text string, color ButtonColor, cmd ButtonCommand) *Keyboard {
	return k.AddButton(text, color, *NewPayloadWithCommandOnly(cmd))
}

func (k *Keyboard) AddButton(text string, color ButtonColor, p Payload) *Keyboard {
	lastRowIdx := len(k.Buttons) - 1
	k.Buttons[lastRowIdx] = append(k.Buttons[lastRowIdx], Button{
		Text:    text,
		Color:   color,
		Payload: p,
	})

	return k
}
