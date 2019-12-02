// +build integration

package lock_test

import (
	"testing"
	"time"

	// "time"

	"github.com/go-redsync/redsync"
	"github.com/joincivil/go-common/pkg/lock"
	"github.com/joincivil/go-common/pkg/numbers"
	"github.com/joincivil/go-common/pkg/strings"
)

func interfaceTest(l lock.DLock) {}

func TestRedisDLock(t *testing.T) {
	pool1 := lock.NewRedisDLockPool("127.0.0.1:6379",
		numbers.IntToPtr(1), numbers.IntToPtr(1), numbers.IntToPtr(10))
	pool2 := lock.NewRedisDLockPool("127.0.0.1:6378", nil, nil, nil)

	dlock := lock.NewRedisDLock([]redsync.Pool{pool1, pool2}, strings.StrToPtr("anamespace"))
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

	kill := make(chan bool)
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
	pool1 := lock.NewRedisDLockPool("127.0.0.1:6379",
		numbers.IntToPtr(1), numbers.IntToPtr(1), numbers.IntToPtr(10))
	pool2 := lock.NewRedisDLockPool("127.0.0.1:6378", nil, nil, nil)

	dlock := lock.NewRedisDLock([]redsync.Pool{pool1, pool2}, nil)
	dlock.MutexTries = numbers.IntToPtr(64)
	dlock.MutexRetryDelayMillis = numbers.IntToPtr(500)

	key := "test-key"

	// Ensure it is not locked first
	_ = dlock.Unlock(key)

	err := dlock.Lock(key, numbers.IntToPtr(1000*30))
	if err != nil {
		t.Errorf("Should have gotten lock: %v", err)
	}

	dlock = lock.NewRedisDLock([]redsync.Pool{pool1, pool2}, nil)
	dlock.MutexTries = numbers.IntToPtr(5)
	dlock.MutexRetryDelayMillis = numbers.IntToPtr(100)

	err = dlock.Lock(key, numbers.IntToPtr(1000*5))
	if err == nil {
		t.Error("Should have returned error with lock timeout")
	}
	if err != lock.ErrNoLockObtained {
		t.Error("Should have returned ErrNoLockObtained")
	}

	// unlock
	_ = dlock.Unlock(key)
}
