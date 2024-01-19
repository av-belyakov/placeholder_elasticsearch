package testhandlerobservables

func GetReportsList() [][2]string {
	reportsList := [][2]string{
		{"reports.CyberCrime-Tracker_1_0.taxonomies.level", "info"},
		{"reports.CyberCrime-Tracker_1_0.taxonomies.namespace", "CCT"},
		{"reports.CyberCrime-Tracker_1_0.taxonomies.predicate", "C2 Search"},
		{"reports.CyberCrime-Tracker_1_0.taxonomies.value", "0 hits"},

		{"reports.CyberCrime-Tracker_1_0.taxonomies.level", "info"},
		{"reports.CyberCrime-Tracker_1_0.taxonomies.namespace", "HTy"},
		{"reports.CyberCrime-Tracker_1_0.taxonomies.predicate", "No search out"},
		{"reports.CyberCrime-Tracker_1_0.taxonomies.value", "78 bit"},

		{"reports.DShield_lookup_1_0.taxonomies.level", "not info"},
		{"reports.DShield_lookup_1_0.taxonomies.namespace", "Jonson A"},
		{"reports.DShield_lookup_1_0.taxonomies.predicate", "Tyo"},
		{"reports.DShield_lookup_1_0.taxonomies.value", "1000"},

		{"reports.URLhaus_2_0.taxonomies.level", "info"},
		{"reports.URLhaus_2_0.taxonomies.namespace", "Lochkarev"},
		{"reports.URLhaus_2_0.taxonomies.predicate", "CCNoK"},
		{"reports.URLhaus_2_0.taxonomies.value", "8bit"},

		{"reports.Urlscan_io_Search_0_1_1.taxonomies.level", "warning"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", "LPPPT1"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", "12-34"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.value", "NONE"},

		{"reports.Urlscan_io_Search_0_1_1.taxonomies.level", "suspicious"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", "urlscan.io"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", "Search out"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.value", "6 results"},

		{"reports.Urlscan_io_Search_0_1_1.taxonomies.level", "safe"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.namespace", "Maltiverse"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.predicate", "Report"},
		{"reports.Urlscan_io_Search_0_1_1.taxonomies.value", "6 neutral"},
	}

	return reportsList
}
