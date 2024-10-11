package coremodule

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/elasticsearchinteractions"
	"placeholder_elasticsearch/eventenrichmentmodule"
	"placeholder_elasticsearch/natsinteractions"
)

type VerifiedElasticsearchFormatCaseOptions struct {
	msgId    string                                             //идентификатор сообщения полученного из модуля API NATS
	cs       *coreStorage                                       //хранилище временной информации
	input    <-chan datamodels.ChanOutputDecodeJSON             //канал приема данных при декодировании JSON
	done     <-chan bool                                        //канал индикации останова обработчика
	natsChan chan<- natsinteractions.SettingsInputChan          //канал для взаимодействия с API модуля NATS
	esmChan  chan<- elasticsearchinteractions.SettingsInputChan //канал для взаимодействия с модулем API Elasticsearch
	eemChan  chan<- eventenrichmentmodule.SettingsChanInputEEM  //канал для отправки запросов к модулю обогащения
	logging  chan<- datamodels.MessageLogging                   //канал логирования
}
