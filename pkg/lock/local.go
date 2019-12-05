package lock

import "sync"

import (
	"strings"
	"time"

	"github.com/joincivil/go-common/pkg/numbers"
	ctime "github.com/joincivil/go-common/pkg/time"
)

const (
	defaultTries            = 32
	defaultRetryDelayMillis = 500
)

type lockMeta struct {
	expireAt *int
}

func NewLocalDLock() *LocalDLock {
	dlock := &LocalDLock{
		locks: make(map[string]*lockMeta, 5),
	}
	return dlock
}

// LocalDLock is an implementation of a distributed lock that is a wrapper about
// a simple in memory map.
// NOTE: This *isn't* distributed and should only be used for small single
// server projects or for testing services that require an DLock implementation.
type LocalDLock struct {
	m     sync.Mutex
	trymx sync.Mutex
	locks map[string]*lockMeta

	Tries            *int
	RetryDelayMillis *int
}

func (m *LocalDLock) Lock(key string, expireMillis *int) error {
	key = strings.ToLower(key)
	l := m.lockMeta(key)

	if l == nil {
		m.acquire(key, expireMillis)
		return nil
	}

	// Lock already acquired, poll to see if it expires
	tries := defaultTries
	if m.Tries != nil {
		tries = *m.Tries
	}

	delay := defaultRetryDelayMillis
	if m.RetryDelayMillis != nil {
		delay = *m.RetryDelayMillis
	}

	for i := 0; i < tries; i++ {
		// Check to see if it was unlocked
		m.trymx.Lock()
		l = m.lockMeta(key)
		if l == nil {
			m.acquire(key, expireMillis)
			m.trymx.Unlock()
			return nil
		}

		// If lock is still active, has it expired if expiry is set
		nowMs := ctime.CurrentEpochSecsInInt() * 1000
		if l.expireAt != nil && nowMs > *l.expireAt {
			m.release(key)
			m.acquire(key, expireMillis)
			m.trymx.Unlock()
			return nil
		}
		m.trymx.Unlock()

		time.Sleep(time.Duration(delay) * time.Millisecond)
	}

	return ErrNoLockObtained
}

// Unlock attempts to unlock the given key.
func (m *LocalDLock) Unlock(key string) error {
	key = strings.ToLower(key)

	l := m.lockMeta(key)
	if l == nil {
		return ErrDidNotUnlock
	}

	m.release(key)

	return nil
}

func (m *LocalDLock) acquire(key string, expireMillis *int) {
	l := &lockMeta{}
	if expireMillis != nil {
		nowMs := ctime.CurrentEpochSecsInInt() * 1000
		l.expireAt = numbers.IntToPtr(nowMs + *expireMillis)
	}

	m.m.Lock()
	m.locks[key] = l
	m.m.Unlock()
}

func (m *LocalDLock) release(key string) {
	m.m.Lock()
	delete(m.locks, key)
	m.m.Unlock()
}

func (m *LocalDLock) lockMeta(key string) *lockMeta {
	m.m.Lock()
	l, ok := m.locks[key]
	m.m.Unlock()
	if !ok {
		return nil
	}
	return l
}
