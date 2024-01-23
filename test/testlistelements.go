package testhandlerobservables

func GetReportsList() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "reports.CyberCrime-Tracker_1_0.taxonomies.level", ElemValue: "info"},
		{ElemName: "reports.CyberCrime-Tracker_1_0.taxonomies.namespace", ElemValue: "CCT"},
		{ElemName: "reports.CyberCrime-Tracker_1_0.taxonomies.predicate", ElemValue: "C2 Search"},
		{ElemName: "reports.CyberCrime-Tracker_1_0.taxonomies.value", ElemValue: "0 hits"},

		{ElemName: "reports.CyberCrime-Tracker_1_0.taxonomies.level", ElemValue: "info"},
		{ElemName: "reports.CyberCrime-Tracker_1_0.taxonomies.namespace", ElemValue: "HTy"},
		{ElemName: "reports.CyberCrime-Tracker_1_0.taxonomies.predicate", ElemValue: "No search out"},
		{ElemName: "reports.CyberCrime-Tracker_1_0.taxonomies.value", ElemValue: "78 bit"},

		{ElemName: "reports.DShield_lookup_1_0.taxonomies.level", ElemValue: "not info"},
		{ElemName: "reports.DShield_lookup_1_0.taxonomies.namespace", ElemValue: "Jonson A"},
		{ElemName: "reports.DShield_lookup_1_0.taxonomies.predicate", ElemValue: "Tyo"},
		{ElemName: "reports.DShield_lookup_1_0.taxonomies.value", ElemValue: "1000"},

		{ElemName: "reports.URLhaus_2_0.taxonomies.level", ElemValue: "info"},
		{ElemName: "reports.URLhaus_2_0.taxonomies.namespace", ElemValue: "Lochkarev"},
		{ElemName: "reports.URLhaus_2_0.taxonomies.predicate", ElemValue: "CCNoK"},
		{ElemName: "reports.URLhaus_2_0.taxonomies.value", ElemValue: "8bit"},

		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.level", ElemValue: "warning"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", ElemValue: "LPPPT1"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", ElemValue: "12-34"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.value", ElemValue: "NONE"},

		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.level", ElemValue: "suspicious"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", ElemValue: "urlscan.io"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", ElemValue: "Search out"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.value", ElemValue: "6 results"},

		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.level", ElemValue: "safe"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", ElemValue: "Maltiverse"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", ElemValue: "Report"},
		{ElemName: "reports.Urlscan_io_Search_0_1_1.taxonomies.value", ElemValue: "6 neutral"},
	}
}

func GetObservableOne() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "observables._id", ElemValue: "~3460985064"},
		{ElemName: "observables.data", ElemValue: "9608643"},
		{ElemName: "observables.dataType", ElemValue: "snort_sid"},
		{ElemName: "observables.ioc", ElemValue: true},
		{ElemName: "observables._type", ElemValue: "Observable"},
		{ElemName: "observables.tlp", ElemValue: float64(2)},
		{ElemName: "observables._createdAt", ElemValue: float64(1690968664227)},
		{ElemName: "observables._createdBy", ElemValue: "uds@crimea-rcm"},
		{ElemName: "observables._updatedAt", ElemValue: float64(1704977151860)},
		{ElemName: "observables.sighted", ElemValue: false},
		{ElemName: "observables.startDate", ElemValue: float64(1690968664227)},
		{ElemName: "observables.tags", ElemValue: []interface{}{
			"misp:Network activity=\"attachment\"",
			"b2m:dumpfile=\"main\"",
		}},

		{ElemName: "observables.attachment.contentType", ElemValue: "text/plain"},
		{ElemName: "observables.attachment.id", ElemValue: "c29438b04791184d3eba39bdb7cf99560ab62068fee9509d50cf59723c398ac1"},
		{ElemName: "observables.attachment.name", ElemValue: "n[n.txt"},
		{ElemName: "observables.attachment.size", ElemValue: float64(817)},
		{ElemName: "observables.attachment.hashes", ElemValue: []interface{}{
			"c29438b04791184d3eba39bdb7cf99560ab62068fee9509d50cf59723c398ac1",
			"58861ef4c118cc3270b9871734ee54852a1374e5",
			"7c531394dc2f483bc6c6c628c02e0788",
		}},
		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.", ElemValue: ""},
	}
}

