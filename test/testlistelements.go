package testhandlerobservables

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
			ElemName:  "event.object.customFields.class-attack.string",
			ElemValue: "Exploite",
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
			ElemName:  "event.details.customFields.class-attack.string",
			ElemValue: "Вирусное заражение",
		},
		{
			ElemName:  "event.details.customFields.attack-type.order",
			ElemValue: float64(7),
		},
		{
			ElemName:  "event.details.customFields.attack-type.string",
			ElemValue: "Распределенная КА",
		},
		{
			ElemName:  "event.details.customFields.event-source.order",
			ElemValue: float64(10),
		},
		{
			ElemName:  "event.details.customFields.event-source.string",
			ElemValue: "Заслон-ПМ. Ведомственное СОА",
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

func GetReportsList() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "observables.reports.CyberCrime-Tracker_1_0.taxonomies.level", ElemValue: "info"},
		{ElemName: "observables.reports.CyberCrime-Tracker_1_0.taxonomies.namespace", ElemValue: "CCT"},
		{ElemName: "observables.reports.CyberCrime-Tracker_1_0.taxonomies.predicate", ElemValue: "C2 Search"},
		{ElemName: "observables.reports.CyberCrime-Tracker_1_0.taxonomies.value", ElemValue: "0 hits"},

		{ElemName: "observables.reports.CyberCrime-Tracker_1_0.taxonomies.level", ElemValue: "info"},
		{ElemName: "observables.reports.CyberCrime-Tracker_1_0.taxonomies.namespace", ElemValue: "HTy"},
		{ElemName: "observables.reports.CyberCrime-Tracker_1_0.taxonomies.predicate", ElemValue: "No search out"},
		{ElemName: "observables.reports.CyberCrime-Tracker_1_0.taxonomies.value", ElemValue: "78 bit"},

		{ElemName: "observables.reports.DShield_lookup_1_0.taxonomies.level", ElemValue: "not info"},
		{ElemName: "observables.reports.DShield_lookup_1_0.taxonomies.namespace", ElemValue: "Jonson A"},
		{ElemName: "observables.reports.DShield_lookup_1_0.taxonomies.predicate", ElemValue: "Tyo"},
		{ElemName: "observables.reports.DShield_lookup_1_0.taxonomies.value", ElemValue: "1000"},

		{ElemName: "observables.reports.URLhaus_2_0.taxonomies.level", ElemValue: "info"},
		{ElemName: "observables.reports.URLhaus_2_0.taxonomies.namespace", ElemValue: "Lochkarev"},
		{ElemName: "observables.reports.URLhaus_2_0.taxonomies.predicate", ElemValue: "CCNoK"},
		{ElemName: "observables.reports.URLhaus_2_0.taxonomies.value", ElemValue: "8bit"},

		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.level", ElemValue: "warning"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", ElemValue: "LPPPT1"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", ElemValue: "12-34"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.value", ElemValue: "NONE"},

		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.level", ElemValue: "suspicious"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", ElemValue: "urlscan.io"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", ElemValue: "Search out"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.value", ElemValue: "6 results"},

		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.level", ElemValue: "safe"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", ElemValue: "Maltiverse"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", ElemValue: "Report"},
		{ElemName: "observables.reports.Urlscan_io_Search_0_1_1.taxonomies.value", ElemValue: "6 neutral"},
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

func GetObservableTwo() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "observables._id", ElemValue: "~542580736"},
		{ElemName: "observables.data", ElemValue: "/dbdata/dump/events/58964/B2M-58964.pcap"},
		{ElemName: "observables.dataType", ElemValue: "url_pcap"},
		{ElemName: "observables.ioc", ElemValue: true},
		{ElemName: "observables._type", ElemValue: "Observable"},
		{ElemName: "observables.tlp", ElemValue: float64(2)},
		{ElemName: "observables._createdAt", ElemValue: float64(1705049449272)},
		{ElemName: "observables._createdBy", ElemValue: "zhmurchuk@mail.rcm"},
		{ElemName: "observables._updatedAt", ElemValue: float64(1705049448855)},
		{ElemName: "observables.sighted", ElemValue: false},
		{ElemName: "observables.startDate", ElemValue: float64(1705049449272)},
		{ElemName: "observables.tags", ElemValue: []interface{}{
			"misp:Network activity=\"ip-src\"",
			"b2m:ip_ext=206.189.15.25",
		}},

		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.", ElemValue: ""},
	}
}

