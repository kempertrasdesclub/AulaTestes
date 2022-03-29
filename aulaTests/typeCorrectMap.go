package aulaTests

import "sync"

type CorrectMap struct {
	m    sync.Mutex
	data map[interface{}]interface{}
}

func (e *CorrectMap) Store(key, value interface{}) {
	e.m.Lock()
	defer e.m.Unlock()

	if e.data == nil {
		e.data = make(map[interface{}]interface{})
	}

	e.data[key] = value
}

func (e *CorrectMap) Load(key interface{}) (value interface{}) {
	e.m.Lock()
	defer e.m.Unlock()

	if e.data == nil {
		e.data = make(map[interface{}]interface{})
	}

	return e.data[key]
}
