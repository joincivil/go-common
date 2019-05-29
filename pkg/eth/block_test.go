package eth_test

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/go-common/pkg/eth"
	ctime "github.com/joincivil/go-common/pkg/time"
)

func TestBlockHeaderCache(t *testing.T) {
	expiryCacheSecs := 1
	cache := eth.NewBlockHeaderCache(int64(expiryCacheSecs))
	header := cache.HeaderByBlockNumber(uint64(1))
	if header != nil {
		t.Error("Should have failed to retrieve any headers")
	}

	ts := uint64(ctime.CurrentEpochSecsInInt64())
	header1 := &types.Header{
		Time: ts,
	}

	cache.AddHeader(uint64(1), header1)

	header = cache.HeaderByBlockNumber(uint64(1))
	if header == nil {
		t.Error("Should have retrieved a header")
	}
	if header.Time != ts {
		t.Error("Should have been the same time as the added Header")
	}

	time.Sleep(time.Duration(expiryCacheSecs+1) * time.Second)

	header = cache.HeaderByBlockNumber(uint64(1))
	if header != nil {
		t.Error("Should have not retrieved a header after duration")
	}
}
