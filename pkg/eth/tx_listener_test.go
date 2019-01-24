package eth_test

import (
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joincivil/go-common/pkg/eth"
	"github.com/joincivil/go-common/pkg/jobs"
)

func TestTxListener(t *testing.T) {
	var wg sync.WaitGroup
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		t.Fatal(err)
	}

	svc := eth.NewTxListener(client, jobs.NewInMemoryJobService())

	sub1, err := svc.StartListener("0x3d8a78e268db358a88ae1006138d65dc06f7369617b9b991d694abbce13fe3aa")
	if err != nil {
		t.Fatalf("sub1: unable to get tx subscription")
	}
	sub2, err := svc.StartListener("0x3d8a78e268db358a88ae1006138d65dc06f7369617b9b991d694abbce13fe3aa")
	if err != nil {
		t.Fatalf("sub2: unable to get tx subscription")
	}

	wg.Add(2)
	go func() {
		for event := range sub1.Updates {
			t.Logf("sub1: %v", event)
		}
		wg.Done()
	}()

	go func() {
		for event := range sub2.Updates {
			t.Logf("sub2: %v", event)
		}
		wg.Done()
	}()

	wg.Wait()
	t.Log("Complete")

}
