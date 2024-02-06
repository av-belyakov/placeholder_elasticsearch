package confighandler

type ConfigApp struct {
	CommonAppConfig
	AppConfigNATS
	AppConfigElasticSearch
	AppConfigMongoDB
	AppConfigRulesProcMsg
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
