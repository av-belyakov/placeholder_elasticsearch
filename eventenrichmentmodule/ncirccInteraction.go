package eventenrichmentmodule

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// SettingsFuncFullNameOrganizationByINN настройки для получения информации
// об организации по её ИНН
type SettingsFuncFullNameOrganizationByINN struct {
	ctx               context.Context
	connectionTimeout time.Duration
	url               string
	token             string
	client            *http.Client
}

// ResponseFullNameOrganizationByINN результат поиска информации об организации по ИНН
type ResponseFullNameOrganizationByINN struct {
	Success bool                                        `json:"success"`
	Count   int                                         `json:"total"`
	Data    []SettingsResponseFullNameOrganizationByINN `json:"data"`
}

type SettingsResponseFullNameOrganizationByINN struct {
	SubjectINN string `json:"settings_inn_of_subject"`
	Name       string `json:"settings_name"`
	Sname      string `json:"settings_sname"`
	Type       string `json:"settings_subject_type"`
}

// NewSettingsFuncFullNameOrganizationByINN возвращает параметры настроек для запроса к НКЦКИ
// информации по организации по её ИНН
func NewSettingsFuncFullNameOrganizationByINN(ctx context.Context, url, token string, connTimeout time.Duration) (*SettingsFuncFullNameOrganizationByINN, error) {
	settings := SettingsFuncFullNameOrganizationByINN{
		ctx:               ctx,
		connectionTimeout: 30 * time.Second,
	}

	if url == "" {
		_, f, l, _ := runtime.Caller(0)
		return &settings, fmt.Errorf("the 'url' parameter must not be empty %s:%d", f, l-1)
	}

	if token == "" {
		_, f, l, _ := runtime.Caller(0)
		return &settings, fmt.Errorf("the 'token' parameter must not be empty %s:%d", f, l-1)
	}

	settings.url = url
	settings.token = token

	if connTimeout > (1 * time.Second) {
		settings.connectionTimeout = connTimeout
	}

	transport := &http.Transport{
		MaxIdleConns:        10,
		IdleConnTimeout:     connTimeout,
		MaxIdleConnsPerHost: 10,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            x509.NewCertPool(),
		},
	}

	settings.client = &http.Client{Transport: transport}

	return &settings, nil
}

// GetFullNameOrganizationByINN делает запрос к НКЦКИ с целью получения полной информации
// об организации по её ИНН
func (settings *SettingsFuncFullNameOrganizationByINN) GetFullNameOrganizationByINN(inn string) (ResponseFullNameOrganizationByINN, error) {
	var (
		rd ResponseFullNameOrganizationByINN = ResponseFullNameOrganizationByINN{}

		f string
		l int
	)

	req, err := http.NewRequestWithContext(settings.ctx, "GET", settings.url, strings.NewReader(""))
	if err != nil {
		_, f, l, _ = runtime.Caller(0)

		return rd, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	req.Header.Add("x-token", settings.token)

	q := req.URL.Query()
	q.Add("fields", "[\"settings_name\",\"settings_sname\",\"settings_inn_of_subject\",\"settings_subject_type\"]")
	q.Add("filter", fmt.Sprintf("[{\"property\":\"settings_inn_of_subject\",\"operator\":\"eq\",\"value\":\"%s\"}]", inn))
	q.Add("limit", "10")
	q.Add("start", "0")
	req.URL.RawQuery = q.Encode()

	res, err := settings.client.Do(req)
	if err != nil {
		_, f, l, _ = runtime.Caller(0)

		return rd, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	if res.StatusCode != http.StatusOK {
		_, f, l, _ = runtime.Caller(0)

		return rd, fmt.Errorf("error sending the request, response status is %s %s:%d", res.Status, f, l-1)
	}

	err = json.NewDecoder(res.Body).Decode(&rd)
	if err != nil {
		_, f, l, _ = runtime.Caller(0)

		return rd, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	return rd, nil
}
