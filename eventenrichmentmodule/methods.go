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

// GetIpAddresses список найденной по ip адресам информации
func (e SettingsChanOutputEEM) GetIpAddresses() []string {
	return e.IpAddresses
}

// GetIsSuccess успешность выполненного запроса
func (e SettingsChanOutputEEM) GetIsSuccess(ip string) bool {
	var foundElem bool

	for _, v := range e.IpAddressesInfo {
		if v.Ip == ip {
			foundElem = v.IsSuccess
		}
	}

	return foundElem
}

// SearchCity поиск названия города
func (e SettingsChanOutputEEM) SearchCity(ip, sourceInfo string) (string, bool) {
	var foundElem string

	for _, v := range e.IpAddressesInfo {
		if v.Ip != ip {
			continue
		}

		if information, ok := v.Info[sourceInfo]; ok {
			return information.GetCity(), ok
		}
	}

	return foundElem, false
}

// SearchCountry поиск названия страны
func (e SettingsChanOutputEEM) SearchCountry(ip, sourceInfo string) (string, bool) {
	var foundElem string

	for _, v := range e.IpAddressesInfo {
		if v.Ip != ip {
			continue
		}

		if information, ok := v.Info[sourceInfo]; ok {
			return information.GetCountry(), ok
		}
	}

	return foundElem, false
}

// SearchCountryCode поиск кода страны
func (e SettingsChanOutputEEM) SearchCountryCode(ip, sourceInfo string) (string, bool) {
	var foundElem string

	for _, v := range e.IpAddressesInfo {
		if v.Ip != ip {
			continue
		}

		if information, ok := v.Info[sourceInfo]; ok {
			return information.GetCountryCode(), ok
		}
	}

	return foundElem, false
}

// GetIsSuccess успешность выполненного запроса
func (gii GeoIpInformation) GetIsSuccess() bool {
	return gii.IsSuccess
}

// GetIp ip адрес по которому осуществлялся поиск
func (gii GeoIpInformation) GetIp() string {
	return gii.Ip
}

// GetIpLocation подробная информация об ip адресе
func (gii GeoIpInformation) GetIpLocation(ip string) map[string]IpLocation {
	return gii.Info
}

// GetCity название города
func (ipl IpLocation) GetCity() string {
	return ipl.City
}

// GetCountry название страны
func (ipl IpLocation) GetCountry() string {
	return ipl.Country
}

// GetCountryCode код страны
func (ipl IpLocation) GetCountryCode() string {
	return ipl.CountryCode
}
