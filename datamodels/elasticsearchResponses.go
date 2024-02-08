package datamodels

// ElasticsearchResponseCase может использоватся для анмаршалинга
// json ответов содержащих информацию о кейсах
type ElasticsearchResponseCase struct {
	Options ElasticsearchResponseCaseOptions `json:"hits"`
}

type ElasticsearchResponseCaseOptions struct {
	Total    OptionsTotal `json:"total"`
	MaxScore float64      `json:"max_score"`
	Hits     []CaseHits   `json:"hits"`
}

type OptionsTotal struct {
	Relation string `json:"relation"`
	Value    int    `json:"value"`
}

type CaseHits struct {
	ID    string `json:"_id"`
	Index string `json:"_index"`
}
