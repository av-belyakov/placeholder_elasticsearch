package memorytemporarystorage

import (
	"sync"
	"time"
)

// CommonStorageTemporary содержит информацию предназначенную для временного хранения
type CommonStorageTemporary struct {
	dataCounter DataCounterStorage
}

// DataCounterStorage
type DataCounterStorage struct {
	acceptedEvents       int
	processedEvents      int
	eventsDoNotMeetRules int
	eventsMeetRules      int
	insertMongoDB        int
	insertElasticsearch  int
	startTime            time.Time
	sync.Mutex
}

type SettingsInputCase struct {
	TimeCreate int64
	EventId    string
}
