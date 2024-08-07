package eventenrichmentmodule

// GetRootId основной идентификатор кейса
func (e SettingsChanOutputEEM) GetRootId() string {
	return e.RootId
}

// GetSource источник сформировавший кейс
func (e SettingsChanOutputEEM) GetSource() string {
	return e.Source
}

// GetSensorsId список идентификаторов сенсоров
func (e SettingsChanOutputEEM) GetSensorsId() []string {
	list := make([]string, 0, len(e.Sensors))

	for _, v := range e.Sensors {
		list = append(list, v.SensorId)
	}

	return list
}

// GetHostId идентификатор сенсора (необходимый для поиска в НКЦКИ)
func (e SettingsChanOutputEEM) GetHostId(sensorId string) string {
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
func (e SettingsChanOutputEEM) GetGeoCode(sensorId string) string {
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
func (e SettingsChanOutputEEM) GetObjectArea(sensorId string) string {
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
func (e SettingsChanOutputEEM) GetSubjectRF(sensorId string) string {
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
func (e SettingsChanOutputEEM) GetINN(sensorId string) string {
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
func (e SettingsChanOutputEEM) GetHomeNet(sensorId string) string {
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
func (e SettingsChanOutputEEM) GetOrgName(sensorId string) string {
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
func (e SettingsChanOutputEEM) GetFullOrgName(sensorId string) string {
	var foundElem string
	for _, v := range e.Sensors {
		if v.SensorId == sensorId {
			foundElem = v.FullOrgName

			break
		}
	}

	return foundElem
}
