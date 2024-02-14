package mongodbinteractions

import (
	"context"
	"errors"
	"fmt"
	"runtime"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"placeholder_elasticsearch/datamodels"
)

// AddNewCase добавляет новый кейс в БД
func (w *wrappers) AddNewCase(
	//data *datamodels.VerifiedTheHiveCase,
	data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	qp := QueryParameters{
		NameDB:         w.NameDB,
		ConnectDB:      w.ConnDB,
		CollectionName: "case_collection",
	}

	obj, ok := data.(*datamodels.VerifiedTheHiveCase)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type *datamodels.VerifiedTheHiveCase' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	//******************************************************************************
	//ищем документы подходящие под фильтр и удаляем их что бы избежать дублирования
	cur, err := qp.Find(bson.D{
		{Key: "source", Value: obj.GetSource()},
		{Key: "event.rootId", Value: obj.GetEvent().GetRootId()},
	})
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}

	listForDelete := []string{}
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
			MsgData: fmt.Sprintf("MongoDB , document with ID:'%s', author: '%s' data will be deleted", modelType.ID, modelType.Source),
			MsgType: "warning",
		}

		listForDelete = append(listForDelete, modelType.ID)
	}

	if len(listForDelete) > 0 {
		if _, err := qp.DeleteManyData(
			bson.D{{
				Key:   "@id",
				Value: bson.D{{Key: "$in", Value: listForDelete}}}},
			options.Delete(),
		); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-1),
				MsgType: "error",
			}
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
		DataMsg:  "subject_case",
		Count:    1,
	}
}

// AddNewCase добавляет новый кейс в БД
func (w *wrappers) AddNewAlert(
	data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings,
) {
	qp := QueryParameters{
		NameDB:         w.NameDB,
		ConnectDB:      w.ConnDB,
		CollectionName: "alert_collection",
	}

	obj, ok := data.(*datamodels.VerifiedTheHiveAlert)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type *datamodels.VerifiedTheHiveCase' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	//поиск схожего документа
	currentData := datamodels.VerifiedTheHiveAlert{}
	err := qp.FindOne(bson.D{
		{Key: "source", Value: obj.GetSource()},
		{Key: "event.rootId", Value: obj.GetEvent().GetRootId()},
	}).Decode(&currentData)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		//если похожего документа нет в БД
		currentData = *obj
	} else {
		//если похожий документ есть, выполняем замену старых значений новыми
		if _, err := currentData.GetEvent().ReplacingOldValues(*obj.GetEvent()); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("error replacing old values event '%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}
		if _, err := currentData.GetAlert().ReplacingOldValues(*obj.GetAlert()); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("error replacing old values alert '%s' %s:%d", err.Error(), f, l-2),
				MsgType: "error",
			}
		}
	}

	if _, err := qp.InsertData([]interface{}{currentData}, []mongo.IndexModel{
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
		DataMsg:  "subject_alert",
		Count:    1,
	}
}
