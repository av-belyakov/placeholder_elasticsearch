package mongodbinteractions

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"placeholder_elasticsearch/datamodels"
)

// MongoDBModule содержит описание каналов для взаимодействия с БД MongoDB
// ChanInputModule - канал для отправки данных В модуль
// ChanOutputModule - канал для принятия данных ИЗ модуля
type MongoDBModule struct {
	ChanInputModule  chan SettingsInputChan
	ChanOutputModule chan ModuleDataBaseInteractionChannel
}

// SettingsInputChan
// Section - секция обработки данных
// Command - команда
// RootId - основной идентификатор
// Source - источник
// Data - данные
type SettingsInputChan struct {
	Section        string
	Command        string
	RootId         string
	Source         string
	Data           interface{}
	VerifiedObject *datamodels.VerifiedTheHiveCase
}

// ModuleDataBaseInteractionChannel описание типов данных циркулирующих между модулем взаимодействия с БД и Ядром приложения
// Section - секция обработки данных
// Command - команда
// RootId - основной идентификатор
// Source - источник
// Data - данные
type ModuleDataBaseInteractionChannel struct {
	Section string
	Command string
	Data    interface{}
}

type wrappers struct {
	AdditionalRequestParameters interface{}
	NameDB                      string
	ConnDB                      *mongo.Client
}

// ConnectionDescriptorMongoDB дескриптор соединения с БД MongoDB
// databaseName - имя базы данных
// connection - дескриптор соединения
// ctx - контекст переносит крайний срок, сигнал отмены и другие значения через границы API
// ctxCancel - метод закрытия контекста
type ConnectionDescriptorMongoDB struct {
	databaseName string
	connection   *mongo.Client
	ctx          context.Context
	ctxCancel    context.CancelFunc
}

// ResultFoundSensorInformation содержит подробную информацию о найденных сенсорах
// RootId - основной идентификатор
// Source - наименование источника
// SensorsId - список искомых идентификаторов сенсоров
// Sensors - найденная по сенсорам информация
type ResultFoundSensorInformation struct {
	RootId    string
	Source    string
	SensorsId []string
	Sensors   []datamodels.SensorInformation
}
