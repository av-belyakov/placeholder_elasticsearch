package eventenrichmentmodule

// SettingsChanInputEEM параметры для ПРИЕМА данных в модуль
type SettingsChanInputEEM struct {
	RootId      string   //основной идентификатор
	Source      string   //источник
	SensorsId   []string //искомые идентификаторы сенсоров
	IpAddresses []string //искомые ip адреса
}

// SettingsChanInputEEM параметры для ОТПРАВКИ данных из модуля
type SettingsChanOutputEEM struct {
	RootId      string                   //основной идентификатор
	Source      string                   //источник
	SensorsId   []string                 //список идентификаторов сенсоров информация по которым не была найдена
	Sensors     []FoundSensorInformation //найденная информация, где ключ карты это идентификатор сенсора
	IpAddresses []GeoIpInformation       //найденная по ip адресам информация
}

// FoundSensorInformation содержит найденную о сенсоре информацию
type FoundSensorInformation struct {
	SensorId    string `json:"sensorId"`    //идентификатор сенсора
	HostId      string `json:"hostId"`      //идентификатор сенсора, специальный, для поиска информации в НКЦКИ
	GeoCode     string `json:"geoCode"`     //географический код страны
	ObjectArea  string `json:"objectArea"`  //сфера деятельности объекта
	SubjectRF   string `json:"subjectRF"`   //субъект Российской Федерации
	INN         string `json:"inn"`         //налоговый идентификатор
	HomeNet     string `json:"homeNet"`     //перечень домашних сетей
	OrgName     string `json:"orgName"`     //наименование организации
	FullOrgName string `json:"fullOrgName"` //полное наименование организации
}

// EventEnrichmentModule инициализированный модуль обогащения кейсов
type EventEnrichmentModule struct {
	ChanInputModule  chan SettingsChanInputEEM  //канал для отправки данных В модуль
	ChanOutputModule chan SettingsChanOutputEEM //канал для принятия данных ИЗ модуля
}
