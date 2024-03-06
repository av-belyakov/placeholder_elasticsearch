package datamodels

// TtpsMessageEs список TTP сообщений
type TtpsMessageEs struct {
	Ttp map[string][]TtpMessage `json:"ttp" bson:"ttp"`
}