func GetObservableThree() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "observables._id", ElemValue: "~1916473544"},
		{ElemName: "observables.data", ElemValue: "91003345,95000138"},
		{ElemName: "observables.dataType", ElemValue: "snort_sid"},
		{ElemName: "observables.ioc", ElemValue: true},
		{ElemName: "observables._type", ElemValue: "Observable"},
		{ElemName: "observables.tlp", ElemValue: float64(3)},
		{ElemName: "observables._createdAt", ElemValue: float64(1690968664227)},
		{ElemName: "observables._createdBy", ElemValue: "zsiem@rcm.lcl"},
		{ElemName: "observables._updatedAt", ElemValue: float64(1704977151860)},
		{ElemName: "observables.sighted", ElemValue: false},
		{ElemName: "observables.startDate", ElemValue: float64(1690968664227)},
		{ElemName: "observables.tags", ElemValue: []interface{}{
			"snort",
			"misp:Network activity=\"snort\"",
		}},

		{ElemName: "observables.attachment.contentType", ElemValue: "text/plain"},
		{ElemName: "observables.attachment.id", ElemValue: "00000000000000000000c398ac1"},
		{ElemName: "observables.attachment.name", ElemValue: "d23t.dll"},
		{ElemName: "observables.attachment.size", ElemValue: float64(817)},
		{ElemName: "observables.attachment.hashes", ElemValue: []interface{}{
			"000000000000000000c29438b04791184d3eba39bdb",
			"7c531394dc2f483bc00000000000000006c6c628c02e0788",
		}},
		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.extraData", ElemValue: ""},
		//{ElemName: "observables.", ElemValue: ""},
	}
}

func GetTtpOne() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "ttp._createdAt", ElemValue: float64(1705041429370)},
		{ElemName: "ttp._createdBy", ElemValue: "prs@rcm"},
		{ElemName: "ttp._id", ElemValue: "~185286688"},
		{ElemName: "ttp.occurDate", ElemValue: float64(1705041420000)},
		{ElemName: "ttp.patternId", ElemValue: "T1110.001"},
		{ElemName: "ttp.tactic", ElemValue: "credential-access"},

		//pattern
		{ElemName: "ttp.extraData.pattern.remoteSupport", ElemValue: true},
		{ElemName: "ttp.extraData.pattern.revoked", ElemValue: true},
		{ElemName: "ttp.extraData.pattern._createdAt", ElemValue: float64(1705041429370)},
		{ElemName: "ttp.extraData.pattern._createdBy", ElemValue: "admin@thehive.local"},
		{ElemName: "ttp.extraData.pattern._id", ElemValue: "~164016"},
		{ElemName: "ttp.extraData.pattern._type", ElemValue: "Pattern"},
		{ElemName: "ttp.extraData.pattern.description", ElemValue: "exploit edge network"},
		{ElemName: "ttp.extraData.pattern.detection", ElemValue: "Monitor authentication logs for system"},
		{ElemName: "ttp.extraData.pattern.name", ElemValue: "Password Guessing"},
		{ElemName: "ttp.extraData.pattern.patternId", ElemValue: "T1110.001"},
		{ElemName: "ttp.extraData.pattern.patternType", ElemValue: "attack-pattern"},
		{ElemName: "ttp.extraData.pattern.url", ElemValue: "https://attack.mitre.org/techniques/T1110/001"},
		{ElemName: "ttp.extraData.pattern.version", ElemValue: "1.4"},
		{ElemName: "ttp.extraData.pattern.platforms", ElemValue: []interface{}{
			"Windows",
			"Azure AD",
			"Office 365",
		}},
		{ElemName: "ttp.extraData.pattern.permissionsRequired", ElemValue: []interface{}{
			"User",
			"Administrator",
		}},
		{ElemName: "ttp.extraData.pattern.dataSources", ElemValue: []interface{}{
			"Application Log: Application Log Content",
			"User Account: User Account Authentication",
		}},
		{ElemName: "ttp.extraData.pattern.tactics", ElemValue: []interface{}{
			"credential-access",
		}},

		//patternParent
		{ElemName: "ttp.extraData.patternParent.remoteSupport", ElemValue: true},
		{ElemName: "ttp.extraData.patternParent.revoked", ElemValue: true},
		{ElemName: "ttp.extraData.patternParent._createdAt", ElemValue: float64(1705040481513)},
		{ElemName: "ttp.extraData.patternParent._createdBy", ElemValue: "root@thehive.gcm"},
		{ElemName: "ttp.extraData.patternParent._id", ElemValue: "~236516"},
		{ElemName: "ttp.extraData.patternParent._type", ElemValue: "Pattern Parent"},
		{ElemName: "ttp.extraData.patternParent.description", ElemValue: "Adversaries may"},
		{ElemName: "ttp.extraData.patternParent.detection", ElemValue: "Adversaries with no prior knowledge"},
		{ElemName: "ttp.extraData.patternParent.name", ElemValue: "Password Guessing"},
		{ElemName: "ttp.extraData.patternParent.patternId", ElemValue: "T1110"},
		{ElemName: "ttp.extraData.patternParent.patternType", ElemValue: "attack-pattern"},
		{ElemName: "ttp.extraData.patternParent.url", ElemValue: "https://attack.mitre.org/techniques/T1110"},
		{ElemName: "ttp.extraData.patternParent.version", ElemValue: "1.14"},
		{ElemName: "ttp.extraData.patternParent.platforms", ElemValue: []interface{}{
			"Office 365",
			"SaaS",
			"IaaS",
			"Linux",
			"macOS",
			"Google Workspace",
		}},
		{ElemName: "ttp.extraData.patternParent.permissionsRequired", ElemValue: []interface{}{
			"Root",
			"User",
			"Administrator",
		}},
		{ElemName: "ttp.extraData.patternParent.dataSources", ElemValue: []interface{}{
			"Application Log: Application Log Content",
			"User Account: User Account Authentication",
			"Command: Command Execution",
		}},
		{ElemName: "ttp.extraData.patternParent.tactics", ElemValue: []interface{}{
			"credential-access",
		}},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
	}
}

