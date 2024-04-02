package memorytemporarystorage

import (
	"time"
)

// CommonStorageTemporary содержит информацию предназначенную для временного хранения
type CommonStorageTemporary struct {
	dataCounter  DataCounterStorage
	alertCounter StorageCounter
	caseCounter  StorageCounter
}

type StorageCounter struct {
	eventsMeetRules     uint64
	insertMongoDB       uint64
	insertElasticsearch uint64
}

// DataCounterStorage
type DataCounterStorage struct {
	acceptedEvents  uint64
	processedEvents uint64
	startTime       time.Time
}

type SettingsInputCase struct {
	TimeCreate int64
	EventId    string
}
