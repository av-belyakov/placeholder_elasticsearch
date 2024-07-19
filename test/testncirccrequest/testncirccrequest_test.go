package testncirccrequest_test

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

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type SettingsFuncFullNameOrganizationByINN struct {
	ctx               context.Context
	url               string
	token             string
	client            *http.Client
	connectionTimeout time.Duration
}

type ResponseFullNameOrganizationByINN struct {
	Data    []SettingsResponseFullNameOrganizationByINN `json:"data"`
	Success bool                                        `json:"success"`
	Count   int                                         `json:"total"`
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
		MaxIdleConns:    10,
		IdleConnTimeout: connTimeout,
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

var _ = Describe("Testncirccrequest", func() {
	Context("Тест 1. Проверка возможности получения подробной информации по ИНН", func() {
		It("При выполнении запроса не должно быть ошибок", func() {
			var (
				url   string = "https://lk.cert.local/api/v2/companies"
				token string = "fdd2c5e743960ec9ea80d1ff8868cc6d8439b02f4d61075efd69a46eaa52ff0e"
				inn   string = "7722377866"
			)

			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

			settingsFullOrgNameByINN, err := NewSettingsFuncFullNameOrganizationByINN(ctx, url, token, (5 * time.Second))
			Expect(err).ShouldNot(HaveOccurred())

			rd, err := settingsFullOrgNameByINN.GetFullNameOrganizationByINN(inn)
			Expect(err).ShouldNot(HaveOccurred())

			fmt.Println("RESULT:")
			for _, v := range rd.Data {
				fmt.Println("Organization name:", v.Name)
				fmt.Println("Full name organization:", v.Sname)
			}

			Expect(rd.Count).ShouldNot(Equal(0))

			/*

				Написал и успешно протестировал модули:
				1. Для выполнения запросов к Zabbix с целью получения подробной информации
				об сенсоре.
				2. Для запросов к НКЦКИ с целью получения информации о наозвании организации
				по ее ИНН.

				Теперь необходимо внедрить эти модули в основное припложение. Для этого
				необходимы следующие действия:
				1. Дописать и протестировать обработку конфигурационных файлов с учетом
				добавления туда конфигурационных параметров, необходимых для работы выше
				указанных модулей.
				2. Продумать модуль временного хранения информации (который будет использоватся
				при недоступности основных источников информации).
				3. Продумать систему попыток переподключения к Zabbix при его недоступности,
				так как при запросе авторизации Zabbix выдает токен, который может устареть
				если Zabbix был не доступен какое то время (возможно он перезапускался).

			*/
		})
	})
})
