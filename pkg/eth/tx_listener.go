package eth

import (
	"context"
	"fmt"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
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

// ReadTransaction is a temporary interface until go-ethereum/pull/19026 is merged
// at which point TxListener should probably use ethereum.TransactionReader interface
type ReadTransaction interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

// TxListener provides methods to interact with Ethereum transactions
type TxListener struct {
	blockchain ReadTransaction
	jobs       jobs.JobService
}

// NewTxListener creates a new TransactionService instance
func NewTxListener(blockchain ReadTransaction, jobs jobs.JobService) *TxListener {
	return &TxListener{blockchain, jobs}
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

	hash := common.HexToHash(txID)

	ticker := time.NewTicker(time.Millisecond * 500)

	for range ticker.C {
		isPending, err := t.checkTx(hash)
		if err != nil {
			updates <- fmt.Sprintf("%v %v", TxListenerTransactionErrorMsgPrefix, err.Error())
			return
		}
		if !isPending {
			updates <- TxListenerTransactionCompleteMsg
			return
		}
		updates <- TxListenerTransactionPendingMsg
	}

}

func (t *TxListener) checkTx(hash common.Hash) (bool, error) {

	// todo(dankins): TransactionByHash is not implemented by SimulatedBackend
	// once https://github.com/ethereum/go-ethereum/pull/19026 is resolved we should be able to use the following:
	// _, isPending, err := t.blockchain.TransactionByHash(context.Background(), hash)

	// this method has no way of determining if a tx just doesn't exist,
	// so there is potential an invalid tx ID will just poll forever
	tx, err := t.blockchain.TransactionReceipt(context.Background(), hash)
	if err == ethereum.NotFound {
		return true, nil
	}
	if err != nil {
		log.Errorf("Error retrieving TransactionReceipt: err: %v\n", err)
		return false, err
	}

	if tx == nil {
		return true, nil
	}

	return false, nil
}
