package eventenrichmentmodule

// SettingsChanInputEEM параметры для ПРИЕМА данных в модуль
// RootId - основной идентификатор
// Source - источник
// SensorsId - искомые идентификаторы сенсоров
type SettingsChanInputEEM struct {
	RootId    string
	Source    string
	SensorsId []string
}

// SettingsChanInputEEM параметры для ОТПРАВКИ данных из модуля
// RootId - основной идентификатор
// Source - источник
// Sensors - найденная информация, где ключ карты это идентификатор сенсора
type SettingsChanOutputEEM struct {
	RootId  string
	Source  string
	Sensors map[string]FoundSensorInformation
}

// FoundSensorInformation содержит найденную о сенсоре информацию
// SensorId - идентификатор сенсора
// HostId - идентификатор сенсора, специальный, для поиска информации в НКЦКИ
// GeoCode - географический код страны
// ObjectArea - сфера деятельности объекта
// SubjectRF - субъект Российской Федерации
// INN - налоговый идентификатор
// HomeNet - перечень домашних сетей
// OrgName - наименование организации
// FullOrgName - полное наименование организации
type FoundSensorInformation struct {
	SensorId    string
	HostId      string
	GeoCode     string
	ObjectArea  string
	SubjectRF   string
	INN         string
	HomeNet     string
	OrgName     string
	FullOrgName string
}

// EventEnrichmentModule инициализированный модуль обогащения кейсов
// ChanInputModule - канал для отправки данных В модуль
// ChanOutputModule - канал для принятия данных ИЗ модуля
type EventEnrichmentModule struct {
	ChanInputModule  chan SettingsChanInputEEM
	ChanOutputModule chan SettingsChanOutputEEM
}
