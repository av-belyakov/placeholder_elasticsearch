package zabbixinteractions

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

type RequiredHostId struct {
	HostId           string
	zabbixConnection *ZabbixConnectionJsonRPC
}

type responseData struct {
	Error  map[string]interface{}   `json:"error"`
	Result []map[string]interface{} `json:"result"`
}

type FullSensorInformationFromZabbixAPI struct {
	SensorId   int
	HostId     string
	GeoCode    string
	ObjectArea string
	SubjectRF  string
	INN        string
	HomeNet    string
}

func GetFullSensorInformationFromZabbixAPI(ctx context.Context, sensorId int, zconn *ZabbixConnectionJsonRPC) (FullSensorInformationFromZabbixAPI, error) {
	fullInfo := FullSensorInformationFromZabbixAPI{SensorId: sensorId}

	var (
		requiredHostId *RequiredHostId
		err            error
	)

	/*
		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		тут надо внимательно посмотреть
		!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	*/
	requiredHostId, err = NewRequiredHostId(sensorId, zconn)
	if err != nil {
		tmp := ctx.Value("auth")
		if auth, ok := tmp.(struct {
			login, passwd string
		}); ok {
			//при возникновении ошибки пытаемся авторизоватся повторно, так как при устаревании
			//авторизационного хеша возможно появления ошибки с сообщением: 'Invalid params. Session terminated, re-login, please.'.
			authHash, err := authorizationZabbixAPI(auth.login, auth.passwd, *zconn)
			if err != nil {
				return fullInfo, err
			}

			zconn.authorizationHash = authHash
			requiredHostId, err = NewRequiredHostId(sensorId, zconn)
			if err != nil {
				return fullInfo, err
			}

		} else {
			return fullInfo, err
		}
	}
	fullInfo.HostId = requiredHostId.GetHostId()

	geoCode, err := requiredHostId.GetGeoCode()
	if err != nil {
		return fullInfo, err
	}
	fullInfo.GeoCode = geoCode

	objectArea, err := requiredHostId.GetObjectArea()
	if err != nil {
		return fullInfo, err
	}
	fullInfo.ObjectArea = objectArea

	subjectRF, err := requiredHostId.GetSubjectRF()
	if err != nil {
		return fullInfo, err
	}
	fullInfo.SubjectRF = subjectRF

	inn, err := requiredHostId.GetINN()
	if err != nil {
		return fullInfo, err
	}
	fullInfo.INN = inn

	homeNet, err := requiredHostId.GetHomeNet()
	if err != nil {
		return fullInfo, err
	}
	fullInfo.HomeNet = homeNet

	return fullInfo, nil
}

func NewRequiredHostId(sensorId int, zconn *ZabbixConnectionJsonRPC) (*RequiredHostId, error) {
	var (
		f string
		l int
	)

	requiredHostId := RequiredHostId{zabbixConnection: zconn}

	strReq := "{ \"jsonrpc\": \"2.0\","
	strReq += " \"method\": \"host.get\","
	strReq += " \"params\": {\"search\":"
	strReq += fmt.Sprintf("{\"host\": %d}},", sensorId)
	strReq += fmt.Sprintf(" \"auth\": \"%s+DDD\",", zconn.GetAuthorizationData())
	strReq += " \"id\": 1}"

	if sensorId == 0 {
		_, f, l, _ = runtime.Caller(0)
		return &requiredHostId, fmt.Errorf("error, the sensor ID cannot be equal to 0 %s:%d", f, l-1)
	}

	_, f, l, _ = runtime.Caller(0)
	res, err := zconn.SendPostRequest(strings.NewReader(strReq))
	if err != nil {
		return &requiredHostId, fmt.Errorf("error send post request, %v %s:%d", err, f, l+1)
	}

	rd := responseData{}
	err = json.NewDecoder(res).Decode(&rd)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		return &requiredHostId, fmt.Errorf("error decode request, %v %s:%d", err, f, l-2)
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

		return &requiredHostId, fmt.Errorf("error send post request, (%s %s) %s:%d", msg, data, f, l+1)
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

func (r *RequiredHostId) GetHostId() string {
	return r.HostId
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
