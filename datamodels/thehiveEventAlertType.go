package datamodels

// EventMessageTheHiveAlert сообщение с информацией о событии
// Operation - операция
// ObjectId - уникальный идентификатор объекта
// ObjectType - тип объекта
// Base - основа
// StartDate - начальная дата
// RootId - главный уникальный идентификатор
// RequestId - уникальный идентификатор запроса
// EventDetails - детальная информация о событии
// Object - объект события
// OrganisationId - уникальный идентификатор организации
// Organisation - наименование организации
type EventMessageTheHiveAlert struct {
	Base           bool              `json:"base" bson:"base"`
	StartDate      string            `json:"startDate" bson:"startDate"` //в формате RFC3339
	RootId         string            `json:"rootId" bson:"rootId"`
	Organisation   string            `json:"organisation" bson:"organisation"`
	OrganisationId string            `json:"organisationId" bson:"organisationId"`
	ObjectId       string            `json:"objectId" bson:"objectId"`
	ObjectType     string            `json:"objectType" bson:"objectType"`
	Operation      string            `json:"operation" bson:"operation"`
	RequestId      string            `json:"requestId" bson:"requestId"`
	Details        EventAlertDetails `json:"details" bson:"details"`
	Object         EventAlertObject  `json:"object" bson:"object"`
}

// EventAlertDetails детальная информация о событии
// SourceRef - ссылка
// Title - заголовок
// Description - описание
// Tags - список тегов
type EventAlertDetails struct {
	SourceRef   string   `json:"sourceRef" bson:"sourceRef"`
	Title       string   `json:"title" bson:"title"`
	Description string   `json:"description" bson:"description"`
	Tags        []string `json:"tags" bson:"tags"`
}

/*
"object" : {
            +  "_id" : "~85935947856",
            +  "id" : "~85935947856",
            +  "createdBy" : "zsiem@cloud.gcm",
            +  "updatedBy" : "zsiem@cloud.gcm",
            +  "createdAt" : "2024-02-06T05:31:27+00:00",
            +  "updatedAt" : "2024-02-06T05:31:35+00:00",
            +  "_type" : "alert",
              "type" : "snort_alert",
              "source" : "zsiеmSystems",
              "sourceRef" : "TSK-8MSK-6-ZPM-240206-1216010",
              "case" : null, string
			  "caseTemplate" : "snort_alert",
			  "objectType": "alert",
            +  "title" : "Редко встречающиеся признаки ВПО с 94.26.228.205",
            +  "description" : "**Задача переданная из смежной системы: Заслон-Пост-Модерн**В формате ГЦМ: **`TSK-8MSK-6-ZPM-240206-1216010`** ID: `1216010`[http://siem.cloud.gcm/tasks/card/TSK-8MSK-6-ZPM-240206-1216010](http://siem.cloud.gcm/tasks/card/TSK-8MSK-6-ZPM-240206-1216010)Автор задачи: **`Security Event Manager`**Тип: **`snort_alert`****Причина по которой создана задача**Название: `Редко встречающиеся признаки ВПО с 94.26.228.205`Описание: `## Данная задача создана автоматически Время начала: 2024-02-06 08:19:43 Время окончания: 2024-02-06 08:19:43 Продолжительность воздействий: 0:00:00`Отработало на СОА: - **`8030060`**   Российская газета, Установлен: Москва,Москва, IP адрес: 10.20.0.60**Полное описание события IDS:**- Время начала: **`06.02.2024 08:19:43`**- Время окончания: **`06.02.2024 08:19:43`**- **IP из домашней подсети**1. **`213.135.81.78`**- **IP из внешней подсети**1. **`94.26.228.205`****Сигнатуры на которых отработал анализатор сетевого трафика:**1. РП: **`56548067`**, Сообщение: Exploit.CVE-2022-22954.HTTP.C&C, Добавлена: 09.06.2023 10:29:05**Фильтрация и выгрузка от** Tue Feb 06 2024 08:31:35 GMT+0300 Размер: **`4.7 MB`**, [Скачать файл](ftp://ftp.cloud.gcm//traffic/8030060/1707197488_2024_02_06____08_31_28_031838.pcap)Контент: Фильтрация успешно завершена. Включена автоматическая выгрузкаФайл на СОА: `/opt/zaslon/zmanager/data/pfilter_storage/1707197488_2024_02_06____08_31_28_031838.pcap`",
            +  "severity" : 2,
            +  "date" : "2024-02-06T05:31:27+00:00",
            +  "tags" : {
                "ATsreason" : [
                  "\"Редко встречающиеся признаки ВПО\""
                ],
                "Sensorid" : [ ],
                "ATsgeoip" : [
                  "\"Россия\""
                ],
                "WebHooksend" : [ ],
                "Sensor:id=\"8030060\"" : [ ],
                "'Webhook:send=ES'" : [ ]
              },
            +  "tlp" : 2,
            +  "pap" : 2,
            +  "status" : "New",
            +  "follow" : true,
            +  "customFields" : {
                "first-time" : {
                  "date" : "2024-02-06T05:19:43+00:00"
                },
                "last-time" : {
                  "date" : "2024-02-06T05:19:43+00:00"
                }
              }
*/
// EventAlertObject объект события
// Follow - следовать
// Tlp - tlp
// Pap - pap
// Severity - строгость
// UnderliningId - уникальный идентификатор
// Id - уникальный идентификатор
// CreatedBy - кем создан
// UpdatedBy - кем обновлен
// CreatedAt - дата создания
// UpdatedAt - дата обновления
// UnderliningType - тип
// Title - заголовок
// Description - описание
// Tags - список тегов
// Status - статус
// CustomFields - настраиваемые поля
// Date - дата
// Type - тип
// Source - источник
// SourceRef - ссылка на источник
// Case - кейс
// CaseTemplate - шаблон обращения
// ObjectType - тип объекта
type EventAlertObject struct {
	Folow           bool                      `json:"folow" bson:"folow"`
	Severity        uint64                    `json:"severity" bson:"severity"`
	Tlp             uint64                    `json:"tlp" bson:"tlp"`
	Pap             uint64                    `json:"pap" bson:"pap"`
	UnderliningId   string                    `json:"_id" bson:"_id"`
	Id              string                    `json:"id" bson:"id"`
	CreatedBy       string                    `json:"createdBy" bson:"createdBy"`
	UpdatedBy       string                    `json:"updatedBy" bson:"updatedBy"`
	CreatedAt       string                    `json:"createdAt" bson:"createdAt"` //формат RFC3339
	UpdatedAt       string                    `json:"updatedAt" bson:"updatedAt"` //формат RFC3339
	UnderliningType string                    `json:"_type" bson:"_type"`
	Title           string                    `json:"title" bson:"title"`
	Description     string                    `json:"description" bson:"description"`
	Status          string                    `json:"status" bson:"status"`
	Date            string                    `json:"date" bson:"date"` //формат RFC3339
	Type            string                    `json:"type" bson:"type"`
	Source          string                    `json:"source" bson:"source"`
	SourceRef       string                    `json:"sourceRef" bson:"sourceRef"`
	Case            string                    `json:"case" bson:"case"`
	CaseTemplate    string                    `json:"caseTemplate" bson:"caseTemplate"`
	ObjectType      string                    `json:"objectType" bson:"objectType"`
	Tags            []string                  `json:"tags" bson:"tags"`
	CustomFields    map[string]CustomerFields `json:"customFields" bson:"customFields"`
	//"artifacts" : [ ], думаю эти не надо, всегда пустые
	//"similarCases" : [ ] думаю эти не надо, всегда пустые
}
