package confighandler

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strconv"

	"github.com/spf13/viper"

	"placeholder_elasticsearch/supportingfunctions"
)

func NewConfig(rootDir string) (ConfigApp, error) {
	conf := ConfigApp{}
	var envList map[string]string = map[string]string{
		"GO_PHELASTIC_MAIN": "",

		//Подключение к NATS
		"GO_PHELASTIC_NHOST": "",
		"GO_PHELASTIC_NPORT": "",

		// Подключение к СУБД Elasticsearch
		"GO_PHELASTIC_ESSEND":   "",
		"GO_PHELASTIC_ESHOST":   "",
		"GO_PHELASTIC_ESPORT":   "",
		"GO_PHELASTIC_ESPREFIX": "",
		"GO_PHELASTIC_ESINDEX":  "",
		"GO_PHELASTIC_ESUSER":   "",
		"GO_PHELASTIC_ESPASSWD": "",

		// Подключение к СУБД MongoDB
		"GO_PHELASTIC_MONGOHOST":   "",
		"GO_PHELASTIC_MONGOPORT":   "",
		"GO_PHELASTIC_MONGOUSER":   "",
		"GO_PHELASTIC_MONGOPASSWD": "",
		"GO_PHELASTIC_MONGONAMEDB": "",
	}

	getFileName := func(sf, confPath string, lfs []fs.DirEntry) (string, error) {
		for _, v := range lfs {
			if v.Name() == sf && !v.IsDir() {
				return path.Join(confPath, v.Name()), nil
			}
		}

		return "", fmt.Errorf("file '%s' is not found", sf)
	}

	setCommonSettings := func(fn string) error {
		viper.SetConfigFile(fn)
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		ls := Logs{}
		if ok := viper.IsSet("LOGGING"); ok {
			if err := viper.GetViper().Unmarshal(&ls); err != nil {
				return err
			}

			conf.CommonAppConfig.LogList = ls.Logging
		}

		z := ZabbixSet{}
		if ok := viper.IsSet("ZABBIX"); ok {
			if err := viper.GetViper().Unmarshal(&z); err != nil {
				return err
			}

			ti := 10
			if z.Zabbix.TimeInterval > 0 && z.Zabbix.TimeInterval <= 30 {
				ti = z.Zabbix.TimeInterval
			}

			hs := "0"
			if z.Zabbix.Handshake != "" && len(z.Zabbix.Handshake) <= 60 {
				hs = z.Zabbix.Handshake
			}

			np := 10051
			if z.Zabbix.NetworkPort != 0 && z.Zabbix.NetworkPort < 65536 {
				np = z.Zabbix.NetworkPort
			}

			conf.CommonAppConfig.Zabbix = ZabbixOptions{
				IsTransmit:   z.Zabbix.IsTransmit,
				TimeInterval: ti,
				NetworkPort:  np,
				NetworkHost:  z.Zabbix.NetworkHost,
				ZabbixHost:   z.Zabbix.ZabbixHost,
				ZabbixKey:    z.Zabbix.ZabbixKey,
				Handshake:    hs,
			}
		}

		return nil
	}

	setSpecial := func(fn string) error {
		viper.SetConfigFile(fn)
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		//Настройки для модуля подключения к NATS
		if viper.IsSet("NATS.host") {
			conf.AppConfigNATS.Host = viper.GetString("NATS.host")
		}
		if viper.IsSet("NATS.port") {
			conf.AppConfigNATS.Port = viper.GetInt("NATS.port")
		}

		// Настройки для модуля подключения к СУБД ElasticSearch
		if viper.IsSet("ElasticSearch.send") {
			conf.AppConfigElasticSearch.Send = viper.GetBool("ElasticSearch.send")
		}
		if viper.IsSet("ElasticSearch.host") {
			conf.AppConfigElasticSearch.Host = viper.GetString("ElasticSearch.host")
		}
		if viper.IsSet("ElasticSearch.port") {
			conf.AppConfigElasticSearch.Port = viper.GetInt("ElasticSearch.port")
		}
		if viper.IsSet("ElasticSearch.prefix") {
			conf.AppConfigElasticSearch.Prefix = viper.GetString("ElasticSearch.prefix")
		}
		if viper.IsSet("ElasticSearch.index") {
			conf.AppConfigElasticSearch.Index = viper.GetString("ElasticSearch.index")
		}
		if viper.IsSet("ElasticSearch.user") {
			conf.AppConfigElasticSearch.User = viper.GetString("ElasticSearch.user")
		}
		if viper.IsSet("ElasticSearch.passwd") {
			conf.AppConfigElasticSearch.Passwd = viper.GetString("ElasticSearch.passwd")
		}

		// Настройки для модуля подключения к СУБД MongoDB
		if viper.IsSet("MongoDB.host") {
			conf.AppConfigMongoDB.Host = viper.GetString("MongoDB.host")
		}
		if viper.IsSet("MongoDB.port") {
			conf.AppConfigMongoDB.Port = viper.GetInt("MongoDB.port")
		}
		if viper.IsSet("MongoDB.user") {
			conf.AppConfigMongoDB.User = viper.GetString("MongoDB.user")
		}
		if viper.IsSet("MongoDB.passwd") {
			conf.AppConfigMongoDB.Passwd = viper.GetString("MongoDB.passwd")
		}
		if viper.IsSet("MongoDB.namedb") {
			conf.AppConfigMongoDB.NameDB = viper.GetString("MongoDB.namedb")
		}

		return nil
	}

	for v := range envList {
		if env, ok := os.LookupEnv(v); ok {
			envList[v] = env
		}
	}

	rootPath, err := supportingfunctions.GetRootPath(rootDir)
	if err != nil {
		return conf, err
	}

	confPath := path.Join(rootPath, "configs")

	list, err := os.ReadDir(confPath)
	if err != nil {
		return conf, err
	}

	fileNameCommon, err := getFileName("config.yaml", confPath, list)
	if err != nil {
		return conf, err
	}

	//читаем общий конфигурационный файл
	if err := setCommonSettings(fileNameCommon); err != nil {
		return conf, err
	}

	var fn string
	if envList["GO_PHELASTIC_MAIN"] == "development" {
		fn, err = getFileName("config_dev.yaml", confPath, list)
		if err != nil {
			return conf, err
		}
	} else {
		fn, err = getFileName("config_prod.yaml", confPath, list)
		if err != nil {
			return conf, err
		}
	}

	if err := setSpecial(fn); err != nil {
		return conf, err
	}

	//Настройки для модуля подключения к NATS
	if envList["GO_PHELASTIC_NHOST"] != "" {
		conf.AppConfigNATS.Host = envList["GO_PHELASTIC_NHOST"]
	}
	if envList["GO_PHELASTIC_NPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHELASTIC_NPORT"]); err == nil {
			conf.AppConfigNATS.Port = p
		}
	}

	//Настройки для модуля подключения к СУБД ElasticSearch
	if envList["GO_PHELASTIC_ESSEND"] != "" {
		if envList["GO_PHELASTIC_ESSEND"] == "true" {
			conf.AppConfigElasticSearch.Send = true
		}

		if envList["GO_PHELASTIC_ESSEND"] == "false" {
			conf.AppConfigElasticSearch.Send = false
		}
	}
	if envList["GO_PHELASTIC_ESHOST"] != "" {
		conf.AppConfigElasticSearch.Host = envList["GO_PHELASTIC_ESHOST"]
	}
	if envList["GO_PHELASTIC_ESPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHELASTIC_ESPORT"]); err == nil {
			conf.AppConfigElasticSearch.Port = p
		}
	}
	if envList["GO_PHELASTIC_ESPREFIX"] != "" {
		conf.AppConfigElasticSearch.Prefix = envList["GO_PHELASTIC_ESPREFIX"]
	}
	if envList["GO_PHELASTIC_ESINDEX"] != "" {
		conf.AppConfigElasticSearch.Index = envList["GO_PHELASTIC_ESINDEX"]
	}
	if envList["GO_PHELASTIC_ESUSER"] != "" {
		conf.AppConfigElasticSearch.User = envList["GO_PHELASTIC_ESUSER"]
	}
	if envList["GO_PHELASTIC_ESPASSWD"] != "" {
		conf.AppConfigElasticSearch.Passwd = envList["GO_PHELASTIC_ESPASSWD"]
	}

	//Настройки для модуля подключения к СУБД MongoDB
	if envList["GO_PHELASTIC_MONGOHOST"] != "" {
		conf.AppConfigMongoDB.Host = envList["GO_PHELASTIC_MONGOHOST"]
	}
	if envList["GO_PHELASTIC_MONGOPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHELASTIC_MONGOPORT"]); err == nil {
			conf.AppConfigMongoDB.Port = p
		}
	}
	if envList["GO_PHELASTIC_MONGOUSER"] != "" {
		conf.AppConfigMongoDB.User = envList["GO_PHELASTIC_MONGOUSER"]
	}
	if envList["GO_PHELASTIC_MONGOPASSWD"] != "" {
		conf.AppConfigMongoDB.Passwd = envList["GO_PHELASTIC_MONGOPASSWD"]
	}
	if envList["GO_PHELASTIC_MONGONAMEDB"] != "" {
		conf.AppConfigMongoDB.NameDB = envList["GO_PHELASTIC_MONGONAMEDB"]
	}

	return conf, nil
}
