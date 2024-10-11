package eventenrichmentmodule

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// GeoIpClient GeoIP клиента для запроса информации из БД GeoIP компании
type GeoIpClient struct {
	port              int
	host              string
	path              string
	ctx               context.Context
	connectionTimeout time.Duration
	client            *http.Client
}

type resultGeoIP struct {
	AddressVersion string          `json:"address_version"`
	IpLocations    []ipLocationSet `json:"ip_locations"`
}

type ipLocationSet struct {
	Source string `json:"source"`
	IpLocation
}

// GeoIpInformation список найденной информации по запрашиваемому ip адресу
type GeoIpInformation struct {
	IsSuccess bool
	Ip        string
	Info      map[string]IpLocation
}

// IpLocation подробная информация об ip адресе
type IpLocation struct {
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

// geoIpClientOptions функциональные параметры
type geoIpClientOptions func(*GeoIpClient) error

// WithPort устанавливает порт для взаимодействия с модулем
func WithPort(v int) geoIpClientOptions {
	return func(gic *GeoIpClient) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		gic.port = v

		return nil
	}
}

// WithHost устанавливает хост для взаимодействия с модулем
func WithHost(v string) geoIpClientOptions {
	return func(gic *GeoIpClient) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		gic.host = v

		return nil
	}
}

// WithPath устанавливает путь запроса по которой осуществляется маршрутизация
func WithPath(v string) geoIpClientOptions {
	return func(gic *GeoIpClient) error {
		gic.path = v

		return nil
	}
}

// WithConnectionTimeout устанавливает время ожидания выполнения запроса
func WithConnectionTimeout(timeout time.Duration) geoIpClientOptions {
	return func(gic *GeoIpClient) error {
		if timeout > (1 * time.Second) {
			gic.connectionTimeout = timeout
		}

		return nil
	}
}

// NewGeoIpClient GeoIP клиент
func NewGeoIpClient(ctx context.Context, opts ...geoIpClientOptions) (*GeoIpClient, error) {
	settings := GeoIpClient{
		ctx:               ctx,
		connectionTimeout: 30 * time.Second,
	}

	for _, opt := range opts {
		if err := opt(&settings); err != nil {
			return &settings, err
		}
	}

	settings.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     settings.connectionTimeout,
			MaxIdleConnsPerHost: 10,
		}}

	return &settings, nil
}

// GetGeoInformation делает запрос к API БД GeoIP
func (gic *GeoIpClient) GetGeoInformation(ip string) (GeoIpInformation, error) {
	result := GeoIpInformation{
		Ip:   ip,
		Info: make(map[string]IpLocation, 0),
	}

	rex := regexp.MustCompile(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)[.]){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`)
	if !rex.MatchString(ip) {
		_, f, l, _ := runtime.Caller(0)
		return result, fmt.Errorf("an invalid ip address '%s' was received %s:%d", ip, f, l-1)
	}

	url := fmt.Sprintf("http://%s:%d/%s/%s/", gic.host, gic.port, gic.path, ip)
	req, err := http.NewRequestWithContext(gic.ctx, "GET", url, strings.NewReader(""))
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return result, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	res, err := gic.client.Do(req)
	defer responseClose(res)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return result, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	if res.StatusCode != http.StatusOK {
		_, f, l, _ := runtime.Caller(0)
		return result, fmt.Errorf("error sending the request, response status is %s %s:%d", res.Status, f, l-1)
	}

	resultGeoIP := resultGeoIP{}
	err = json.NewDecoder(res.Body).Decode(&resultGeoIP)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return result, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	result.IsSuccess = true
	for _, v := range resultGeoIP.IpLocations {
		result.Info[v.Source] = IpLocation{
			City:        v.City,
			Country:     v.Country,
			CountryCode: v.CountryCode,
		}
	}

	return result, nil
}
