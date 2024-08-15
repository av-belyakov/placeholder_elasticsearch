package zabbixinteractions

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"
)

type ZabbixAuthorizationData struct {
	Id      int                    `json:"id"`
	JsonRPC string                 `json:"jsonrpc"`
	Result  string                 `json:"result"`
	Error   map[string]interface{} `json:"error"`
}

type ZabbixAuthorizationErrorMessage struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}

// NewZabbixConnectionJsonRPC создает объект соединения с Zabbix API с использование Json-RPC
func NewZabbixConnectionJsonRPC(settings SettingsZabbixConnectionJsonRPC) (*ZabbixConnectionJsonRPC, error) {
	var zc ZabbixConnectionJsonRPC

	connTimeout := 30 * time.Second
	if *settings.ConnectionTimeout > (1 * time.Second) {
		connTimeout = *settings.ConnectionTimeout
	}

	if settings.Host == "" {
		_, f, l, _ := runtime.Caller(0)
		return &zc, fmt.Errorf("the value 'Host' should not be empty %s:%d", f, l-2)
	}

	if settings.ConnectionTimeout == nil {
		t := time.Duration(5 * time.Second)
		settings.ConnectionTimeout = &t
	}

	client := &http.Client{Transport: &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     connTimeout,
		MaxIdleConnsPerHost: 10,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            x509.NewCertPool(),
		},
	}}

	zc = ZabbixConnectionJsonRPC{
		url:             fmt.Sprintf("https://%s/api_jsonrpc.php", settings.Host),
		host:            settings.Host,
		login:           settings.Login,
		passwd:          settings.Passwd,
		applicationType: "application/json-rpc",
		connClient:      client,
	}

	authHash, err := authorizationZabbixAPI(settings.Login, settings.Passwd, zc)
	if err != nil {
		return &zc, err
	}

	zc.authorizationHash = authHash

	return &zc, nil
}

// authorizationZabbixAPI делает запрос к Zabbix с целью получения хеша авторизации
// необходимого для дальнейшей работы с API
func authorizationZabbixAPI(login, passwd string, zc ZabbixConnectionJsonRPC) (string, error) {
	req := strings.NewReader(fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"user.login\",\"params\":{\"username\":\"%s\",\"password\":\"%s\"},\"id\":1}", login, passwd))
	res, err := zc.connClient.Post(zc.url, zc.applicationType, req)
	defer func() {
		if res == nil || res.Body == nil {
			return
		}

		res.Body.Close()
	}()
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return "", fmt.Errorf("error authorization, %v %s:%d", err, f, l-2)
	}

	if res.StatusCode != http.StatusOK {
		_, f, l, _ := runtime.Caller(0)
		return "", fmt.Errorf("error authorization, response status is %s %s:%d", res.Status, f, l-2)
	}

	result := ZabbixAuthorizationData{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return "", fmt.Errorf("error authorization, %v %s:%d", err, f, l-2)
	}

	if len(result.Error) > 0 {
		_, f, l, _ := runtime.Caller(0)
		var shortMsg, fullMsg string

		for k, v := range result.Error {
			if k == "message" {
				shortMsg = fmt.Sprint(v)
			}
			if k == "data" {
				fullMsg = fmt.Sprint(v)
			}
		}

		return "", fmt.Errorf("error authorization, (%s %s) %s:%d", shortMsg, fullMsg, f, l-1)
	}

	return result.Result, nil
}

// GetAuthorizationData возвращает авторизационный хеш
func (zc *ZabbixConnectionJsonRPC) GetAuthorizationData() string {
	return zc.authorizationHash
}

// SendPostRequest отправляет HTTP POST запрос с параметрами запроса вида JSON
func (zc *ZabbixConnectionJsonRPC) SendPostRequest(data *strings.Reader) (io.Reader, error) {
	var result io.Reader

	res, err := zc.connClient.Post(zc.url, zc.applicationType, data)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return result, fmt.Errorf("error sending the request, %v %s:%d", err, f, l-2)
	}

	if res.StatusCode != http.StatusOK {
		_, f, l, _ := runtime.Caller(0)
		return result, fmt.Errorf("error sending the request, response status is %s %s:%d", res.Status, f, l-1)
	}

	return res.Body, nil
}
