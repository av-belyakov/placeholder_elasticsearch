package natsinteractions

// NewResponseMessage формирует новый тип ResponseMessageFromMispToTheHave с предустановленными значениями
func NewResponseMessage() *ResponseMessageFromMispToTheHave {
	return &ResponseMessageFromMispToTheHave{
		Success: true,
		Service: "ES",
		Commands: []ResponseCommandForTheHive{
			{
				Command: "addtag",
				String:  "Webhook: send=\"ES\"",
			},
		},
	}
}

// ResponseMessageAddNewCommand добавляет объект с новой командой для сообщения
func (rm *ResponseMessageFromMispToTheHave) ResponseMessageAddNewCommand(rcm ResponseCommandForTheHive) {
	rm.Commands = append(rm.Commands, rcm)
}

// GetResponseMessageFromMispToTheHave возвращает сформированное сообщение для отправки в TheHive
func (rm *ResponseMessageFromMispToTheHave) GetResponseMessageFromMispToTheHave() ResponseMessageFromMispToTheHave {
	return *rm
}

// GetDataReceptionChannel возвращает канал для приема сообщений от модуля API NATS
func (mnats ModuleNATS) GetDataReceptionChannel() <-chan SettingsOutputChan {
	return mnats.chanOutputNATS
}

// SendingDataOutput отправляет данные в канал приема сообщений от модуля API NATS
func (mnats ModuleNATS) SendingDataOutput(data SettingsOutputChan) {
	mnats.chanOutputNATS <- data
}

// GetDataDeliveryChannel возвращает канал для приема сообщений в модул API NATS
func (mnats ModuleNATS) GetDataDeliveryChannel() chan SettingsInputChan {
	return mnats.chanInputNATS
}

// SendingDataInput отправляет данные в канал передачи сообщений в модул API NATS
func (mnats ModuleNATS) SendingDataInput(data SettingsInputChan) {
	mnats.chanInputNATS <- data
}
