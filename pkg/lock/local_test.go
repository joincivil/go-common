package lock_test

import (
	"testing"
	"time"

	"github.com/joincivil/go-common/pkg/lock"
	"github.com/joincivil/go-common/pkg/numbers"
)

func interfaceTest2(l lock.DLock) {}

func TestLocalDLock(t *testing.T) {
	dlock := lock.NewLocalDLock()
	key := "test-key"

	interfaceTest2(dlock)

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

func TestLocalDLockExpire(t *testing.T) {
	key := "test-key"

	dlock := lock.NewLocalDLock()
	dlock.Tries = numbers.IntToPtr(64)
	dlock.RetryDelayMillis = numbers.IntToPtr(500)

	// Ensure it is not locked first
	_ = dlock.Unlock(key)

	// Test the lock acquisition runs out of retries error
	err := dlock.Lock(key, numbers.IntToPtr(1000*30))
	if err != nil {
		t.Errorf("Should have gotten lock: %v", err)
	}

	dlock.Tries = numbers.IntToPtr(5)
	dlock.RetryDelayMillis = numbers.IntToPtr(100)

	err = dlock.Lock(key, numbers.IntToPtr(1000*5))
	if err == nil {
		t.Error("Should have returned error with lock timeout")
	}
	if err != lock.ErrNoLockObtained {
		t.Error("Should have returned ErrNoLockObtained")
	}

	// unlock
	_ = dlock.Unlock(key)

	// Test the expiration of the lock
	dlock.Tries = numbers.IntToPtr(64)
	dlock.RetryDelayMillis = numbers.IntToPtr(500)

	err = dlock.Lock(key, numbers.IntToPtr(1000*2))
	if err != nil {
		t.Errorf("Should have gotten lock: %v", err)
	}

	err = dlock.Lock(key, numbers.IntToPtr(1000*2))
	if err != nil {
		t.Errorf("Should not have returned error with expired: err: %v", err)
	}

	// unlock
	_ = dlock.Unlock(key)
}
