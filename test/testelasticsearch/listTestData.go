package testelasticsearch

import (
	"placeholder_elasticsearch/datamodels"
	"placeholder_elasticsearch/datamodels/commonalert"
	"placeholder_elasticsearch/datamodels/commonalertartifact"
	commonevent "placeholder_elasticsearch/datamodels/commonevent"
	commonobjectevent "placeholder_elasticsearch/datamodels/commonobjectevent"
)

//*******************************************
//		тестовые данные для формировании
//	индексов содержащих Events
//*******************************************

//*******************************************
//		тестовые данные для формировании
//	индексов содержащих Alerts
//*******************************************

var EventForEsAlertTestOne datamodels.EventMessageForEsAlert = datamodels.EventMessageForEsAlert{
	CommonEventType: commonevent.CommonEventType{
		Base:           false,
		StartDate:      "2024-02-06T15:20:41+03:00",
		RootId:         "~84625227848",
		ObjectId:       "~4192",
		ObjectType:     "alert",
		Organisation:   "GCM",
		OrganisationId: "~4192",
		Operation:      "new",
		RequestId:      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
	},
	Details: datamodels.EventMessageForEsAlertDetails{
		SourceRef:   "TSK-8MSK-6-ZPM-240206-1215999",
		Title:       "222",
		Description: "111",
		Tags: map[string][]string{
			"ats:reason": {"INFO Controlled FGS"},
			"sensor:id":  {"8030066"},
		},
		TagsAll: []string{
			"ATs:reason=\"INFO Controlled FGS\"",
			"Sensor:id=\"8030066\"",
		},
	},
	Object: datamodels.EventMessageForEsAlertObject{
		CommonEventAlertObject: commonobjectevent.CommonEventAlertObject{
			Tlp:             1,
			UnderliningId:   "~85455464790",
			Id:              "~85455464790",
			CreatedBy:       "ddddd",
			CreatedAt:       "1970-01-01T03:00:00+03:00",
			UpdatedAt:       "1970-01-01T03:00:00+03:00",
			UnderliningType: "aalllert",
			Title:           "vbb er3",
			Description:     "any more",
			Status:          "None",
			Date:            "2024-02-06T15:37:52+03:00",
			Type:            "snort_alert",
			ObjectType:      "",
			Source:          "zsiеmSystems",
			SourceRef:       "TSK-8MSK-6-ZPM-240206-1215999",
			Case:            "alert",
			CaseTemplate:    "alert_snort",
		},
		Tags: map[string][]string{
			"sensor:id": {"8030012"},
		},
		TagsAll: []string{
			"'Sensor:id=\"8030012\"'",
			"'Webhook:send=ES'",
		},
		CustomFields: datamodels.CustomFields{
			"first-time": &datamodels.CustomFieldDateType{
				Order: 0,
				Date:  "2024-02-06T15:20:30+03:00",
			},
			"last-time": &datamodels.CustomFieldDateType{
				Order: 0,
				Date:  "2024-02-06T15:20:30+03:00",
			},
		},
	},
}

var AlertForEsAlertTestOne datamodels.AlertMessageForEsAlert = datamodels.AlertMessageForEsAlert{
	CommonAlertType: commonalert.CommonAlertType{
		Tlp:       2,
		Date:      "1970-01-01T03:00:00+03:00",
		CreatedAt: "2024-02-07T11:11:11+03:00",
		// UpdatedAt: ,
		UpdatedBy:       "webhook@cloud.gcm",
		UnderliningId:   "~88026357960",
		Status:          "New",
		Type:            "snort",
		UnderliningType: "__Snort",
		Description:     "free alerts",
		CaseTemplate:    "sonr",
		SourceRef:       "TSK-8MSK-6-ZPM-240206-1216137",
	},
	Tags: map[string][]string{
		"sensor:id":  {"8030105"},
		"ats:reason": {"Редко встречающиеся признаки ВПО"},
	},
	TagsAll: []string{
		"Sensor:id=\"8030105\"",
		"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
		"'Webhook:send=ES'",
	},
	CustomFields: datamodels.CustomFields{
		"first-time": &datamodels.CustomFieldDateType{
			Order: 0,
			Date:  "2024-01-01T05:22:30+03:00",
		},
		"last-time": &datamodels.CustomFieldDateType{
			Order: 0,
			Date:  "2024-01-17T00:18:13+03:00",
		},
	},
	Artifacts: map[string][]datamodels.ArtifactForEsAlert{
		"coordinates": {
			{
				CommonArtifactType: commonalertartifact.CommonArtifactType{
					Ioc:           false,
					Tlp:           1,
					UnderliningId: "~84302220012",
					Id:            "~84302220012",
					CreatedAt:     "2024-01-26T13:02:01+03:00",
					//UpdatedAt: ,
					StartDate: "2024-01-26T13:02:01+03:00",
					CreatedBy: "friman@email.net",
					Data:      "63.5656 89.12",
					DataType:  "coordinates",
					Message:   "Any message",
				},
				Tags: map[string][]string{
					"sensor:id":     {"1111111"},
					"geoip:country": {"CH"},
				},
				TagsAll: []string{
					"Sensor:id=\"1111111\"",
					"geoip:country=CH",
					"'Webhook:send=ES'",
				},
			},
		},
		"ipaddr": {
			{
				CommonArtifactType: commonalertartifact.CommonArtifactType{
					Ioc:           true,
					Tlp:           2,
					UnderliningId: "~306522241",
					Id:            "~306522241",
					CreatedAt:     "2024-01-16T03:32:01+03:00",
					StartDate:     "2024-01-04T19:32:01+03:00",
					CreatedBy:     "example@email.net",
					Data:          "5.63.123.99",
					DataType:      "ipaddr",
					Message:       "ffdffd fdg",
				},
				Tags: map[string][]string{
					"sensor:id":     {"3411"},
					"geoip:country": {"RU"},
				},
				TagsAll: []string{
					"Sensor:id=\"3411\"",
					"geoip:country=RU",
					"'Webhook:send=ES'",
				},
			},
		},
	},
}

