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
	startTime            time.Time
	sync.Mutex
}

// DataCounter счетчик данных
// AcceptedEvents       количество принятых событий
// ProcessedEvents      количество обработанных событий
// EventsDoNotMeetRules количество событий не соответствующих правилам
// EventsMeetRules количество событий соответствующих правилам
// StartTime время инициализации счетчика
/*type DataCounter struct {
	AcceptedEvents       int
	ProcessedEvents      int
	EventsDoNotMeetRules int
	EventsMeetRules      int
	StartTime            time.Time
}*/

type SettingsInputCase struct {
	TimeCreate int64
	EventId    string
}
