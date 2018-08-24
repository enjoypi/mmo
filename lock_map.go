package ext

import (
	"sync"
)

type lockMap struct {
	m     map[interface{}]interface{}
	mutex sync.RWMutex
}

func NewLockMap() ParallelMap {
	l := lockMap{}
	l.m = make(map[interface{}]interface{})
	return &l
}

func (l *lockMap) Set(k, v interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.m[k] = v
}

func (l *lockMap) Get(k interface{}) interface{} {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	v, ok := l.m[k]
	if ok {
		return v
	} else {
		return nil
	}
}

func (l *lockMap) Delete(k interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.m, k)
}

func (l *lockMap) Len() int {
	return len(l.m)
}
