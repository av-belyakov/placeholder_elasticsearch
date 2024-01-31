package datamodels

type CustomerFields interface {
	Set(interface{}, interface{})
	Get() (string, int, string, string)
}

// ResponseMessageFromMispToTheHave содержит ответ для TheHive получаемый от MISP
type ResponseMessageFromMispToTheHave struct {
	Success  bool                        `json:"success" bson:"success"`
	Service  string                      `json:"service" bson:"service"`
	Error    error                       `json:"error" bson:"error"`
	Commands []ResponseCommandForTheHive `json:"commands" bson:"commands"`
}

// ResponseCommandForTheHive ответы с командами для TheHive
type ResponseCommandForTheHive struct {
	Command string `json:"command" bson:"command"`
	String  string `json:"string" bson:"string"`
	Name    string `json:"name" bson:"name"`
}

// MainMessageTheHive основное сообщение получаемое через NATS
type MainMessageTheHive struct {
	SourceMessageTheHive
	EventMessageTheHive
	ObservablesMessageTheHive
	TtpsMessageTheHive
}

// SourceMessageTheHive сообщение с информацией об источнике
// Source - источник
type SourceMessageTheHive struct {
	Source string `json:"source" bson:"source"`
}

type CustomFieldStringType struct {
	Order  int    `json:"order" bson:"order"`
	String string `json:"string" bson:"string"`
}

type CustomFieldDateType struct {
	Order int    `json:"order" bson:"order"`
	Date  string `json:"date" bson:"date"`
}

type CustomFieldIntegerType struct {
	Order   int `json:"order" bson:"order"`
	Integer int `json:"integer" bson:"integer"`
}
