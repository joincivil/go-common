package lock_test

import (
	"flag"
	"sync"
	"testing"
	"time"

	// "time"

	"github.com/go-redsync/redsync"
	"github.com/joincivil/go-common/pkg/lock"
	"github.com/joincivil/go-common/pkg/numbers"
	"github.com/joincivil/go-common/pkg/strings"
)

var redis2x = flag.Bool("redis2x", false, "enable two redis pool test")

func interfaceTest(l lock.DLock) {}

func redisDLock() *lock.RedisDLock {
	pools := []redsync.Pool{}
	pool1 := lock.NewRedisDLockPool(
		"127.0.0.1:6379",
		numbers.IntToPtr(1),
		numbers.IntToPtr(1),
		numbers.IntToPtr(10),
	)
	pools = append(pools, pool1)

	if redis2x != nil && *redis2x {
		pool2 := lock.NewRedisDLockPool("127.0.0.1:6378", nil, nil, nil)
		pools = append(pools, pool2)
	}

	return lock.NewRedisDLock(pools, strings.StrToPtr("anamespace"))
}

func TestRedisDLock(t *testing.T) {
	dlock := redisDLock()
	key := "test-key"

	interfaceTest(dlock)

	err := dlock.Unlock(key)
	if err == nil || err != lock.ErrDidNotUnlock {
		t.Errorf("Should have gotten error with unlock")
	}

	// get the lock
	err = dlock.Lock(key, numbers.IntToPtr(1000*10))
	if err != nil {
		t.Errorf("Should have not returned error on lock: err: %v", err)
	}

	go func(k string, dl lock.DLock) {
		// unlock the lock after a few secs
		time.Sleep(2 * time.Second)
		unlockErr := dl.Unlock(k)
		if unlockErr != nil {
			t.Errorf("Should not have gotten error with unlock: err: %v", err)
		}
	}(key, dlock)

	kill := make(chan struct{})
	go func() {
		select {
		case <-time.After(5 * time.Second):
			t.Error("should not have timed out, should have been unlocked")
		case <-kill:
		}
	}()

	// attempt to get the lock while it is locked.  should get the lock after
	// the unlock is called.
	err = dlock.Lock(key, numbers.IntToPtr(1000*10))
	if err != nil {
		t.Errorf("Should have gotten lock after unlock")
	}
	close(kill)

	// unlock
	_ = dlock.Unlock(key)
}

func TestRedisDLockExpire(t *testing.T) {
	dlock := redisDLock()

	dlock.MutexTries = numbers.IntToPtr(64)
	dlock.MutexRetryDelayMillis = numbers.IntToPtr(500)

	key := "test-key"

	// Ensure it is not locked first
	_ = dlock.Unlock(key)

	err := dlock.Lock(key, numbers.IntToPtr(1000*30))
	if err != nil {
		t.Errorf("Should have gotten lock: %v", err)
	}

	dlock2 := redisDLock()
	dlock2.MutexTries = numbers.IntToPtr(5)
	dlock2.MutexRetryDelayMillis = numbers.IntToPtr(100)

	err = dlock2.Lock(key, numbers.IntToPtr(1000*5))
	if err == nil {
		t.Error("Should have returned error with lock timeout")
	}
	if err != lock.ErrNoLockObtained {
		t.Error("Should have returned ErrNoLockObtained")
	}

	// unlock
	_ = dlock.Unlock(key)
	_ = dlock2.Unlock(key)
}

func TestRedisDLockMultiLock(t *testing.T) {
	dlock := redisDLock()
	key := "test-key"

	interfaceTest(dlock)

	err := dlock.Unlock(key)
	if err == nil || err != lock.ErrDidNotUnlock {
		t.Errorf("Should have gotten error with unlock")
	}

	// get the lock
	err = dlock.Lock(key, numbers.IntToPtr(1000*10))
	if err != nil {
		t.Errorf("Should have not returned error on lock: err: %v", err)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err1 := dlock.Lock(key, numbers.IntToPtr(1000*10))
		if err1 != nil {
			t.Errorf("Should have gotten lock after unlock")
		}

		time.Sleep(1 * time.Second)
		err1 = dlock.Unlock(key)
		if err1 != nil {
			t.Errorf("Should not gotten lock after unlock")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		err2 := dlock.Lock(key, numbers.IntToPtr(1000*10))
		if err2 != nil {
			t.Errorf("Should have gotten lock after unlock")
		}

		time.Sleep(1 * time.Second)
		err2 = dlock.Unlock(key)
		if err2 != nil {
			t.Errorf("Should not gotten lock after unlock")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(200 * time.Millisecond)
		err3 := dlock.Lock(key, numbers.IntToPtr(1000*10))
		if err3 != nil {
			t.Errorf("Should have gotten lock after unlock")
		}
		time.Sleep(1 * time.Second)
		err3 = dlock.Unlock(key)
		if err3 != nil {
			t.Errorf("Should not gotten lock after unlock")
		}
	}()

	time.Sleep(500 * time.Millisecond)
	err = dlock.Unlock(key)
	if err != nil {
		t.Errorf("Should have unlocked properly")
	}

	wg.Wait()
}
