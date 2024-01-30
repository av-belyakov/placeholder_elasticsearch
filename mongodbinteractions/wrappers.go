package mongodbinteractions

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"placeholder_elasticsearch/datamodels"
)

func (w *wrappers) AddNewCase(data []interface{}, logging chan<- datamodels.MessageLogging) {
	qp := QueryParameters{
		NameDB:         w.NameDB,
		CollectionName: "case_collection",
		ConnectDB:      w.ConnDB,
	}

	if _, err := qp.InsertData(data, []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "@id", Value: 1},
				{Key: "@timestamp", Value: 1},
			},
			Options: &options.IndexOptions{},
		}, {
			Keys: bson.D{
				{Key: "source", Value: 1},
			},
			Options: &options.IndexOptions{},
		},
	}); err != nil {
		fmt.Println("func 'AddNewCase' ERROR:", err)

		/*
			Для тестов пока отключаем

			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}*/
	}
}
