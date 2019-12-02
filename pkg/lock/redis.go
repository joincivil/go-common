package lock

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

const (
	defaultMaxIdle         = 2
	defaultMaxActive       = 4
	defaultIdleTimeoutSecs = 3
)

// NewRedisLockingPool returns a Redis pool for go-redsync
func NewRedisDLockPool(addr string, maxIdle *int, maxActive *int,
	idleTimeout *int) redsync.Pool {
	mi := defaultMaxIdle
	if maxIdle != nil {
		mi = *maxIdle
	}

	ma := defaultMaxActive
	if maxActive != nil {
		ma = *maxActive
	}

	it := time.Duration(defaultIdleTimeoutSecs) * time.Second
	if idleTimeout != nil {
		it = time.Duration(*idleTimeout) * time.Second
	}

	return &redis.Pool{
		MaxIdle:     mi,
		MaxActive:   ma,
		IdleTimeout: it,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func NewRedisDLock(pools []redsync.Pool, namespace *string) *RedisDLock {
	rl := &RedisDLock{}
	rl.namespace = namespace
	rl.client = redsync.New(pools)
	rl.mutexes = make(map[string]*redsync.Mutex, 100)
	return rl
}

// RedisDLock is an implementation of a distributed lock using Redis.
// This first version expects a single instance of Redis for locking, which isn't
// truly "distributed", but will work for our simple setup.
// Basically a wrapper around https://github.com/go-redsync/redsync that implements
// the Locking interface
type RedisDLock struct {
	client    *redsync.Redsync
	m         sync.Mutex
	mutexes   map[string]*redsync.Mutex
	namespace *string

	MutexTries            *int
	MutexRetryDelayMillis *int
}

// Lock attempts to obtain a lock on the given key. expireMillis sets how long
// this key should be locked for to prevent deadlocks.
func (r *RedisDLock) Lock(key string, expireMillis *int) error {
	fullKey := r.fullKey(key)

	r.m.Lock()
	mt, ok := r.mutexes[fullKey]
	if !ok {
		mt = r.newMutex(fullKey, expireMillis)
		r.mutexes[fullKey] = mt
	}
	r.m.Unlock()

	err := mt.Lock()
	if err == redsync.ErrFailed {
		return ErrNoLockObtained
	} else if err != nil {
		return err
	}
	return nil
}

// Unlock attempts to unlock the given key.
func (r *RedisDLock) Unlock(key string) error {
	fullKey := r.fullKey(key)

	r.m.Lock()
	mt, ok := r.mutexes[fullKey]
	if !ok {
		// No mutex, so nothing to unlock
		r.m.Unlock()
		return ErrDidNotUnlock
	}
	r.m.Unlock()

	// Unlock the dist lock
	res := mt.Unlock()
	if !res {
		return ErrDidNotUnlock
	}

	r.m.Lock()
	delete(r.mutexes, fullKey)
	r.m.Unlock()

	return nil
}

// newMutex returns a initialized redsync.Mutex struct
func (r *RedisDLock) newMutex(key string, expireMillis *int) *redsync.Mutex {
	options := make([]redsync.Option, 0, 3)
	if expireMillis != nil {
		expiry := redsync.SetExpiry(time.Duration(*expireMillis) * time.Millisecond)
		options = append(options, expiry)
	}
	if r.MutexTries != nil {
		tries := redsync.SetTries(*r.MutexTries)
		options = append(options, tries)
	}
	if r.MutexRetryDelayMillis != nil {
		delay := redsync.SetRetryDelay(time.Duration(*r.MutexRetryDelayMillis) * time.Millisecond)
		options = append(options, delay)
	}

	mt := r.client.NewMutex(r.fullKey(key), options...)
	return mt
}

func (r *RedisDLock) fullKey(key string) string {
	key = strings.ToLower(key)
	if r.namespace == nil {
		return key
	}
	return fmt.Sprintf("%v-%v", r.namespace, key)
}
