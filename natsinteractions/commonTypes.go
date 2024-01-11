package natsinteractions

// ModuleNATS инициализированный модуль
// ChanOutputMISP - канал для отправки полученных данных из модуля
type ModuleNATS struct {
	chanOutputNATS chan SettingsOutputChan
}

type SettingsOutputChan struct {
	MsgId string
	Data  []byte
}

type SettingsInputChan struct {
	Command, EventId, TaskId string
}

func (mnats ModuleNATS) GetDataReceptionChannel() <-chan SettingsOutputChan /*[]byte*/ {
	return mnats.chanOutputNATS
}

func (mnats ModuleNATS) SendingDataOutput(data SettingsOutputChan) {
	mnats.chanOutputNATS <- data
}
