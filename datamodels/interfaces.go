package datamodels

type InformationFromEventEnricher interface {
	GetterRootId
	GetterSource
	GetterSourcesId
	GetterSensorInformation
}

type GetterSourcesId interface {
	GetSourcesId() []string
}

type GetterRootId interface {
	GetRootId() string
}

type GetterSource interface {
	GetSource() string
}

type GetterSensorInformation interface {
	GetSensorId(string) (string, bool)
	GetHostId(string) (string, bool)
	GetGeoCode(string) (string, bool)
	GetObjectArea(string) (string, bool)
	GetSubjectRF(string) (string, bool)
	GetINN(string) (string, bool)
	GetHomeNet(string) (string, bool)
	GetOrgName(string) (string, bool)
	GetFullOrgName(string) (string, bool)
}