func GetTtpTwo() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "ttp._createdAt", ElemValue: float64(1705032829013)},
		{ElemName: "ttp._createdBy", ElemValue: "dv-cloud@gcm"},
		{ElemName: "ttp._id", ElemValue: "~104177744"},
		{ElemName: "ttp.occurDate", ElemValue: float64(1705032780000)},
		{ElemName: "ttp.patternId", ElemValue: "T1190"},
		{ElemName: "ttp.tactic", ElemValue: "initial-access"},

		//pattern
		{ElemName: "ttp.extraData.pattern.remoteSupport", ElemValue: true},
		{ElemName: "ttp.extraData.pattern.revoked", ElemValue: false},
		{ElemName: "ttp.extraData.pattern._createdAt", ElemValue: float64(1705032829013)},
		{ElemName: "ttp.extraData.pattern._createdBy", ElemValue: "prs@rcm"},
		{ElemName: "ttp.extraData.pattern._id", ElemValue: "~104177744"},
		{ElemName: "ttp.extraData.pattern._type", ElemValue: "Pattern"},
		{ElemName: "ttp.extraData.pattern.detection", ElemValue: "Monitor application logs"},
		{ElemName: "ttp.extraData.pattern.name", ElemValue: "Password Guessing"},
		{ElemName: "ttp.extraData.pattern.patternId", ElemValue: "T1190"},
		{ElemName: "ttp.extraData.pattern.patternType", ElemValue: "attack-pattern"},
		{ElemName: "ttp.extraData.pattern.description", ElemValue: "It`s very important message"},
		{ElemName: "ttp.extraData.pattern.url", ElemValue: "https://attack.mitre.org/techniques/T1110/001"},
		{ElemName: "ttp.extraData.pattern.version", ElemValue: "2.4"},
		{ElemName: "ttp.extraData.pattern.platforms", ElemValue: []interface{}{
			"Windows",
			"Azure AD",
			"Office 365",
			"IaaS",
		}},
		{ElemName: "ttp.extraData.pattern.permissionsRequired", ElemValue: []interface{}{
			"User",
			"Admin",
			"FreeUser",
			"JobUser",
		}},
		{ElemName: "ttp.extraData.pattern.dataSources", ElemValue: []interface{}{
			"Application Log: Application Log Content",
			"User Account: User Account Authentication",
		}},
		{ElemName: "ttp.extraData.pattern.tactics", ElemValue: []interface{}{
			"credential-access",
		}},

		//patternParent
		{ElemName: "ttp.extraData.patternParent.remoteSupport", ElemValue: true},
		{ElemName: "ttp.extraData.patternParent.revoked", ElemValue: true},
		{ElemName: "ttp.extraData.patternParent._createdAt", ElemValue: float64(1705040481513)},
		{ElemName: "ttp.extraData.patternParent._createdBy", ElemValue: "root@example.gcm"},
		{ElemName: "ttp.extraData.patternParent._id", ElemValue: "~16400016"},
		{ElemName: "ttp.extraData.patternParent._type", ElemValue: "Pattern Parent"},
		//{ElemName: "ttp.extraData.patternParent.description", ElemValue: "applications are often websites/web"},
		{ElemName: "ttp.extraData.patternParent.detection", ElemValue: "Adversaries with no prior knowledge"},
		{ElemName: "ttp.extraData.patternParent.name", ElemValue: "Password Guessing"},
		{ElemName: "ttp.extraData.patternParent.patternId", ElemValue: "T1220"},
		{ElemName: "ttp.extraData.patternParent.patternType", ElemValue: "attack-pattern"},
		{ElemName: "ttp.extraData.patternParent.url", ElemValue: "https://attack.mitre.org/techniques/T1220"},
		{ElemName: "ttp.extraData.patternParent.version", ElemValue: "1.1411"},
		{ElemName: "ttp.extraData.patternParent.platforms", ElemValue: []interface{}{
			"Office 365",
			"SaaS",
			"IaaS",
			"Linux",
			"macOS",
			"Google Workspace",
		}},
		{ElemName: "ttp.extraData.patternParent.permissionsRequired", ElemValue: []interface{}{
			"Root",
			"User",
			"Administrator",
		}},
		{ElemName: "ttp.extraData.patternParent.dataSources", ElemValue: []interface{}{
			"Application Log: Application Log Content",
			"User Account: User Account Authentication",
			"Command: Command Execution",
		}},
		{ElemName: "ttp.extraData.patternParent.tactics", ElemValue: "resurce-development"},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
	}
}

