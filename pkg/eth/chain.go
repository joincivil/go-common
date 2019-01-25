package eth

import (
	"context"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/go-common/pkg/retry"
)

// RetryChainReader is a ChainReader that includes some version of ChainReader
// functions that are wrapped by a retry mechanism
type RetryChainReader struct {
	ethereum.ChainReader
}

// HeaderByNumberWithRetry is a version of HeaderByNumber that has a retry
// mechanism
func (r *RetryChainReader) HeaderByNumberWithRetry(blockNumber uint64, maxAttempts int,
	baseWaitMs int) (*types.Header, error) {

	retryFn := func() (interface{}, error) {
		blockNum := big.NewInt(0)
		blockNum.SetUint64(blockNumber)
		return r.HeaderByNumber(context.Background(), blockNum)
	}

	result, err := retry.Wrapper(retryFn, maxAttempts, baseWaitMs)
	if err != nil {
		return nil, err
	}

	return result.(*types.Header), nil
}
