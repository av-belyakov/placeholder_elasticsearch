package confighandler

type ConfigApp struct {
	CommonAppConfig
	AppConfigNATS
	AppConfigElasticSearch
	AppConfigMongoDB
}

type CommonAppConfig struct {
	LogList []LogSet
	Zabbix  ZabbixOptions
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
	IsTransmit   bool   `yaml:"isTransmit"`
	TimeInterval int    `yaml:"timeInterval"`
	NetworkPort  int    `yaml:"networkPort"`
	NetworkHost  string `yaml:"networkHost"`
	ZabbixHost   string `yaml:"zabbixHost"`
	ZabbixKey    string `yaml:"zabbixKey"`
	Handshake    string `yaml:"handshake"`
}

type AppConfigNATS struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type AppConfigElasticSearch struct {
	Send   bool   `yaml:"send"`
	Prefix string `yaml:"prefix"`
	Index  string `yaml:"index"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
}

type AppConfigMongoDB struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	NameDB string `yaml:"namedb"`
}