func GetEventOne() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "event.objectId", ElemValue: "~419385432"},
		{ElemName: "event.objectType", ElemValue: "case"},
		{ElemName: "event.base", ElemValue: true},
		{ElemName: "event.startDate", ElemValue: float64(1705061267325)},
		{ElemName: "event.organisation", ElemValue: "RCM"},
		{ElemName: "event.organisationId", ElemValue: "~20488"},
		{ElemName: "event.operation", ElemValue: "update"},
		{ElemName: "event.rootId", ElemValue: "~419385432"},
		{ElemName: "event.requestId", ElemValue: "019f0dbc0ab90bbe:-58339429:18b66b86afa:-8000:780802"},

		//----------- object
		{ElemName: "event.object.id", ElemValue: "~85771464712"},
		{ElemName: "event.object.tags", ElemValue: []interface{}{
			"'Webhook:send=ES'",
			"ATs:reason=\"bmpol WhiteList\"",
			"Sensor:id=\"8030071\"",
			"ATs:geoip=\"Нидерланды\"",
			"ATs:reason=\"Зафиксированы признаки взаимодействия ВПО класса trojan\"",
			"ATs:reason=\"bmpol List 4\"",
			"ATs:reason=\"ПНПО\"",
			"APPA:Direction=\"outbound\"",
		}},
		{ElemName: "event.object.summary", ElemValue: "trigger"},
		{ElemName: "event.object.owner", ElemValue: "b.polyakov@cloud.gcm"},
		{ElemName: "event.object.updatedBy", ElemValue: "d.sergeev@cloud.gcm"},
		{ElemName: "event.object.title", ElemValue: "ПНПО \"Ammyy Admin\""},
		{ElemName: "event.object.severity", ElemValue: float64(2)},
		{ElemName: "event.object.endDate", ElemValue: float64(0)},
		{ElemName: "event.object.caseId", ElemValue: float64(34411)},
		{ElemName: "event.object.description", ElemValue: "Атака направлена **наружу**"},
		{ElemName: "event.object.flag", ElemValue: true},
		{ElemName: "event.object.tlp", ElemValue: float64(3)},
		{ElemName: "event.object.pap", ElemValue: float64(5)},
		//object.customFields
		{
			ElemName:  "event.object.customFields.ncircc-class-attack.order",
			ElemValue: float64(3),
		},
		{
			ElemName:  "event.object.customFields.class-attack.order",
			ElemValue: float64(2),
		},
		{
			ElemName:  "event.object.customFields.first-time.order",
			ElemValue: float64(10),
		},
		{
			ElemName:  "event.object.customFields.first-time.date",
			ElemValue: float64(1705052465000),
		},
		{
			ElemName:  "event.object.customFields.last-time.order",
			ElemValue: float64(221),
		},
		{
			ElemName:  "event.object.customFields.last-time.date",
			ElemValue: float64(1705052479000),
		},
		{
			ElemName:  "event.object.customFields.sphere.order",
			ElemValue: float64(4),
		},
		{
			ElemName:  "event.object.customFields.sphere.string",
			ElemValue: "Здравоохранение",
		},
		{
			ElemName:  "event.object.customFields.state.order",
			ElemValue: float64(17),
		},
		{
			ElemName:  "event.object.customFields.state.string",
			ElemValue: "Город федерального значения Севастополь",
		},
		{
			ElemName:  "event.object.customFields.id-soa.order",
			ElemValue: float64(13),
		},
		{
			ElemName:  "event.object.customFields.id-soa.string",
			ElemValue: "220041",
		},
		{
			ElemName:  "event.object.customFields.ir-name.order",
			ElemValue: float64(11),
		},
		{
			ElemName:  "event.object.customFields.ir-name.string",
			ElemValue: "ГБУЗ Севастополя \"Медицинский информационно-аналитический центр\"",
		},

		//по этим двум полям пока не понятно
		{ElemName: "event.object.stats", ElemValue: ""},
		{ElemName: "event.object.permissions", ElemValue: ""},
		//---------------------------------------------------
		{ElemName: "event.object._type", ElemValue: "case"},
		{ElemName: "event.object._id", ElemValue: "~85771464712"},
		{ElemName: "event.object.startDate", ElemValue: float64(1704980275686)},
		{ElemName: "event.object.impactStatus", ElemValue: "With Impact"},
		{ElemName: "event.object.status", ElemValue: "Open"},
		{ElemName: "event.object.createdBy", ElemValue: "b.polyakov@cloud.gcm"},
		{ElemName: "event.object.createdAt", ElemValue: float64(1704980275725)},
		{ElemName: "event.object.updatedAt", ElemValue: float64(1705062426568)},
		{ElemName: "event.object.resolutionStatus", ElemValue: "True Positive"},

		//----------- details
		{ElemName: "event.details.summary", ElemValue: "FP (Обращение на getz-club.ru) с 185.4.65.151"},
		{ElemName: "event.details.status", ElemValue: "Resolved"},
		{ElemName: "event.details.impactStatus", ElemValue: "NotApplicable"},
		{ElemName: "event.details.endDate", ElemValue: float64(1705063488183)},
		{ElemName: "event.details.resolutionStatus", ElemValue: "FalsePositive"},
		//details.customFields
		{
			ElemName:  "event.details.customFields.notification.order",
			ElemValue: float64(100),
		},
		{
			ElemName:  "event.details.customFields.class-attack.order",
			ElemValue: float64(2),
		},
		{
			ElemName:  "event.details.customFields.is-incident.order",
			ElemValue: float64(81),
		},
		{
			ElemName:  "event.details.customFields.is-incident.boolean",
			ElemValue: true,
		},
		{
			ElemName:  "event.details.customFields.CNC.order",
			ElemValue: float64(44),
		},
		{
			ElemName:  "event.details.customFields.CNC.string",
			ElemValue: "185.158.114.53",
		},
		{
			ElemName:  "event.details.customFields.first-time.order",
			ElemValue: float64(20),
		},
		{
			ElemName:  "event.details.customFields.first-time.date",
			ElemValue: float64(1705052465000),
		},
		{
			ElemName:  "event.details.customFields.last-time.order",
			ElemValue: float64(121),
		},
		{
			ElemName:  "event.details.customFields.last-time.date",
			ElemValue: float64(1705052479000),
		},
	}
}
