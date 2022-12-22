package wen

import (
	"sync"
	"time"
)

type KeyLock struct {
	m sync.Map
}

func (k *KeyLock) TryLock(key interface{}) bool {
	_, ok := k.m.LoadOrStore(key, struct{}{})
	return !ok
}

func (k *KeyLock) WaitLock(key interface{}, retry int) bool {
	for i := 0; i < retry; i++ {
		if k.TryLock(key) {
			return true
		} else {
			time.Sleep(10 * time.Microsecond)
		}
	}
	return false
}

func (k *KeyLock) UnLock(key interface{}) {
	k.m.Delete(key)
}
