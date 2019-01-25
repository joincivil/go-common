package eth_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joincivil/go-common/pkg/eth"
)

func TestRetryChainReader(t *testing.T) {
	client, _ := ethclient.Dial("https://rinkeby.infura.io/")
	retryClient := &eth.RetryChainReader{ChainReader: client}

	header, err := retryClient.HeaderByNumberWithRetry(1, 3, 200)
	if err != nil {
		t.Errorf("Should not have received an error: err: %v", err)
	}

	if header == nil {
		t.Errorf("Should not have been a nil header")
	}
}

func TestRetryChainReaderBadEthClient(t *testing.T) {
	client, _ := ethclient.Dial("http://civil.co")
	retryClient := &eth.RetryChainReader{ChainReader: client}

	header, err := retryClient.HeaderByNumberWithRetry(1, 3, 200)
	if err == nil {
		t.Errorf("Should have received an error: err: %v", err)
	}

	if header != nil {
		t.Errorf("Should have been a nil header")
	}
}
