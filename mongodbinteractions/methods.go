package mongodbinteractions

import "placeholder_elasticsearch/datamodels"

func NewResultFoundSensorInformation() *ResultFoundSensorInformation {
	return &ResultFoundSensorInformation{
		SensorsId: []string(nil),
		Sensors:   []datamodels.SensorInformation(nil),
	}
}

// GetRootId основной идентификатор кейса
func (e ResultFoundSensorInformation) GetRootId() string {
	return e.RootId
}

// GetSource источник сформировавший кейс
func (e ResultFoundSensorInformation) GetSource() string {
	return e.Source
}

// GetSensorsId список идентификаторов сенсоров
func (e ResultFoundSensorInformation) GetSensorsId() []string {
	list := make([]string, 0, len(e.Sensors))

	for _, v := range e.Sensors {
		list = append(list, v.SensorId)
	}

	return list
}

// GetHostId идентификатор сенсора (необходимый для поиска в НКЦКИ)
func (e ResultFoundSensorInformation) GetHostId(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.HostId

			break
		}
	}

	return foundElem
}

// GetGeoCode географический код
func (e ResultFoundSensorInformation) GetGeoCode(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.GeoCode

			break
		}
	}

	return foundElem
}

// GetObjectArea сфера деятельности объекта
func (e ResultFoundSensorInformation) GetObjectArea(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.ObjectArea

			break
		}
	}

	return foundElem
}

// GetSubjectRF субъект Российской Федерации
func (e ResultFoundSensorInformation) GetSubjectRF(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.SubjectRF

			break
		}
	}

	return foundElem
}

// GetINN налоговый идентификатор
func (e ResultFoundSensorInformation) GetINN(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.INN

			break
		}
	}

	return foundElem
}

// GetHomeNet перечень домашних сетей
func (e ResultFoundSensorInformation) GetHomeNet(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.HomeNet

			break
		}
	}

	return foundElem
}

// GetOrgName наименование организации
func (e ResultFoundSensorInformation) GetOrgName(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.OrgName

			break
		}
	}

	return foundElem
}

// GetFullOrgName полное наименование организации
func (e ResultFoundSensorInformation) GetFullOrgName(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.FullOrgName

			break
		}
	}

	return foundElem
}