func GetTtpThree() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "ttp._createdAt", ElemValue: float64(1705041429370)},
		{ElemName: "ttp._createdBy", ElemValue: "ag@rcm.ekb"},
		{ElemName: "ttp._id", ElemValue: "~698159200"},
		{ElemName: "ttp.occurDate", ElemValue: float64(1705041420000)},
		{ElemName: "ttp.patternId", ElemValue: "T1110.001"},
		{ElemName: "ttp.tactic", ElemValue: "initial-access"},

		//pattern
		{ElemName: "ttp.extraData.pattern.remoteSupport", ElemValue: true},
		{ElemName: "ttp.extraData.pattern.revoked", ElemValue: true},
		{ElemName: "ttp.extraData.pattern._createdAt", ElemValue: float64(1705041429370)},
		{ElemName: "ttp.extraData.pattern._createdBy", ElemValue: "adminxxx@thehive.local"},
		{ElemName: "ttp.extraData.pattern._id", ElemValue: "~61546640"},
		{ElemName: "ttp.extraData.pattern._type", ElemValue: "Pattern"},
		{ElemName: "ttp.extraData.pattern.description", ElemValue: "Spearphishing attachment"},
		{ElemName: "ttp.extraData.pattern.detection", ElemValue: "ACSC Email Spoofing"},
		{ElemName: "ttp.extraData.pattern.name", ElemValue: "Spearphishing Attachment"},
		{ElemName: "ttp.extraData.pattern.patternId", ElemValue: "T1110.001"},
		{ElemName: "ttp.extraData.pattern.patternType", ElemValue: "attack-pattern"},
		{ElemName: "ttp.extraData.pattern.url", ElemValue: "https://attack.mitre.org/techniques/T1110/001"},
		{ElemName: "ttp.extraData.pattern.version", ElemValue: "2.2"},
		{ElemName: "ttp.extraData.pattern.platforms", ElemValue: []interface{}{
			"macOS",
			"Windows",
			"Office 365",
		}},
		{ElemName: "ttp.extraData.pattern.permissionsRequired", ElemValue: []interface{}{
			"User",
		}},
		{ElemName: "ttp.extraData.pattern.dataSources", ElemValue: []interface{}{
			"Network Traffic: Network Traffic Content",
			"Application Log: Application Log Content",
			"File: File Creation",
		}},
		{ElemName: "ttp.extraData.pattern.tactics", ElemValue: []interface{}{
			"initial-access",
		}},

		//patternParent
		{ElemName: "ttp.extraData.patternParent.remoteSupport", ElemValue: true},
		{ElemName: "ttp.extraData.patternParent.revoked", ElemValue: true},
		{ElemName: "ttp.extraData.patternParent._createdAt", ElemValue: float64(1705040481513)},
		{ElemName: "ttp.extraData.patternParent._createdBy", ElemValue: "rootxx@thehive.gcm"},
		{ElemName: "ttp.extraData.patternParent._id", ElemValue: "~346894448"},
		{ElemName: "ttp.extraData.patternParent._type", ElemValue: "Pattern cxv Parent"},
		{ElemName: "ttp.extraData.patternParent.description", ElemValue: "Phishing may also"},
		{ElemName: "ttp.extraData.patternParent.detection", ElemValue: "SSL/TLS inspection"},
		{ElemName: "ttp.extraData.patternParent.name", ElemValue: "Password Guessing"},
		{ElemName: "ttp.extraData.patternParent.patternId", ElemValue: "T1110"},
		{ElemName: "ttp.extraData.patternParent.patternType", ElemValue: "attack-pattern"},
		{ElemName: "ttp.extraData.patternParent.url", ElemValue: "https://attack.mitre.org/techniques/T1110"},
		{ElemName: "ttp.extraData.patternParent.version", ElemValue: "2.14"},
		{ElemName: "ttp.extraData.patternParent.platforms", ElemValue: []interface{}{
			"Office 365",
			"Linux",
			"macOS",
			"Google Workspace",
		}},
		{ElemName: "ttp.extraData.patternParent.permissionsRequired", ElemValue: []interface{}{
			"Root",
			"User",
			"Administrator",
		}},
		{ElemName: "ttp.extraData.patternParent.dataSources", ElemValue: []interface{}{
			"Application Log: Application Log Content",
			"User Account: User Account Authentication",
			"Command: Command Execution",
		}},
		{ElemName: "ttp.extraData.patternParent.tactics", ElemValue: []interface{}{
			"initial-access",
		}},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
		//{ElemName: "ttp.patternParent.", ElemValue: ""},
	}
}

