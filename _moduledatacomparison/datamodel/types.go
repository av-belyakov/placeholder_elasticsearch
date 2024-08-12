package datamodel

// ChanOutputDecodeJSON содержит данные получаемые при декодировании JSON формата
// обрабатываемого обработчиком HandlerMessageFromHive
// FieldName - наименование поля
// ValueType - тип передаваемого значения (string, int и т.д.)
// Value - любые передаваемые данные
// FieldBranch - 'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
type ChanOutputDecodeJSON struct {
	FieldName   string
	ValueType   string
	Value       interface{}
	FieldBranch string
}

// SupportiveObservables вспомогательный тип для для обработки observables
/*type SupportiveObservables struct {
	currentKey string
	//previousFieldReports      string
	listAcceptedFields        []string
	listAcceptedFieldsReports []string
	observableTmp             datamodels.ObservableMessageEs
	observables               map[string][]datamodels.ObservableMessageEs
}
*/