var EventForEsAlertTestTwo datamodels.EventMessageForEsAlert = datamodels.EventMessageForEsAlert{
	CommonEventType: commonevent.CommonEventType{
		Base:           true,                        //замена
		StartDate:      "2024-02-13T05:12:24+03:00", //замена
		RootId:         "~84625227848",
		ObjectId:       "~4192",
		ObjectType:     "ALERT",   //замена
		Organisation:   "GCM-MSK", //замена
		OrganisationId: "~419211", //замена
		Operation:      "update",  //замена
		RequestId:      "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497",
	},
	Details: datamodels.EventMessageForEsAlertDetails{
		SourceRef:   "TSK-8MSK-6-ZPM-240206-1215999",
		Title:       "протоколы: **smtp/tcp**",            //замена
		Description: "использует протоколы: **smtp/tcp**", //замена
		//замена
		Tags: map[string][]string{
			"ats:reason": {
				"INFO Controlled FGS",
				"Редко встречающиеся признаки ВПО",
			},
			"sensor:id":             {"8030066"},
			"ats:geoip":             {"Китай"},
			"misp:Network activity": {"snort"},
		},
		TagsAll: []string{
			"ATs:reason=\"INFO Controlled FGS\"",
			"Sensor:id=\"8030066\"",
			"'APPA:Direction=\"inbound\"'",
			"ATs:geoip=\"Китай\"",
			"misp:Network activity=\"snort\"",
			"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
		},
	},
	Object: datamodels.EventMessageForEsAlertObject{
		CommonEventAlertObject: commonobjectevent.CommonEventAlertObject{
			Tlp:             1,
			UnderliningId:   "~85455464790",
			Id:              "~85455464790",
			CreatedBy:       "d.zablotsky@cloud.gcm",       //замена
			CreatedAt:       "2024-02-10T23:25:14+03:00",   //замена
			UpdatedAt:       "2024-02-06T15:15:14+03:00",   //замена
			UnderliningType: "ALERT",                       //замена
			Title:           "Атака направлена **внутрь**", //замена
			Description:     "Вирусное заражение",          //замена
			Status:          "None",
			Date:            "2024-02-06T15:37:52+03:00",
			Type:            "snort_alert",
			ObjectType:      "",
			Source:          "zsiеmSystems",
			SourceRef:       "TSK-8MSK-6-ZPM-240206-1215999",
			Case:            "alert",
			CaseTemplate:    "Alert_Snort", //замена
		},
		//замена
		Tags: map[string][]string{
			"sensor:id":             {"8030105"},
			"misp:Payload delivery": {"email-src"},
		},
		TagsAll: []string{
			"'Webhook:send=ES'",
			"'Sensor:id=\"8030105\"'",
			"misp:Payload delivery=\"email-src\"",
		},
		CustomFields: datamodels.CustomFields{
			"first-time": &datamodels.CustomFieldDateType{
				Order: 0,
				Date:  "2024-02-06T15:20:30+03:00",
			},
			//замена
			"last-time": &datamodels.CustomFieldDateType{
				Order: 0,
				Date:  "2024-02-07T22:48:13+03:00",
			},
		},
	},
}

