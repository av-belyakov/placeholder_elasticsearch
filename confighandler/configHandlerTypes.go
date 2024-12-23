package confighandler

// ConfigApp конфигурационная структура
type ConfigApp struct {
	CommonAppConfig        //общие настройки
	AppConfigNATS          //настройки для доступа к API NATS
	AppConfigElasticSearch //настройки для доступа к API Elasticsearch
	AppConfigMongoDB       //настройки для доступа к MongoDB
	AppConfigRulesProcMsg  //правила фильтрации
	AppConfigMapping       //настройки сопоставления значений
}

type CommonAppConfig struct {
	LogList       []LogSet
	NCIRCC        NCIRCCOptions
	Zabbix        ZabbixOptions
	GeoIpJsonRPC  GeoIPJsonRPCOptions
	ZabbixJsonRPC ZabbixJsonRPCOptions
}

type Logs struct {
	Logging []LogSet
}

type LogSet struct {
	WritingStdout bool   `yaml:"writingStdout"`
	WritingFile   bool   `yaml:"writingFile"`
	MaxFileSize   int    `yaml:"maxFileSize"`
	MsgTypeName   string `yaml:"msgTypeName"`
	PathDirectory string `yaml:"pathDirectory"`
}

type ZabbixSet struct {
	Zabbix ZabbixOptions
}

type ZabbixOptions struct {
	NetworkPort int         `yaml:"networkPort"`
	NetworkHost string      `yaml:"networkHost"`
	ZabbixHost  string      `yaml:"zabbixHost"`
	EventTypes  []EventType `yaml:"eventType"`
}

type NCIRCCSet struct {
	NCIRCC NCIRCCOptions
}

type NCIRCCOptions struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

type GEOIPJSONRPCSet struct {
	GEOIPJSONRPC GeoIPJsonRPCOptions
}

type GeoIPJsonRPCOptions struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
	Path string `yaml:"path"`
}

type ZABBIXJSONRPCSet struct {
	ZABBIXJSONRPC ZabbixJsonRPCOptions
}

type ZabbixJsonRPCOptions struct {
	ConnectionTimeout int    `yaml:"connectionTimeout"`
	NetworkHost       string `yaml:"networkHost"`
	Login             string `yaml:"login"`
	Passwd            string `yaml:"passwd"`
}

type EventType struct {
	IsTransmit bool      `yaml:"isTransmit"`
	EventType  string    `yaml:"eventType"`
	ZabbixKey  string    `yaml:"zabbixKey"`
	Handshake  Handshake `yaml:"handshake"`
}

type Handshake struct {
	TimeInterval int    `yaml:"timeInterval"`
	Message      string `yaml:"message"`
}

type AppConfigNATS struct {
	Port         int    `yaml:"port"`
	Host         string `yaml:"host"`
	SubjectCase  string `yaml:"subject_case"`
	SubjectAlert string `yaml:"subject_alert"`
}

type AppConfigElasticSearch struct {
	Port        int    `yaml:"port"`
	PrefixCase  string `yaml:"prefix_case"`
	IndexCase   string `yaml:"index_case"`
	PrefixAlert string `yaml:"prefix_alert"`
	IndexAlert  string `yaml:"index_alert"`
	User        string `yaml:"user"`
	Passwd      string `yaml:"passwd"`
	Host        string `yaml:"host"`
}

type AppConfigMongoDB struct {
	Port   int    `yaml:"port"`
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	NameDB string `yaml:"namedb"`
}

type AppConfigRulesProcMsg struct {
	Directory string `yaml:"directory"`
	FileCase  string `yaml:"file_case"`
	FileAlert string `yaml:"file_alert"`
}

type AppConfigMapping struct {
	AreaActivity []ObjectAreaActivity
}

type MAPPINGAREAACTIVITYSet struct {
	MAPPINGAREAACTIVITY []ObjectAreaActivity
}

type ObjectAreaActivity struct {
	ApprovedName   string   `yaml:"approvedName"`
	VariationsName []string `yaml:"variationsName"`
}
