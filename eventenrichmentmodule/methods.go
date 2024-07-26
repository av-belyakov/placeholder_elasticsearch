package eventenrichmentmodule

// GetRootId основной идентификатор кейса
func (e *SettingsChanOutputEEM) GetRootId() string {
	return e.RootId
}

// GetSource источник сформировавший кейс
func (e *SettingsChanOutputEEM) GetSource() string {
	return e.Source
}

// GetSourcesId список идентификаторов сенсоров
func (e *SettingsChanOutputEEM) GetSourcesId() []string {
	list := make([]string, 0, len(e.Sensors))

	for k := range e.Sensors {
		list = append(list, k)
	}

	return list
}

// GetSensorId идентификатор сенсора
func (e *SettingsChanOutputEEM) GetSensorId(key string) (sensorId string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		sensorId = sensor.SensorId

		return
	}

	return
}

// GetHostId идентификатор сенсора (необходимый для поиска в НКЦКИ)
func (e *SettingsChanOutputEEM) GetHostId(key string) (hostId string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		hostId = sensor.HostId

		return
	}

	return
}

// GetGeoCode географический код
func (e *SettingsChanOutputEEM) GetGeoCode(key string) (geoCode string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		geoCode = sensor.GeoCode

		return
	}

	return
}

// GetObjectArea сфера деятельности объекта
func (e *SettingsChanOutputEEM) GetObjectArea(key string) (objectArea string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		objectArea = sensor.ObjectArea

		return
	}

	return
}

// GetSubjectRF субъект Российской Федерации
func (e *SettingsChanOutputEEM) GetSubjectRF(key string) (subjectRF string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		subjectRF = sensor.SubjectRF

		return
	}

	return
}

// GetINN налоговый идентификатор
func (e *SettingsChanOutputEEM) GetINN(key string) (inn string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		inn = sensor.INN

		return
	}

	return
}

// GetHomeNet перечень домашних сетей
func (e *SettingsChanOutputEEM) GetHomeNet(key string) (homeNet string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		homeNet = sensor.HomeNet

		return
	}

	return
}

// GetOrgName наименование организации
func (e *SettingsChanOutputEEM) GetOrgName(key string) (orgName string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		orgName = sensor.OrgName

		return
	}

	return
}

// GetFullOrgName полное наименование организации
func (e *SettingsChanOutputEEM) GetFullOrgName(key string) (fullOrgName string, ok bool) {
	var sensor FoundSensorInformation
	if sensor, ok = e.Sensors[key]; ok {
		fullOrgName = sensor.FullOrgName

		return
	}

	return
}