func GetEventForAlertOne() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "event.base", ElemValue: true},
		{ElemName: "event.startDate", ElemValue: float64(1707223072562)},
		{ElemName: "event.rootId", ElemValue: "~84625227848"},
		{ElemName: "event.objectId", ElemValue: "~84625227848"},
		{ElemName: "event.objectType", ElemValue: "alert"},
		{ElemName: "event.organisation", ElemValue: "GCM"},
		{ElemName: "event.organisationId", ElemValue: "~4192"},
		{ElemName: "event.operation", ElemValue: "update"},
		{ElemName: "event.requestId", ElemValue: "55994d44f3b276c1:6483e5ec:18d786c2f83:-8000:138497"},

		//----------- details
		{ElemName: "event.details.sourceRef", ElemValue: "TSK-8MSK-6-ZPM-240206-1215999"},
		{ElemName: "event.details.title", ElemValue: "Зафиксированна КА"},
		{ElemName: "event.details.description", ElemValue: "**Задача переданная из смежной системы"},
		{ElemName: "event.details.tags", ElemValue: []interface{}{
			"ATs:geoip=\"Индия\"",
			"'Webhook:send=ES'",
			"Sensor:id=\"8030066\"",
			"ATs:reason=\"INFO Controlled FGS\"",
			"ATs:reason=\"Зафиксированы признаки взаимодействия ВПО класса trojan\"",
		}},

		//----------- object
		{ElemName: "event.object.follow", ElemValue: true},
		{ElemName: "event.object.severity", ElemValue: float64(2)},
		{ElemName: "event.object.tlp", ElemValue: float64(3)},
		{ElemName: "event.object.pap", ElemValue: float64(5)},
		{ElemName: "event.object._id", ElemValue: "~85455464790"},
		{ElemName: "event.object.id", ElemValue: "~85771464712"},
		{ElemName: "event.object.createdBy", ElemValue: "v.kovanko@cloud.gcm"},
		{ElemName: "event.object.updatedBy", ElemValue: "y.kovalenko@cloud.gcm"},
		{ElemName: "event.object.createdAt", ElemValue: float64(1707223072375)},
		{ElemName: "event.object.updatedAt", ElemValue: float64(0)},
		{ElemName: "event.object._type", ElemValue: "alert"},
		{ElemName: "event.object.title", ElemValue: "Редко встречающиеся признаки ВПО"},
		{ElemName: "event.object.description", ElemValue: "из смежной системы: Заслон-Пост-Модерн**"},
		{ElemName: "event.object.status", ElemValue: "Ignored"},
		{ElemName: "event.object.date", ElemValue: float64(1707223017220)},
		{ElemName: "event.object.type", ElemValue: "snort_alert"},
		{ElemName: "event.object.objectType", ElemValue: "alert"},
		{ElemName: "event.object.source", ElemValue: "zsiеmSystems"},
		{ElemName: "event.object.sourceRef", ElemValue: "TSK-8MSK-6-ZPM-240206-1215999"},
		{ElemName: "event.object.case", ElemValue: "aallert"},
		{ElemName: "event.object.caseTemplate", ElemValue: "snort_alert"},
		{ElemName: "event.object.tags", ElemValue: []interface{}{
			"Sensor:id=\"8030012\"",
			"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
			"'Webhook:send=ES'",
		}},
		//object.customFields
		{
			ElemName:  "event.object.customFields.first-time.date",
			ElemValue: float64(1707222030000),
		},
		{
			ElemName:  "event.object.customFields.last-time.date",
			ElemValue: float64(1707222030000),
		},
	}
}

