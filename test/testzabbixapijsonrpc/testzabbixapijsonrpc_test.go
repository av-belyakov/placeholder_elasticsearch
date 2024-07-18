package testzabbixapijsonrpc_test

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"placeholder_elasticsearch/zabbixinteractions"
)

var _ = Describe("Testzabbixapijsonrpc", Ordered, func() {
	var (
		ctx context.Context
		//ctxCancel context.CancelFunc

		zabbixConnHandler *zabbixinteractions.ZabbixConnectionJsonRPC
		zabbixConnErr     error
	)

	BeforeAll(func() {
		connTimeout := time.Duration(3 * time.Second)

		ctx, _ /*ctxCancel*/ = context.WithCancel(context.Background())
		zabbixConnHandler, zabbixConnErr = zabbixinteractions.NewZabbixConnectionJsonRPC(
			ctx,
			zabbixinteractions.SettingsZabbixConnectionJsonRPC{
				Host:   "192.168.9.45", //правильный
				Login:  "Cherry",
				Passwd: "v-2ymX!aVg3eS*hC",
				//Host:              "192.168.9.145", //не правильный
				ConnectionTimeout: &connTimeout,
			})
	})

	Context("Тест 0. Выполняем создание нового JSON RPC соединения", func() {
		It("При инициализации нового JSON RPS соединения ошибок быть не должно", func() {
			fmt.Println("RESULT:", zabbixConnHandler)

			Expect(zabbixConnErr).ShouldNot(HaveOccurred())
		})
	})

	Context("Тест 1. Выполняем POST запрос к Zabbix", func() {
		It("Не должно быть ошибок при запросе", func() {
			//hostName := 690023
			hostName := 8030160

			requiredHostId, err := NewRequiredHostId(hostName, zabbixConnHandler)
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println("_____ HostId _____:", requiredHostId.HostId)

			geoCode, err := requiredHostId.GetGeoCode()
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println("_____ GeoCode _____:", geoCode)

			objectArea, err := requiredHostId.GetObjectArea()
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println("_____ ObjectArea _____:", objectArea)

			subjectRF, err := requiredHostId.GetSubjectRF()
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println("_____ SubjectRF _____:", subjectRF)

			inn, err := requiredHostId.GetINN()
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println("_____ INN _____:", inn)

			homeNet, err := requiredHostId.GetHomeNet()
			Expect(err).ShouldNot(HaveOccurred())
			fmt.Println("_____ HomeNet _____:", homeNet)

			Expect(true).Should(BeTrue())
		})
	})
})

type RequiredHostId struct {
	HostId           string
	zabbixConnection *zabbixinteractions.ZabbixConnectionJsonRPC
}

type responseData struct {
	Error  map[string]interface{}   `json:"error"`
	Result []map[string]interface{} `json:"result"`
}

