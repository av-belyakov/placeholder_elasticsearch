package coremodule

import (
	"sync"
	"time"
)

var (
	cs   *coreStorage
	once sync.Once
)

type coreStorage struct {
	storage map[string]messageDescriptors
	mx      sync.Mutex
}

type messageDescriptors struct {
	timeCreate int64
	request    string
}

func newStorage() *coreStorage {
	once.Do(func() {
		cs = &coreStorage{storage: make(map[string]messageDescriptors)}

		go checkLiveTime(cs)
	})

	return cs
}

func checkLiveTime(cs *coreStorage) {
	for range time.Tick(5 * time.Second) {
		go func() {
			for k, v := range cs.storage {
				if time.Now().Unix() > (v.timeCreate + 360) {
					cs.deleteElement(k)
				}
			}
		}()
	}
}

func (cs *coreStorage) setRequest(key, value string) {
	cs.mx.Lock()
	defer cs.mx.Unlock()

	cs.storage[key] = messageDescriptors{
		timeCreate: time.Now().Unix(),
		request:    value,
	}
}

func (cs *coreStorage) getRequest(key string) (string, bool) {
	if elem, ok := cs.storage[key]; ok {
		return elem.request, ok
	}

	return "", false
}

func (cs *coreStorage) deleteElement(key string) {
	cs.mx.Lock()
	defer cs.mx.Unlock()

	delete(cs.storage, key)
}
