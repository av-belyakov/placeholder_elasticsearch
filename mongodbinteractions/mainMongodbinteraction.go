package mongodbinteractions

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"placeholder_elasticsearch/confighandler"
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/memorytemporarystorage"
)

// MongoDBModule содержит описание каналов для взаимодействия с БД MongoDB
// ChanInputModule - канал для отправки данных В модуль
// ChanOutputModule - канал для принятия данных ИЗ модуля
type MongoDBModule struct {
	ChanInputModule, ChanOutputModule chan ModuleDataBaseInteractionChannel
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

func HandlerMongoDB(conf confighandler.AppConfigMongoDB,
	storageApp *memorytemporarystorage.CommonStorageTemporary,
	logging chan<- datamodels.MessageLogging) (*MongoDBModule, error) {
	channels := &MongoDBModule{
		ChanInputModule:  make(chan ModuleDataBaseInteractionChannel),
		ChanOutputModule: make(chan ModuleDataBaseInteractionChannel),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	description := ConnectionDescriptorMongoDB{
		databaseName: conf.NameDB,
		ctx:          ctx,
		ctxCancel:    cancel,
	}

	conn, err := NewConnection(ctx, conf)
	if err != nil {
		return channels, err
	}

	description.connection = conn
	description.Routing(channels, logging)

	return channels, nil
}

func NewConnection(ctx context.Context, conf confighandler.AppConfigMongoDB) (*mongo.Client, error) {
	clientOption := options.Client()
	clientOption.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    conf.NameDB,
		Username:      conf.User,
		Password:      conf.Passwd,
	})

	confPath := fmt.Sprintf("mongodb://%s:%d/%s", conf.Host, conf.Port, conf.NameDB)
	connect, err := mongo.Connect(ctx, options.Client().ApplyURI(confPath))
	if err != nil {
		return connect, err
	}

	if err = connect.Ping(ctx, readpref.Primary()); err != nil {
		return connect, err
	}

	log.Printf("Create connection with MongoDB (%s:%d)\n", conf.Host, conf.Port)

	return connect, nil
}

func (conn ConnectionDescriptorMongoDB) Routing(channels *MongoDBModule, logging chan<- datamodels.MessageLogging) {
	go func() {
		for data := range channels.ChanInputModule {
			switch data.Section {
			case "add new case":
				fmt.Println("func 'Routing' data.Section:", data.Section)
			}
		}
	}()
}

/*
func Routing(
	chanOutput chan<- datamodels.ModuleDataBaseInteractionChannel,
	nameDb string,
	cdmdb ConnectionDescriptorMongoDB,
	tst *memorytemporarystoragecommoninformation.TemporaryStorageType,
	chanInput <-chan datamodels.ModuleDataBaseInteractionChannel) {

	ws := wrappersSetting{
		NameDB:       nameDb,
		ConnectionDB: cdmdb,
	}

	for data := range chanInput {
		switch data.Section {
		case "handling stix object":
			go ws.wrapperFuncTypeHandlingSTIXObject(chanOutput, data, tst)

		case "handling managing collection stix objects":
			go ws.wrapperFuncTypeHandlingManagingCollectionSTIXObjects(chanOutput, data, tst)

		case "handling managing differences objects collection":
			go ws.wrapperFuncTypeHandlingManagingDifferencesObjectsCollection(chanOutput, data, tst)

		case "handling search requests":
			go ws.wrapperFuncTypeHandlingSearchRequests(chanOutput, data, tst)

		case "handling reference book":
			go ws.wrapperFuncTypeHandlingReferenceBook(chanOutput, data, tst)

		case "handling technical part":
			go ws.wrapperFuncTypeTechnicalPart(chanOutput, data, tst)

		case "handling statistical requests":
			go ws.wrapperFuncTypeHandlingStatisticalRequests(chanOutput, data, tst)

		}
	}
}
*/
