package datamodels

type CustomerFields interface {
	Set(interface{}, interface{})
	Get() (string, int, string, string)
}

// ResponseMessageFromMispToTheHave содержит ответ для TheHive получаемый от MISP
type ResponseMessageFromMispToTheHave struct {
	Success  bool                        `json:"success"`
	Service  string                      `json:"service"`
	Error    error                       `json:"error"`
	Commands []ResponseCommandForTheHive `json:"commands"`
}

// ResponseCommandForTheHive ответы с командами для TheHive
type ResponseCommandForTheHive struct {
	Command string `json:"command"`
	String  string `json:"string"`
	Name    string `json:"name"`
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
	Source string `json:"source"`
}

type CustomFieldStringType struct {
	Order  int    `json:"order"`
	String string `json:"string"`
}

type CustomFieldDateType struct {
	Order int    `json:"order"`
	Date  string `json:"date"`
}

type CustomFieldIntegerType struct {
	Order   int `json:"order"`
	Integer int `json:"integer"`
}
