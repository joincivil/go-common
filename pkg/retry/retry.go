package retry

import (
	"time"

	log "github.com/golang/glog"
)

// WrappedFn is the definition of the function that the Wrapper will attempt to try.
type WrappedFn func() (interface{}, error)

// Wrapper retries a retry function WrappedFn a maxAttempts number of times,
// backing off based on on baseWaitMs.  The backoff is calculated the baseWaitMs * attempt
// number in milliseconds.
// The return interface{} needs to be type inferred based on the true return type of the
// retryFn.
func Wrapper(retryFunc WrappedFn, maxAttempts int, baseWaitMs int) (interface{}, error) {
	attempt := 1
	var result interface{}
	var err error

	for {
		result, err = retryFunc()
		if err != nil {

			if attempt >= maxAttempts {
				log.Errorf("Retry func, maxed retries err: %v", err)
				return nil, err
			}

			log.Infof(
				"Error, sleep/attempt again, waiting %v ms: err: %v",
				baseWaitMs*attempt,
				err,
			)

			time.Sleep(time.Duration(baseWaitMs) * time.Duration(attempt) * time.Millisecond)
			attempt++

			log.Infof("Retrying, attempt %v...", attempt)
			continue
		}
		return result, err
	}
}
