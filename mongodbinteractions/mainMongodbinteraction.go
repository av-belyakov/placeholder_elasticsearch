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
)

// MongoDBModule содержит описание каналов для взаимодействия с БД MongoDB
// ChanInputModule - канал для отправки данных В модуль
// ChanOutputModule - канал для принятия данных ИЗ модуля
type MongoDBModule struct {
	ChanInputModule  chan SettingsInputChan
	ChanOutputModule chan ModuleDataBaseInteractionChannel
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

func HandlerMongoDB(
	conf confighandler.AppConfigMongoDB,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) (*MongoDBModule, error) {

	channels := &MongoDBModule{
		ChanInputModule:  make(chan SettingsInputChan),
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
	description.Routing(channels, logging, counting)

	return channels, nil
}

func NewConnection(ctx context.Context, conf confighandler.AppConfigMongoDB) (*mongo.Client, error) {
	confPath := fmt.Sprintf("mongodb://%s:%d/%s", conf.Host, conf.Port, conf.NameDB)

	connect, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    conf.NameDB,
		Username:      conf.User,
		Password:      conf.Passwd,
	}).ApplyURI(confPath))
	if err != nil {
		return connect, err
	}

	if err = connect.Ping(ctx, readpref.Primary()); err != nil {
		return connect, err
	}

	log.Printf("Create connection with MongoDB (%s:%d)\n", conf.Host, conf.Port)

	return connect, nil
}

func (conn ConnectionDescriptorMongoDB) Routing(
	channels *MongoDBModule,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	ws := wrappers{
		NameDB: conn.databaseName,
		ConnDB: conn.connection,
	}

	go func() {
		defer func() {
			_ = conn.connection.Disconnect(context.TODO())
		}()

		for data := range channels.ChanInputModule {
			switch data.Section {
			case "handling case":
				if data.Command == "add new case" {
					go ws.AddNewCase(data.Data, logging, counting)
				}

			case "handling alert":
				if data.Command == "add new alert" {
					go ws.AddNewAlert(data.Data, logging, counting)
				}
			}
		}
	}()
}
