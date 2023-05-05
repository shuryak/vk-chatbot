package objects

const (
	ButtonText     = "text"
	ButtonVKPay    = "vkpay"
	ButtonVKApp    = "open_app"
	ButtonLocation = "location"
	ButtonOpenLink = "open_link"
	ButtonCallback = "callback"
)

const (
	Primary   = "primary"
	Secondary = "secondary"
	Negative  = "negative"
	Positive  = "positive"
)

type Attachment interface {
	ToAttachment() string
}

type JSONObject interface {
	ToJSON() string
}
