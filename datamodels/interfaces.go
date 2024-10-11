package datamodels

type InformationFromEventEnricher interface {
	GetterRootId
	GetterSource
	GetterSensorsId
	GetterSensorInformation
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

type GetterIpAddresses interface {
	GetIpAddresses() []GetterIpAddressesInformation
}

type GetterIpAddressesInformation interface {
	GetIsSuccess() bool
	GetIp() string
	GetIpLocation() map[string]GetterIpLocation
}

type GetterIpLocation interface {
	GetCity() string
	GetCountry() string
	GetCountryCode() string
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
