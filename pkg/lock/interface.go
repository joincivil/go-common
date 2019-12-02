package lock

import "errors"

var (
	// ErrNoLockObtained is the error when a lock is unable to be attained
	ErrNoLockObtained = errors.New("unable to obtain lock")
	// ErrDidNotUnlock is the error when an unlock did not occur
	ErrDidNotUnlock = errors.New("did not unlock")
)

// DLock defines an interface for a distributed lock
type DLock interface {
	// Lock attempts to obtain a lock on the given key. expireMillis sets how long
	// this key should be locked for to prevent deadlocks.
	Lock(key string, expireMillis *int) error

	// Unlock attempts to unlock the given key.
	Unlock(key string) error
}
