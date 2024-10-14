package datamodels

type InformationFromEventEnricher interface {
	GetterRootId
	GetterSource
	GetterSensorsId
	GetterSensorInformation
	GetterIpAddresses
	GetterIpAddressesInformation
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
	GetHostId(sid string) string
	GetGeoCode(sid string) string
	GetObjectArea(sid string) string
	GetSubjectRF(sid string) string
	GetINN(sid string) string
	GetHomeNet(sid string) string
	GetOrgName(sid string) string
	GetFullOrgName(sid string) string
}

type GetterIpAddresses interface {
	GetIpAddresses() []string
}

type GetterIpAddressesInformation interface {
	GetIsSuccess(ip string) bool
	SearchCity(ip, sourceInfo string) (string, bool)
	SearchCountry(ip, sourceInfo string) (string, bool)
	SearchCountryCode(ip, sourceInfo string) (string, bool)
}
