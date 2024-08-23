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

// GetSensorsEventenrichment запрос на получения дополнительной информации о сенсоре
func (w *wrappers) GetSensorsEventenrichment(
	rootId string,
	source string,
	data interface{},
	chanOutput chan<- ModuleDataBaseInteractionChannel,
	logging chan<- datamodels.MessageLogging) {

	listSensorId, ok := data.([]string)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type []string' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	options := options.Find().SetAllowDiskUse(true)
	collection := w.ConnDB.Database(w.NameDB).Collection("collection_sensor_information")
	cur, err := collection.Find(
		context.Background(),
		(bson.D{
			{Key: "sensorId", Value: bson.D{
				{Key: "$in", Value: listSensorId}},
			},
		}), options)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-1),
			MsgType: "error",
		}

		return
	}

	result := NewResultFoundSensorInformation()
	result.RootId = rootId
	result.Source = source

	for cur.Next(context.Background()) {
		var sensorInfo datamodels.SensorInformation
		if err := cur.Decode(&sensorInfo); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-1),
				MsgType: "error",
			}

			return
		}

		result.SensorsId = append(result.SensorsId, sensorInfo.SensorId)
		result.Sensors = append(result.Sensors, sensorInfo)
	}

	chanOutput <- ModuleDataBaseInteractionChannel{
		Section: "handling eventenrichment",
		Command: "sensor eventenrichment response",
		Data:    result,
	}
}

// AddNewSensorsEventenrichment добавляет дополнительную информацию о сенсорах
func (w *wrappers) AddNewSensorsEventenrichment(
	data interface{},
	logging chan<- datamodels.MessageLogging) {

	requestMongoDB := func(sensorId string, sensorSettings datamodels.InformationFromEventEnricher) error {
		collection := w.ConnDB.Database(w.NameDB).Collection("collection_sensor_information")
		opts := options.FindOne()
		currentData := datamodels.NewSensorInformation()
		err := collection.FindOne(
			context.TODO(),
			bson.D{
				{Key: "sensorId", Value: sensorId},
			},
			opts,
		).Decode(currentData)
		if !errors.Is(err, mongo.ErrNoDocuments) {
			if _, err := collection.DeleteMany(
				context.TODO(),
				bson.D{{
					Key:   "sensorId",
					Value: bson.D{{Key: "$in", Value: []string{currentData.SensorId}}}}},
				options.Delete()); err != nil {

				return err
			}
		}

		newData := datamodels.NewSensorInformation()
		newData.SetSensorId(sensorId)
		newData.SetHostId(sensorSettings.GetHostId(sensorId))
		newData.SetGeoCode(sensorSettings.GetGeoCode(sensorId))
		newData.SetObjectArea(sensorSettings.GetObjectArea(sensorId))
		newData.SetSubjectRF(sensorSettings.GetSubjectRF(sensorId))
		newData.SetINN(sensorSettings.GetINN(sensorId))
		newData.SetHomeNet(sensorSettings.GetHomeNet(sensorId))
		newData.SetOrgName(sensorSettings.GetOrgName(sensorId))
		newData.SetFullOrgName(sensorSettings.GetFullOrgName(sensorId))

		//если похожего документа нет в БД
		if _, err := collection.InsertMany(context.TODO(), []interface{}{newData}); err != nil {
			return err
		}

		return nil
	}

	sensorSettings, ok := data.(datamodels.InformationFromEventEnricher)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type *datamodels.InformationFromEventEnricher' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	for _, sensorId := range sensorSettings.GetSensorsId() {
		if err := requestMongoDB(sensorId, sensorSettings); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-1),
				MsgType: "error",
			}

			continue
		}
	}
}

// AddNewCase добавляет новый кейс в БД
func (w *wrappers) AddNewCase(
	data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {
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
		{Key: "event.commoneventtype.rootId", Value: obj.GetEvent().GetRootId()},
	})
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
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
				{Key: "event.commoneventtype.rootId", Value: 1},
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

// AddNewAlert выполняет замену уже существующего объекта типа Alert
// либо добавляет новый, если в БД нет объекта с заданными параметрами
func (w *wrappers) AddNewAlert(
	data interface{},
	logging chan<- datamodels.MessageLogging,
	counting chan<- datamodels.DataCounterSettings) {
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
		{Key: "event.commoneventtype.rootId", Value: obj.GetEvent().GetRootId()},
	}).Decode(&currentData)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		//если похожего документа нет в БД
		currentData = *obj
	} else {
		//если похожий документ есть, удаляем старый документ и выполняем
		//замену старых значений в полученном из БД документе новыми значениями
		if _, err := qp.DeleteManyData(
			bson.D{{
				Key:   "@id",
				Value: bson.D{{Key: "$in", Value: []string{obj.GetID()}}}}},
			options.Delete(),
		); err != nil {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-1),
				MsgType: "error",
			}

			return
		}

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

	if _, err := qp.InsertData([]interface{}{currentData},
		[]mongo.IndexModel{
			{
				Keys: bson.D{
					{Key: "@id", Value: 1},
				},
				Options: &options.IndexOptions{},
			}, {
				Keys: bson.D{
					{Key: "source", Value: 1},
					{Key: "event.commoneventtype.rootId", Value: 1},
				},
				Options: &options.IndexOptions{},
			}}); err != nil {
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
