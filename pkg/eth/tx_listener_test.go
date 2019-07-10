package eth_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/joincivil/go-common/pkg/eth"
	"github.com/joincivil/go-common/pkg/jobs"
)

func sendTx(t *testing.T, sim bind.ContractBackend, key *ecdsa.PrivateKey) *types.Transaction {
	// generate a transaction and confirm you can retrieve it
	code := `6060604052600a8060106000396000f360606040526008565b00`
	var gas uint64 = 3000000
	tx := types.NewContractCreation(0, big.NewInt(0), gas, big.NewInt(1), common.FromHex(code))
	tx, _ = types.SignTx(tx, types.HomesteadSigner{}, key)

	err := sim.SendTransaction(context.Background(), tx)
	if err != nil {
		t.Fatal("error sending transaction")
	}

	return tx

}

func TestTxListener(t *testing.T) {
	ethHelper, err := eth.NewSimulatedBackendHelper()
	client := ethHelper.Blockchain.(*backends.SimulatedBackend)
	if err != nil {
		t.Fatalf("error with NewSimulatedBackendHelper: %v", err)
	}

	tx := sendTx(t, ethHelper.Blockchain, ethHelper.Key)
	txHash := tx.Hash()
	client.Commit()

	svc := eth.NewTxListener(client, jobs.NewInMemoryJobService())

	sub1, err := svc.StartListener(txHash.String())
	if err != nil {
		t.Fatalf("sub1: unable to get tx subscription %v", txHash.String())
	}

	for event := range sub1.Updates {
		t.Logf("sub1: %v", event)
	}
	t.Log("Complete")

}

func TestTxListenerWait(t *testing.T) {
	ethHelper, err := eth.NewSimulatedBackendHelper()
	client := ethHelper.Blockchain.(*backends.SimulatedBackend)
	if err != nil {
		t.Fatalf("error with NewSimulatedBackendHelper: %v", err)
	}

	txHash := &common.Hash{} // will be error not found

	svc := eth.NewTxListenerWithWaitPeriod(client, jobs.NewInMemoryJobService(), 2*time.Second)

	sub1, err := svc.StartListener(txHash.String())
	if err != nil {
		t.Fatalf("sub1: unable to get tx subscription %v", txHash.String())
	}
	count := 0
	for event := range sub1.Updates {
		count = count + 1
		t.Logf("sub1: %v", event)

	}
	if count < 4 {
		t.Fatalf("checking every half second for 2 seconds it should have collected at least 4 not found errors before terminating the listener but it only found %v", count)
	}
	t.Log("Complete")

	tx := sendTx(t, ethHelper.Blockchain, ethHelper.Key)
	txHash2 := tx.Hash()
	client.Commit()

	svc = eth.NewTxListenerWithWaitPeriod(client, jobs.NewInMemoryJobService(), 2*time.Second)

	sub1, err = svc.StartListener(txHash2.String())
	if err != nil {
		t.Fatalf("sub1: unable to get tx subscription %v", txHash2.String())
	}

	count = 0
	for event := range sub1.Updates {
		count = count + 1
		t.Logf("sub1: %v", event)
	}

	if count != 1 {
		t.Fatalf("shouldnt have waited should have killed the waiting routine")
	}

	t.Log("Complete")

}