func NewRequiredHostId(sensorId int, zconn *zabbixinteractions.ZabbixConnectionJsonRPC) (*RequiredHostId, error) {
	requiredHostId := RequiredHostId{zabbixConnection: zconn}

	strReq := "{ \"jsonrpc\": \"2.0\","
	strReq += " \"method\": \"host.get\","
	strReq += " \"params\": {\"search\":"
	strReq += fmt.Sprintf("{\"host\": %d}},", sensorId)
	strReq += fmt.Sprintf(" \"auth\": \"%s\",", zconn.GetAuthorizationData())
	strReq += " \"id\": 1}"

	var (
		f string
		l int
	)

	res, err := zconn.SendPostRequest(strings.NewReader(strReq))
	if err != nil {
		_, f, l, _ = runtime.Caller(0)
		return &requiredHostId, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	rd := responseData{}
	err = json.NewDecoder(res).Decode(&rd)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return &requiredHostId, fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	if len(rd.Error) > 0 {
		var msg, data string

		for k, v := range rd.Error {
			if k == "message" {
				msg = fmt.Sprint(v)
			}

			if k == "data" {
				data = fmt.Sprint(v)
			}
		}

		return &requiredHostId, fmt.Errorf("%s. %s %s:%d", msg, data, f, l-2)
	}

DONE:
	for _, v := range rd.Result {
		for key, value := range v {
			if key == "hostid" {
				requiredHostId.HostId = fmt.Sprint(value)

				break DONE
			}
		}
	}

	return &requiredHostId, nil
}

func (r *RequiredHostId) GetGeoCode() (string, error) {
	//geo_code
	strReq := "{ \"jsonrpc\": \"2.0\","
	strReq += " \"method\": \"item.get\","
	strReq += " \"params\": {"
	strReq += " \"output\": \"extend\","
	strReq += fmt.Sprintf(" \"hostids\": \"%s\",", r.HostId)
	strReq += " \"search\": {\"key_\": \"geo_code\"},"
	strReq += " \"sortfield\": \"name\"},"
	strReq += fmt.Sprintf(" \"auth\": \"%s\",", r.zabbixConnection.GetAuthorizationData())
	strReq += " \"id\": 1}"

	return r.sendRequest(strReq)
}

func (r *RequiredHostId) GetObjectArea() (string, error) {
	//object_area
	strReq := "{ \"jsonrpc\": \"2.0\","
	strReq += " \"method\": \"item.get\","
	strReq += " \"params\": {"
	strReq += " \"output\": \"extend\","
	strReq += fmt.Sprintf(" \"hostids\": \"%s\",", r.HostId)
	strReq += " \"search\": {\"key_\": \"object_area\"},"
	strReq += " \"sortfield\": \"name\"},"
	strReq += fmt.Sprintf(" \"auth\": \"%s\",", r.zabbixConnection.GetAuthorizationData())
	strReq += " \"id\": 1}"

	return r.sendRequest(strReq)
}

func (r *RequiredHostId) GetSubjectRF() (string, error) {
	//subject_RF
	strReq := "{ \"jsonrpc\": \"2.0\","
	strReq += " \"method\": \"item.get\","
	strReq += " \"params\": {"
	strReq += " \"output\": \"extend\","
	strReq += fmt.Sprintf(" \"hostids\": \"%s\",", r.HostId)
	strReq += " \"search\": {\"key_\": \"subject_RF\"},"
	strReq += " \"sortfield\": \"name\"},"
	strReq += fmt.Sprintf(" \"auth\": \"%s\",", r.zabbixConnection.GetAuthorizationData())
	strReq += " \"id\": 1}"

	return r.sendRequest(strReq)
}

func (r *RequiredHostId) GetINN() (string, error) {
	//inn
	strReq := "{ \"jsonrpc\": \"2.0\","
	strReq += " \"method\": \"item.get\","
	strReq += " \"params\": {"
	strReq += " \"output\": \"extend\","
	strReq += fmt.Sprintf(" \"hostids\": \"%s\",", r.HostId)
	strReq += " \"search\": {\"key_\": \"inn\"},"
	strReq += " \"sortfield\": \"name\"},"
	strReq += fmt.Sprintf(" \"auth\": \"%s\",", r.zabbixConnection.GetAuthorizationData())
	strReq += " \"id\": 1}"

	return r.sendRequest(strReq)
}

func (r *RequiredHostId) GetHomeNet() (string, error) {
	//home_net
	strReq := "{ \"jsonrpc\": \"2.0\","
	strReq += " \"method\": \"item.get\","
	strReq += " \"params\": {"
	strReq += " \"output\": \"extend\","
	strReq += fmt.Sprintf(" \"hostids\": \"%s\",", r.HostId)
	strReq += " \"search\": {\"key_\": \"home_net\"},"
	strReq += " \"sortfield\": \"name\"},"
	strReq += fmt.Sprintf(" \"auth\": \"%s\",", r.zabbixConnection.GetAuthorizationData())
	strReq += " \"id\": 1}"

	return r.sendRequest(strReq)
}

func (r *RequiredHostId) sendRequest(str string) (string, error) {
	var (
		f string
		l int
	)

	res, err := r.zabbixConnection.SendPostRequest(strings.NewReader(str))
	if err != nil {
		_, f, l, _ = runtime.Caller(0)
		return "", fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	rd := responseData{}
	err = json.NewDecoder(res).Decode(&rd)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return "", fmt.Errorf("%v %s:%d", err, f, l-2)
	}

	if len(rd.Error) > 0 {
		var msg, data string

		for k, v := range rd.Error {
			if k == "message" {
				msg = fmt.Sprint(v)
			}

			if k == "data" {
				data = fmt.Sprint(v)
			}
		}

		return "", fmt.Errorf("%s. %s %s:%d", msg, data, f, l-2)
	}

	for _, v := range rd.Result {
		for key, value := range v {
			if key == "lastvalue" {
				return fmt.Sprint(value), nil
			}
		}
	}

	return "", nil
}
