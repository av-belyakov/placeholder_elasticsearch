package zabbixinteractions

import (
	"context"
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
	JsonRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      int    `json:"id"`
}

// NewZabbixConnectionJsonRPC создает объект соединения с Zabbix API с
// использование Json-RPC
// ctx - должен быть context.WithCancel()
// settings - настройки
func NewZabbixConnectionJsonRPC(ctx context.Context, settings SettingsZabbixConnectionJsonRPC) (*ZabbixConnectionJsonRPC, error) {
	var zc ZabbixConnectionJsonRPC

	connTimeout := 30 * time.Second
	if *settings.ConnectionTimeout > (1 * time.Second) {
		connTimeout = *settings.ConnectionTimeout
	}

	transport := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: connTimeout,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            x509.NewCertPool(),
		},
	}

	client := &http.Client{Transport: transport}

	if settings.Host == "" {
		_, f, l, _ := runtime.Caller(0)
		return &zc, fmt.Errorf("the value 'Host' should not be empty %s:%d", f, l-2)
	}

	if settings.ConnectionTimeout == nil {
		t := time.Duration(5 * time.Second)
		settings.ConnectionTimeout = &t
	}

	zc = ZabbixConnectionJsonRPC{
		url:             fmt.Sprintf("https://%s/api_jsonrpc.php", settings.Host),
		host:            settings.Host,
		applicationType: "application/json-rpc",
		connClient:      client,
	}

	req := strings.NewReader(fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"user.login\",\"params\":{\"username\":\"%s\",\"password\":\"%s\"},\"id\":1}", settings.Login, settings.Passwd))
	res, err := client.Post(zc.url, zc.applicationType, req)
	defer func() {
		res.Body.Close()
	}()
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return &zc, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	if res.StatusCode != http.StatusOK {
		_, f, l, _ := runtime.Caller(0)
		return &zc, fmt.Errorf("error authorization, response status is %s %s:%d", res.Status, f, l-2)
	}

	authorizationData := ZabbixAuthorizationData{}
	if err = json.NewDecoder(res.Body).Decode(&authorizationData); err != nil {
		return &zc, err
	}

	zc.authorizationHash = authorizationData.Result

	return &zc, nil
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
		return result, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	if res.StatusCode != http.StatusOK {
		_, f, l, _ := runtime.Caller(0)
		return result, fmt.Errorf("error sending the request, response status is %s %s:%d", res.Status, f, l-1)
	}

	return res.Body, nil
}
