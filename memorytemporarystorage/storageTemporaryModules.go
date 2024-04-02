package memorytemporarystorage

import (
	"sync"
	"sync/atomic"
	"time"
)

var once sync.Once
var cst CommonStorageTemporary

func NewTemporaryStorage() *CommonStorageTemporary {
	once.Do(func() {
		cst = CommonStorageTemporary{}
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

func (cst *CommonStorageTemporary) GetAlertCounter() StorageCounter {
	return cst.alertCounter
}

func (cst *CommonStorageTemporary) GetCaseCounter() StorageCounter {
	return cst.caseCounter
}

func (cst *CommonStorageTemporary) IncrementAcceptedEvents() {
	atomic.AddUint64(&cst.dataCounter.acceptedEvents, 1)
}

func (cst *CommonStorageTemporary) GetAcceptedEvents() uint64 {
	return cst.dataCounter.acceptedEvents
}

func (cst *CommonStorageTemporary) IncrementProcessedEvents() {
	atomic.AddUint64(&cst.dataCounter.processedEvents, 1)
}

func (cst *CommonStorageTemporary) GetProcessedEvents() uint64 {
	return cst.dataCounter.processedEvents
}

func (cst *CommonStorageTemporary) IncrementAlertEventsMeetRules() {
	atomic.AddUint64(&cst.alertCounter.eventsMeetRules, 1)
}

func (cst *CommonStorageTemporary) GetAlertEventsMeetRules() uint64 {
	return cst.alertCounter.eventsMeetRules
}

func (cst *CommonStorageTemporary) IncrementAlertInsertMongoDB() {
	atomic.AddUint64(&cst.alertCounter.insertMongoDB, 1)
}

func (cst *CommonStorageTemporary) GetAlertInsertMongoDB() uint64 {
	return cst.alertCounter.insertMongoDB
}

func (cst *CommonStorageTemporary) IncrementAlertInsertElasticsearch() {
	atomic.AddUint64(&cst.alertCounter.insertElasticsearch, 1)
}

func (cst *CommonStorageTemporary) GetAlertInsertElasticsearch() uint64 {
	return cst.alertCounter.insertElasticsearch
}

func (cst *CommonStorageTemporary) IncrementCaseEventsMeetRules() {
	atomic.AddUint64(&cst.caseCounter.eventsMeetRules, 1)
}

func (cst *CommonStorageTemporary) GetCaseEventsMeetRules() uint64 {
	return cst.caseCounter.eventsMeetRules
}

func (cst *CommonStorageTemporary) IncrementCaseInsertMongoDB() {
	atomic.AddUint64(&cst.caseCounter.insertMongoDB, 1)
}

func (cst *CommonStorageTemporary) GetCaseInsertMongoDB() uint64 {
	return cst.caseCounter.insertMongoDB
}

func (cst *CommonStorageTemporary) IncrementCaseInsertElasticsearch() {
	atomic.AddUint64(&cst.caseCounter.insertElasticsearch, 1)
}

func (cst *CommonStorageTemporary) GetCaseInsertElasticsearch() uint64 {
	return cst.caseCounter.insertElasticsearch
}

// SetStartTimeDataCounter добавляет время начала сетчика
func (cst *CommonStorageTemporary) SetStartTimeDataCounter(t time.Time) {
	cst.dataCounter.startTime = t
}

// GetStartTimeDataCounter возвращает время начала сетчика
func (cst *CommonStorageTemporary) GetStartTimeDataCounter() time.Time {
	return cst.dataCounter.startTime
}