var AlertForEsAlertTestTwo datamodels.AlertMessageForEsAlert = datamodels.AlertMessageForEsAlert{
	CommonAlertType: commonalert.CommonAlertType{
		Tlp:             3, //замена
		Date:            "1970-01-01T03:00:00+03:00",
		CreatedAt:       "2024-02-10T10:00:41+03:00", //замена
		UpdatedAt:       "2024-02-11T12:34:48+03:00", //замена
		UpdatedBy:       "webexample@cloud.gcm",      //замена
		UnderliningId:   "~88026357960",
		Status:          "Update",       //замена
		Type:            "snort_alert",  //замена
		UnderliningType: "snort_alert",  //замена
		Description:     "free alerts!", //замена
		CaseTemplate:    "snort",        //замена
		SourceRef:       "TSK-8MSK-6-ZPM-240206-1216137",
	},
	//замена
	Tags: map[string][]string{
		"ats:reason": {"Редко встречающиеся признаки ВПО"},
	},
	//замена
	TagsAll: []string{
		"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
		"'Webhook:send=ES'",
		"APPA:Direction=\"inbound\"",
	},
	CustomFields: datamodels.CustomFields{
		//замена
		"first-time": &datamodels.CustomFieldDateType{
			Order: 0,
			Date:  "2024-01-22T15:13:10+03:00",
		},
		"last-time": &datamodels.CustomFieldDateType{
			Order: 0,
			Date:  "2024-01-17T00:18:13+03:00",
		},
	},
	Artifacts: map[string][]datamodels.ArtifactForEsAlert{
		"coordinates": {
			{
				CommonArtifactType: commonalertartifact.CommonArtifactType{
					Ioc:           true, //замена
					Tlp:           3,    //замена
					UnderliningId: "",   //НЕ замена
					Id:            "~84302220012",
					CreatedAt:     "2024-01-27T22:17:17+03:00", //замена
					StartDate:     "2024-01-26T13:02:01+03:00",
					CreatedBy:     "friman@email.net",
					Data:          "63.5656 89.1211", //замена
					DataType:      "coordinates",
					Message:       "Any message",
				},
				//добавление
				Tags: map[string][]string{
					"sensor:id":     {"12345667"},
					"geoip:country": {"CH", "US"},
				},
				//добавление
				TagsAll: []string{
					"Sensor:id=\"12345667\"",
					"geoip:country=CH",
					"geoip:country=US",
					"'Webhook:send=ES'",
				},
			},
			//добавление
			{
				CommonArtifactType: commonalertartifact.CommonArtifactType{
					Ioc:           false,
					Tlp:           1,
					UnderliningId: "",
					Id:            "~8430120011",
					CreatedAt:     "2024-01-27T22:17:17+03:00",
					StartDate:     "2024-01-26T13:02:01+03:00",
					CreatedBy:     "friman@email.net",
					Data:          "89.12 11.53",
					DataType:      "coordinates",
					Message:       "funy description",
				},
				Tags: map[string][]string{
					"sensor:id":     {"43522"},
					"geoip:country": {"RU"},
				},
				TagsAll: []string{
					"Sensor:id=\"12345667\"",
					"geoip:country=RU",
					"'Webhook:send=ES'",
				},
			},
		},
		//добавление
		"ip_home": {
			{
				CommonArtifactType: commonalertartifact.CommonArtifactType{
					Ioc:           false,
					Tlp:           1,
					UnderliningId: "~7344456683",
					Id:            "~7344456683",
					CreatedAt:     "2024-01-17T13:12:01+03:00",
					StartDate:     "2024-01-04T19:32:01+03:00",
					CreatedBy:     "example@email.net",
					Data:          "5.63.123.99",
					DataType:      "ip_home",
					Message:       "ffdffd fdg",
				},
				Tags: map[string][]string{
					"geoip:country": {"RU"},
				},
				TagsAll: []string{
					"geoip:country=RU",
					"'Webhook:send=ES'",
				},
			},
		},
		"ipaddr": {
			//добавление
			{
				CommonArtifactType: commonalertartifact.CommonArtifactType{
					Ioc:           true,
					Tlp:           3,
					UnderliningId: "~502221144",
					Id:            "~502221144",
					CreatedAt:     "2024-02-06T13:12:01+03:00",
					StartDate:     "2024-01-04T19:32:01+03:00",
					CreatedBy:     "FOOOE@email.net",
					Data:          "89.6.33.41",
					DataType:      "ipaddr",
					Message:       "fast message",
				},
				Tags: map[string][]string{
					"sensor:id":     {"8999"},
					"geoip:country": {"KZ"},
				},
				TagsAll: []string{
					"Sensor:id=\"8999\"",
					"geoip:country=KZ",
					"'Webhook:send=ES'",
				},
			},
		},
	},
}
