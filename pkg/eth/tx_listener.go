package eth

import (
	"context"
	"fmt"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	log "github.com/golang/glog"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/go-common/pkg/jobs"
)

const (
	// TxListenerTransactionCompleteMsg is the message sent when the transaction is completed
	TxListenerTransactionCompleteMsg = "Transaction is complete"

	// TxListenerTransactionPendingMsg is the message sent when the transaction completion is
	// pending
	TxListenerTransactionPendingMsg = "Transaction is pending"

	// TxListenerTransactionErrorMsgPrefix is the message sent when there is an error with transaction
	// polling
	TxListenerTransactionErrorMsgPrefix = "Error: err:"

	txListenerPrefix = "TxListener"
)

type waiter struct {
	Mutex       *sync.Mutex
	DoneWaiting bool
	QuitChan    chan bool
}

// TxListener provides methods to interact with Ethereum transactions
type TxListener struct {
	blockchain ethereum.TransactionReader
	jobs       jobs.JobService
	timeout    time.Duration
}

// NewTxListenerWithWaitPeriod creates a new TransactionService instance with a wait period
func NewTxListenerWithWaitPeriod(blockchain ethereum.TransactionReader, jobs jobs.JobService, timeout time.Duration) *TxListener {
	return &TxListener{blockchain, jobs, timeout}
}

// NewTxListener creates a new TransactionService instance
func NewTxListener(blockchain ethereum.TransactionReader, jobs jobs.JobService) *TxListener {
	return &TxListener{blockchain, jobs, 0 * time.Second}
}

// StartListener begins listening for an ethereum transaction
func (t *TxListener) StartListener(txID string) (*jobs.Subscription, error) {
	jobID := fmt.Sprintf("%v-%v", txListenerPrefix, txID)
	job, err := t.jobs.StartJob(jobID, func(updates chan<- string) {
		t.PollForTxCompletion(txID, updates)
	})
	if err != nil && err != jobs.ErrJobAlreadyExists {
		return nil, err
	}

	if err == jobs.ErrJobAlreadyExists {
		job, err = t.jobs.GetJob(jobID)
		if err != nil {
			return nil, err
		}
	}

	subscription := job.Subscribe()

	return subscription, nil
}

// StopSubscription will stop subscribing to job updates
// this will not cancel the actual job
func (t *TxListener) StopSubscription(receipt *jobs.Subscription) error {
	return t.jobs.StopSubscription(receipt)
}

// PollForTxCompletion will continuously poll until a transaction is complete
func (t *TxListener) PollForTxCompletion(txID string, updates chan<- string) {
	w := &waiter{
		Mutex:       &sync.Mutex{},
		DoneWaiting: false,
		QuitChan:    make(chan bool),
	}
	go t.waitForTimeout(w)
	defer close(w.QuitChan)

	hash := common.HexToHash(txID)

	ticker := time.NewTicker(time.Millisecond * 500)

	for range ticker.C {
		isPending, err := t.checkTx(hash)
		if err != nil {
			updates <- fmt.Sprintf("%v %v", TxListenerTransactionErrorMsgPrefix, err.Error())
			w.Mutex.Lock()
			if err == ethereum.NotFound && !w.DoneWaiting {
				w.Mutex.Unlock()
				continue
			}
			w.Mutex.Unlock()
			return
		}
		if !isPending {
			updates <- TxListenerTransactionCompleteMsg
			return
		}
		updates <- TxListenerTransactionPendingMsg
	}

}

func (t *TxListener) waitForTimeout(w *waiter) {
	timer := time.NewTimer(t.timeout)
	select {
	case <-timer.C:
		w.Mutex.Lock()
		w.DoneWaiting = true
		w.Mutex.Unlock()
	case <-w.QuitChan:
		log.Infof("Quitting waiting for transaction to appear because it was found or there was another error\n")
		timer.Stop()
	}
}

func (t *TxListener) checkTx(hash common.Hash) (bool, error) {
	_, isPending, err := t.blockchain.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Errorf("Error retrieving TransactionByHash: err: %v\n", err)
		return false, err
	}
	return isPending, nil
}
