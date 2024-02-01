package mongodbinteractions

import (
	"context"
	"fmt"
	"runtime"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"placeholder_elasticsearch/datamodels"
)

// AddNewCase добавляет новый кейс в БД
func (w *wrappers) AddNewCase(
	data *datamodels.VerifiedTheHiveCase,
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	qp := QueryParameters{
		NameDB:         w.NameDB,
		CollectionName: "case_collection",
		ConnectDB:      w.ConnDB,
	}

	//******************************************************************************
	//ищем документы подходящие под фильтр и удаляем их что бы избежать дублирования
	cur, err := qp.Find(bson.D{
		{Key: "source", Value: data.GetSource()},
		{Key: "event.rootId", Value: data.GetEvent().GetRootId()},
	})
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}

	listForDelete := []string(nil)
	for cur.Next(context.Background()) {
		var modelType struct {
			ID     string `bson:"@id"`
			Source string `bson:"source"`
		}
		if err := cur.Decode(&modelType); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}

			continue
		}

		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("MongoDB , document with ID:'%s', author: '%s' to be deleted", modelType.ID, modelType.Source),
			MsgType: "warning",
		}

		listForDelete = append(listForDelete, modelType.ID)
	}

	if _, err := qp.DeleteManyData(listForDelete, options.Delete()); err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}
	//******************************************************************************

	if _, err := qp.InsertData([]interface{}{data}, []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "@id", Value: 1},
			},
			Options: &options.IndexOptions{},
		}, {
			Keys: bson.D{
				{Key: "source", Value: 1},
				{Key: "event.rootId", Value: 1},
			},
			Options: &options.IndexOptions{},
		},
	}); err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	//счетчик
	counting <- datamodels.DataCounterSettings{
		DataType: "update count insert MongoDB",
		Count:    1,
	}
}