func GetAlertForAlertOne() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "alert.tlp", ElemValue: float64(2)},
		{ElemName: "alert.severity", ElemValue: float64(3)},
		{ElemName: "alert.date", ElemValue: float64(1707220877000)},
		{ElemName: "alert.createdAt", ElemValue: float64(1707221547881)},
		{ElemName: "alert.type", ElemValue: "snort_alert"},
		{ElemName: "alert.description", ElemValue: "внутрь, протоколы: **smtp/tcp**"},
		{ElemName: "alert.status", ElemValue: "Update"},
	}
}

func GetAlertForAlertTwo() []struct {
	ElemName  string
	ElemValue interface{}
} {
	return []struct {
		ElemName  string
		ElemValue interface{}
	}{
		{ElemName: "alert.follow", ElemValue: true},
		{ElemName: "alert.tlp", ElemValue: float64(2)},
		{ElemName: "alert.severity", ElemValue: float64(3)},
		{ElemName: "alert.date", ElemValue: float64(1707195545887)},
		{ElemName: "alert.createdAt", ElemValue: float64(1707195545896)},
		{ElemName: "alert.updatedAt", ElemValue: float64(1707223500033)},
		{ElemName: "alert.updatedBy", ElemValue: "webhook@cloud.gcm"},
		{ElemName: "alert._id", ElemValue: "~88026357960"},
		{ElemName: "alert.status", ElemValue: "New"},
		{ElemName: "alert.type", ElemValue: "snort_alert"},
		{ElemName: "alert._type", ElemValue: "alert"},
		{ElemName: "alert.description", ElemValue: "Атака направлена **внутрь**, использует протоколы: **smtp/tcp**"},
		{ElemName: "alert.caseTemplate", ElemValue: "snort_alert"},
		{ElemName: "alert.sourceRef", ElemValue: "TSK-8MSK-6-ZPM-240206-1216137"},
		{ElemName: "alert.tags", ElemValue: []interface{}{
			"Sensor:id=\"8030105\"",
			"ATs:reason=\"Редко встречающиеся признаки ВПО\"",
			"'Webhook:send=ES'",
			"APPA:Direction=\"inbound\"",
		}},
		//object.customFields
		{
			ElemName:  "alert.customFields.first-time.date",
			ElemValue: float64(1707194932000),
		},
		{
			ElemName:  "alert.customFields.last-time.date",
			ElemValue: float64(1707194932000),
		},
		//-------- url_arkime ---------
		{ElemName: "alert.artifacts.ioc", ElemValue: true},
		{ElemName: "alert.artifacts.sighted", ElemValue: true},
		{ElemName: "alert.artifacts.ignoreSimilarity", ElemValue: true},
		{ElemName: "alert.artifacts.tlp", ElemValue: float64(2)},
		{ElemName: "alert.artifacts._id", ElemValue: "~84302291040"},
		{ElemName: "alert.artifacts.id", ElemValue: "~84302291040"},
		{ElemName: "alert.artifacts._type", ElemValue: "case_artifact"},
		{ElemName: "alert.artifacts.createdAt", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.updatedAt", ElemValue: float64(0)},
		{ElemName: "alert.artifacts.startDate", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.createdBy", ElemValue: "e.anisimova@cloud.gcm"},
		{ElemName: "alert.artifacts.updatedBy", ElemValue: ""},
		{ElemName: "alert.artifacts.data", ElemValue: "http://anisimova.cloud.gcm:8005/sessions?expression=file%20%3D%3D%20%2Fopt%2Farkime%2Fraw%2F2024_02_06_15_03_01_522219____1707221109_2024_02_06____15_05_09_364449.pcap&date=-1"},
		{ElemName: "alert.artifacts.dataType", ElemValue: "url_arkime"},
		{ElemName: "alert.artifacts.message", ElemValue: "Ссылка на Arkime"},
		{ElemName: "alert.artifacts.tags", ElemValue: []interface{}{
			"home",
		}},
		//-------- url ---------
		{ElemName: "alert.artifacts.ioc", ElemValue: true},
		{ElemName: "alert.artifacts.sighted", ElemValue: true},
		{ElemName: "alert.artifacts.ignoreSimilarity", ElemValue: true},
		{ElemName: "alert.artifacts.tlp", ElemValue: float64(2)},
		{ElemName: "alert.artifacts._id", ElemValue: "~82693193920"},
		{ElemName: "alert.artifacts.id", ElemValue: "~82693193920"},
		{ElemName: "alert.artifacts._type", ElemValue: "case_artifact"},
		{ElemName: "alert.artifacts.createdAt", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.updatedAt", ElemValue: float64(0)},
		{ElemName: "alert.artifacts.startDate", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.createdBy", ElemValue: "zsiem@cloud.gcm"},
		{ElemName: "alert.artifacts.updatedBy", ElemValue: ""},
		{ElemName: "alert.artifacts.data", ElemValue: "cdnsvc.microft.dynssl.com"},
		{ElemName: "alert.artifacts.dataType", ElemValue: "url"},
		{ElemName: "alert.artifacts.message", ElemValue: "некий URL 11"},
		{ElemName: "alert.artifacts.tags", ElemValue: []interface{}{
			"dns_alert:resolve_ip",
		}},
		//-------- url_pcap ---------
		{ElemName: "alert.artifacts.ioc", ElemValue: true},
		{ElemName: "alert.artifacts.sighted", ElemValue: true},
		{ElemName: "alert.artifacts.ignoreSimilarity", ElemValue: true},
		{ElemName: "alert.artifacts.tlp", ElemValue: float64(1)},
		{ElemName: "alert.artifacts._id", ElemValue: ""},
		{ElemName: "alert.artifacts.id", ElemValue: "~84998705312"},
		{ElemName: "alert.artifacts._type", ElemValue: "case_artifact"},
		{ElemName: "alert.artifacts.createdAt", ElemValue: float64(1707221114151)},
		{ElemName: "alert.artifacts.updatedAt", ElemValue: float64(1707221114151)},
		{ElemName: "alert.artifacts.startDate", ElemValue: float64(1707221114151)},
		{ElemName: "alert.artifacts.createdBy", ElemValue: ""},
		{ElemName: "alert.artifacts.updatedBy", ElemValue: "zsiem@cloud.gcm"},
		{ElemName: "alert.artifacts.data", ElemValue: "ftp://ftp.cloud.gcm/traffic/8030030/1707221109_2024_02_06____15_05_09_364449.pcap"},
		{ElemName: "alert.artifacts.dataType", ElemValue: "url_pcap"},
		{ElemName: "alert.artifacts.message", ElemValue: "Download a piece of traffic"},
		{ElemName: "alert.artifacts.tags", ElemValue: []interface{}{
			"geo:country=\"Россия\"",
			"geo:iso=\"RU\"",
			"geo:region=\"Калининградская область\"",
			"external",
			"geo:city=\"Калининград\"",
		}},
		//-------- url ---------
		{ElemName: "alert.artifacts.ioc", ElemValue: true},
		{ElemName: "alert.artifacts.sighted", ElemValue: false},
		{ElemName: "alert.artifacts.ignoreSimilarity", ElemValue: true},
		{ElemName: "alert.artifacts.tlp", ElemValue: float64(3)},
		{ElemName: "alert.artifacts._id", ElemValue: "~88180887800"},
		{ElemName: "alert.artifacts.id", ElemValue: "~88180887800"},
		{ElemName: "alert.artifacts._type", ElemValue: "case_artifact"},
		{ElemName: "alert.artifacts.createdAt", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.updatedAt", ElemValue: float64(0)},
		{ElemName: "alert.artifacts.startDate", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.createdBy", ElemValue: "zsiemsss@cloud.gcm"},
		{ElemName: "alert.artifacts.updatedBy", ElemValue: ""},
		{ElemName: "alert.artifacts.data", ElemValue: "ftp.microft.dynssl.com"},
		{ElemName: "alert.artifacts.dataType", ElemValue: "url"},
		{ElemName: "alert.artifacts.message", ElemValue: "некий URL 22"},
		{ElemName: "alert.artifacts.tags", ElemValue: []interface{}{
			"home_ip",
		}},
		//-------- ip_home ---------
		{ElemName: "alert.artifacts.ioc", ElemValue: true},
		{ElemName: "alert.artifacts.sighted", ElemValue: false},
		{ElemName: "alert.artifacts.ignoreSimilarity", ElemValue: true},
		{ElemName: "alert.artifacts.tlp", ElemValue: float64(3)},
		{ElemName: "alert.artifacts._id", ElemValue: "~84612182192"},
		{ElemName: "alert.artifacts.id", ElemValue: "~84612182192"},
		{ElemName: "alert.artifacts._type", ElemValue: "case_artifact"},
		{ElemName: "alert.artifacts.createdAt", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.updatedAt", ElemValue: float64(0)},
		{ElemName: "alert.artifacts.startDate", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.createdBy", ElemValue: "rre.zsiem@cloud.gcm"},
		{ElemName: "alert.artifacts.updatedBy", ElemValue: ""},
		{ElemName: "alert.artifacts.data", ElemValue: "8030160:95.167.52.134"},
		{ElemName: "alert.artifacts.dataType", ElemValue: "ip_home"},
		{ElemName: "alert.artifacts.message", ElemValue: "некий URL 22"},
		{ElemName: "alert.artifacts.tags", ElemValue: []interface{}{
			"home_ip",
		}},
		//-------- url ---------
		{ElemName: "alert.artifacts.ioc", ElemValue: false},
		{ElemName: "alert.artifacts.sighted", ElemValue: true},
		{ElemName: "alert.artifacts.ignoreSimilarity", ElemValue: true},
		{ElemName: "alert.artifacts.tlp", ElemValue: float64(4)},
		{ElemName: "alert.artifacts._id", ElemValue: "~86870781992"},
		{ElemName: "alert.artifacts.id", ElemValue: "~86870781992"},
		{ElemName: "alert.artifacts._type", ElemValue: "case_artifact"},
		{ElemName: "alert.artifacts.createdAt", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.updatedAt", ElemValue: float64(0)},
		{ElemName: "alert.artifacts.startDate", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.createdBy", ElemValue: "zsiemsss@cloud.gcm"},
		{ElemName: "alert.artifacts.updatedBy", ElemValue: ""},
		{ElemName: "alert.artifacts.data", ElemValue: "wmiprvse.xyz"},
		{ElemName: "alert.artifacts.dataType", ElemValue: "url"},
		{ElemName: "alert.artifacts.message", ElemValue: "некий URL 33"},
		{ElemName: "alert.artifacts.tags", ElemValue: []interface{}{
			"dns_alert:resolve_ip",
		}},
		//-------- snort_sid ---------
		{ElemName: "alert.artifacts.ioc", ElemValue: false},
		{ElemName: "alert.artifacts.sighted", ElemValue: true},
		{ElemName: "alert.artifacts.ignoreSimilarity", ElemValue: true},
		{ElemName: "alert.artifacts.tlp", ElemValue: float64(4)},
		{ElemName: "alert.artifacts._id", ElemValue: "~85113589920"},
		{ElemName: "alert.artifacts.id", ElemValue: "~85113589920"},
		{ElemName: "alert.artifacts._type", ElemValue: "case_artifact"},
		{ElemName: "alert.artifacts.createdAt", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.updatedAt", ElemValue: float64(0)},
		{ElemName: "alert.artifacts.startDate", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.createdBy", ElemValue: "framp@cloud.gcm"},
		{ElemName: "alert.artifacts.updatedBy", ElemValue: ""},
		{ElemName: "alert.artifacts.data", ElemValue: "61453, 56407, 46937, 30293288, 44801769"},
		{ElemName: "alert.artifacts.dataType", ElemValue: "snort_sid"},
		{ElemName: "alert.artifacts.message", ElemValue: "правила снорта"},
		{ElemName: "alert.artifacts.tags", ElemValue: []interface{}{
			"snort",
			"20.9 MB",
		}},
		//-------- snort_sid ---------
		{ElemName: "alert.artifacts.ioc", ElemValue: false},
		{ElemName: "alert.artifacts.sighted", ElemValue: true},
		{ElemName: "alert.artifacts.ignoreSimilarity", ElemValue: true},
		{ElemName: "alert.artifacts.tlp", ElemValue: float64(4)},
		{ElemName: "alert.artifacts._id", ElemValue: "~82697081024"},
		{ElemName: "alert.artifacts.id", ElemValue: "~82697081024"},
		{ElemName: "alert.artifacts._type", ElemValue: "case_artifact"},
		{ElemName: "alert.artifacts.createdAt", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.updatedAt", ElemValue: float64(0)},
		{ElemName: "alert.artifacts.startDate", ElemValue: float64(1707221115602)},
		{ElemName: "alert.artifacts.createdBy", ElemValue: "framp@cloud.gcm"},
		{ElemName: "alert.artifacts.updatedBy", ElemValue: ""},
		{ElemName: "alert.artifacts.data", ElemValue: "3004388"},
		{ElemName: "alert.artifacts.dataType", ElemValue: "snort_sid"},
		{ElemName: "alert.artifacts.message", ElemValue: "правила Snort"},
		{ElemName: "alert.artifacts.tags", ElemValue: []interface{}{
			"snort",
			"external",
		}},
	}
}
