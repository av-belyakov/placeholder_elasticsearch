package datamodels

// ElasticsearchResponseCase может использоватся для анмаршалинга
// json ответов содержащих информацию о кейсах
type ElasticsearchResponseCase struct {
	Options ElasticsearchResponseCaseOptions `json:"hits"`
}

// ElasticsearchResponseAlert может использоватся для анмаршалинга
// json ответов содержащих информацию об алертах
type ElasticsearchResponseAlert struct {
	Options ElasticsearchResponseAlertOptions `json:"hits"`
}

type ElasticsearchResponseCaseOptions struct {
	Total    OptionsTotal                            `json:"total"`
	MaxScore float64                                 `json:"max_score"`
	Hits     []ElasticsearchPatternVerifiedForEsCase `json:"hits"`
}

type ElasticsearchResponseAlertOptions struct {
	Total    OptionsTotal                             `json:"total"`
	MaxScore float64                                  `json:"max_score"`
	Hits     []ElasticsearchPatternVerifiedForEsAlert `json:"hits"`
}

type ElasticsearchPatternVerifiedForEsCase struct {
	Source VerifiedEsCase `json:"_source"`
	ServiseOption
}

type ElasticsearchPatternVerifiedForEsAlert struct {
	Source VerifiedForEsAlert `json:"_source"`
	ServiseOption
}

// OptionsTotal опции в результате поиска
// Relation - отношение (==, >, <)
// Value - количество найденных значений
type OptionsTotal struct {
	Relation string `json:"relation"`
	Value    int    `json:"value"`
}

type ServiseOption struct {
	ID    string `json:"_id"`
	Index string `json:"_index"`
}
