package memorytemporarystorage

import (
	"sync"
	"time"
)

var once sync.Once
var cst CommonStorageTemporary

func NewTemporaryStorage() *CommonStorageTemporary {
	once.Do(func() {
		cst = CommonStorageTemporary{
			dataCounter: DataCounterStorage{
				insertElasticsearch: make(map[string]int),
			},
		}

		//go checkTimeDelete(&cst)
	})

	return &cst
}

// checkTimeDeleteTemporaryStorageSearchQueries очистка информации о задаче по истечении определенного времени или неактуальности данных
/*func checkTimeDelete(cst *CommonStorageTemporary) {
	c := time.Tick(3 * time.Second)

	for range c {
		go func() {
			for k, v := range cst.HiveFormatMessage.Storages {
				if v.isProcessedMisp && v.isProcessedElasticsearsh && v.isProcessedNKCKI {
					deleteHiveFormatMessageElement(k, cst)
				}
			}
		}()

		go func() {
			for k, v := range cst.temporaryInputCase.Cases {
				if time.Now().Unix() > (v.TimeCreate + 54000) {
					deleteTemporaryCase(k, cst)
				}
			}
		}()
	}
}*/

// GetDataCounter возвращает информацию по сетчикам
func (cst *CommonStorageTemporary) GetDataCounter() DataCounterStorage {
	return DataCounterStorage{
		acceptedEvents:       cst.dataCounter.acceptedEvents,
		processedEvents:      cst.dataCounter.processedEvents,
		eventsMeetRules:      cst.dataCounter.eventsMeetRules,
		eventsDoNotMeetRules: cst.dataCounter.eventsDoNotMeetRules,
		insertMongoDB:        cst.dataCounter.insertMongoDB,
		insertElasticsearch:  cst.dataCounter.insertElasticsearch,
		startTime:            cst.dataCounter.startTime,
	}
}

// GetStartTimeDataCounter возвращает время начала сетчика
func (cst *CommonStorageTemporary) GetStartTimeDataCounter() time.Time {
	return cst.dataCounter.startTime
}

// SetStartTimeDataCounter добавляет время начала сетчика
func (cst *CommonStorageTemporary) SetStartTimeDataCounter(t time.Time) {
	cst.dataCounter.Lock()
	defer cst.dataCounter.Unlock()

	cst.dataCounter.startTime = t
}

// GetAcceptedEventsDataCounter возвращает сетчик принятых событий
func (cst *CommonStorageTemporary) GetAcceptedEventsDataCounter() int {
	return cst.dataCounter.acceptedEvents
}

// SetAcceptedEventsDataCounter увеличивает сетчик принятых событий
func (cst *CommonStorageTemporary) SetAcceptedEventsDataCounter(num int) {
	cst.dataCounter.Lock()
	defer cst.dataCounter.Unlock()

	cst.dataCounter.acceptedEvents += num
}

// GetProcessedEventsDataCounter возвращает сетчик обработанных событий
func (cst *CommonStorageTemporary) GetProcessedEventsDataCounter() int {
	return cst.dataCounter.processedEvents
}

// SetProcessedEventsDataCounter увеличивает сетчик обработанных событий
func (cst *CommonStorageTemporary) SetProcessedEventsDataCounter(num int) {
	cst.dataCounter.Lock()
	defer cst.dataCounter.Unlock()

	cst.dataCounter.processedEvents += num
}

// GetEventsMeetRulesDataCounter возвращает время начала сетчика
func (cst *CommonStorageTemporary) GetEventsMeetRulesDataCounter() int {
	return cst.dataCounter.eventsMeetRules
}

// SetEventsMeetRulesDataCounter увеличивает сетчик событий соответствующих правилу
func (cst *CommonStorageTemporary) SetEventsMeetRulesDataCounter(num int) {
	cst.dataCounter.Lock()
	defer cst.dataCounter.Unlock()

	cst.dataCounter.eventsMeetRules += num
}

// GetEventsDoNotMeetRulesDataCounter возвращает сетчик событий не соответствующих правилу
func (cst *CommonStorageTemporary) GetEventsDoNotMeetRulesDataCounter() int {
	return cst.dataCounter.eventsDoNotMeetRules
}

// SetEventsDoNotMeetRulesDataCounter увеличивает сетчик событий не соответствующих правилу
func (cst *CommonStorageTemporary) SetEventsDoNotMeetRulesDataCounter(num int) {
	cst.dataCounter.Lock()
	defer cst.dataCounter.Unlock()

	cst.dataCounter.eventsDoNotMeetRules += num
}

// GetInsertMongoDBDataCounter возвращает сетчик вставленных в MongoDB документов
func (cst *CommonStorageTemporary) GetInsertMongoDBDataCounter() int {
	return cst.dataCounter.insertMongoDB
}

// SetInsertMongoDBDataCounter увеличивает сетчик вставленных в MongoDB документов
func (cst *CommonStorageTemporary) SetInsertMongoDBDataCounter(num int) {
	cst.dataCounter.Lock()
	defer cst.dataCounter.Unlock()

	cst.dataCounter.insertMongoDB += num
}

// GetAllInsertElasticsearchDataCounter возвращает сетчик вставленных в Elasticsearch документов
func (cst *CommonStorageTemporary) GetAllInsertElasticsearchDataCounter() map[string]int {
	return cst.dataCounter.insertElasticsearch
}

// GetInsertElasticsearchDataCounter возвращает сетчик вставленных в Elasticsearch документов
func (cst *CommonStorageTemporary) GetInsertElasticsearchDataCounter(item string) (int, bool) {
	if v, ok := cst.dataCounter.insertElasticsearch[item]; ok {
		return v, true
	}

	return 0, false
}

// SetInsertElasticsearchDataCounter увеличивает сетчик вставленных в Elasticsearch документов
func (cst *CommonStorageTemporary) SetInsertElasticsearchDataCounter(item string, num int) {
	cst.dataCounter.Lock()
	defer cst.dataCounter.Unlock()

	if _, ok := cst.dataCounter.insertElasticsearch[item]; !ok {
		cst.dataCounter.insertElasticsearch[item] = 0
	}

	cst.dataCounter.insertElasticsearch[item] += num
}
