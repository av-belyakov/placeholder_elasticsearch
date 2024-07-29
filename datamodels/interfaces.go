package datamodels

type InformationFromEventEnricher interface {
	GetterRootId
	GetterSource
	GetterSensorsId
	GetterSensorInformation
}

type GetterRootId interface {
	GetRootId() string
}

type GetterSource interface {
	GetSource() string
}

type GetterSensorsId interface {
	GetSensorsId() []string
}

type GetterSensorInformation interface {
	GetHostId(string) string
	GetGeoCode(string) string
	GetObjectArea(string) string
	GetSubjectRF(string) string
	GetINN(string) string
	GetHomeNet(string) string
	GetOrgName(string) string
	GetFullOrgName(string) string
}
