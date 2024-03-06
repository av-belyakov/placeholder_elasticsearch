package datamodels

// TtpsMessageTheHive список TTP сообщений
type TtpsMessageTheHive struct {
	Ttp []TtpMessage `json:"ttp" bson:"ttp"`
}
