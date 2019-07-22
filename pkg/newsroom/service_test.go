package newsroom_test

import (
	"context"
	"fmt"
	shell "github.com/ipfs/go-ipfs-api"
	"io"
	"io/ioutil"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/go-common/pkg/eth"
	"github.com/joincivil/go-common/pkg/generated/contract"
	"github.com/joincivil/go-common/pkg/newsroom"
)

const testHash = "Qma1sdZfxWYWBUcEJGc7WeZX4aFvPoJYqrbzARAJPSjhCe"

const charterJSON = `{
	"name": "test newsroom",
	"tagline": "Stay classy, San Diego"
}`

func deployContracts(t *testing.T, helper *eth.Helper) *eth.DeployerContractAddresses {
	addresses := &eth.DeployerContractAddresses{}
	deployer := eth.NewDeployer(helper, addresses)
	blockchain := helper.Blockchain.(*backends.SimulatedBackend)

	// create ECRecovery library
	err := deployer.DeployECRecovery()
	if err != nil {
		t.Fatal("error with DeployECRecovery")
	}
	// create MessagesAndCodes library
	err = deployer.DeployMessagesAndCodes()
	if err != nil {
		t.Fatal("error with DeployECRecovery")
	}
	// create CivilTokenController contract
	err = deployer.DeployCivilTokenController()
	if err != nil {
		t.Fatal("error creating newsroom factory")
	}

	// create multisig factory
	err = deployer.DeployMultiSigWalletFactory()
	if err != nil {
		t.Fatal("error creating newsroom factory")
	}

	// create newsroom factory
	err = deployer.DeployNewsroomFactory()
	if err != nil {
		t.Fatal("error creating newsroom factory")
	}

	// create newsroom in group contract
	err = deployer.DeployCreateNewsroomInGroup()
	if err != nil {
		t.Fatal("error creating newsroom factory")
	}

	// create CVL token contract
	err = deployer.DeployCVLToken()
	if err != nil {
		t.Fatal("error creating newsroom factory")
	}

	err = deployer.DeployDLL()
	if err != nil {
		t.Fatal("error with DeployDLL")
	}
	err = deployer.DeployAttributeStore()
	if err != nil {
		t.Fatal("error with DeployAttributeStore")
	}

	err = deployer.DeployPLCR()
	if err != nil {
		t.Fatal("error with DeployPLCR")
	}
	pConfig := eth.NewDefaultParameterizerConfig()
	err = deployer.DeployParameterizer(pConfig.AsArray())
	if err != nil {
		t.Fatalf("error with DeployParameterizer: %v", err)
	}

	// deploy government
	gConfig := eth.NewDefaultGovernmentConfig()
	err = deployer.DeployGovernment(gConfig.AsArray(), helper.Auth.From, helper.Auth.From)
	if err != nil {
		t.Fatalf("error with DeployGovernment: %v", err)
	}

	// deploy TCR
	err = deployer.DeployTCR()
	if err != nil {
		t.Fatalf("error with DeployTCR: %v", err)
	}
	blockchain.Commit()

	ctrl, err := contract.NewCivilTokenControllerContract(addresses.CivilTokenController, helper.Blockchain)
	if err != nil {
		t.Fatalf("error creating CivilTokenController instance: %v", err)
	}

	// CreateNewsroomInGroup needs to be a Manager because it whitelists newsrooms
	_, err = ctrl.AddManager(helper.Transact(), addresses.CreateNewsroomInGroup)
	if err != nil {
		t.Fatalf("error AddManager CreateNewsroomInGroup: %v", err)
	}
	// grant "core" to Parameterizer
	_, err = ctrl.AddToCore(helper.Transact(), addresses.Parameterizer)
	if err != nil {
		t.Fatalf("error AddToCore Parameterizer: %v", err)
	}
	// grant "core" to PLCR
	_, err = ctrl.AddToCore(helper.Transact(), addresses.PLCR)
	if err != nil {
		t.Fatalf("error AddToCore PLCR: %v", err)
	}
	// grant "core" to TCR
	_, err = ctrl.AddToCore(helper.Transact(), addresses.TCR)
	if err != nil {
		t.Fatalf("error AddToCore TCR: %v", err)
	}
	// allow the default account to send CVL (not needed in prod, but used for testing)
	_, err = ctrl.AddToCore(helper.Transact(), helper.Auth.From)
	if err != nil {
		t.Fatalf("error AddToCore: %v", err)
	}

	blockchain.Commit()

	return addresses
}

type MockIPFS struct {
}

func (m MockIPFS) Cat(path string) (io.ReadCloser, error) {
	if path == testHash {
		reader := ioutil.NopCloser(strings.NewReader(charterJSON))
		return reader, nil
	}
	return ioutil.NopCloser(strings.NewReader("")), fmt.Errorf("not found")
}

func (m MockIPFS) Add(r io.Reader, options ...shell.AddOpts) (string, error) {
	return testHash, nil
}

func TestNewsroomService(t *testing.T) {
	ipfs := MockIPFS{}

	ethHelper, err := eth.NewSimulatedBackendHelper()
	if err != nil {
		t.Fatal("error starting simulated backend")
	}
	addresses := deployContracts(t, ethHelper)

	svc, err := newsroom.NewService(ethHelper, ipfs, *addresses)
	if err != nil {
		t.Fatal("error starting newsroom Service")
	}
	blockchain := ethHelper.Blockchain.(*backends.SimulatedBackend)

	createNewsroom := func(name string) common.Address {
		tx, err := svc.CreateNewsroom(name, testHash)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		blockchain.Commit()
		address, err := svc.GetNewsroomAddressFromTransaction(tx)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		return address
	}

	t.Run("CreateNewsroom", func(t *testing.T) {
		address := createNewsroom("test")

		if (address == common.Address{}) {
			t.Fatal("error creating newsroom, invalid address")
		}
	})

	t.Run("GetNewsroomName", func(t *testing.T) {
		newsroomName := "test"
		newsroomAddress := createNewsroom(newsroomName)
		name, err := svc.GetNewsroomName(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		if name != newsroomName {
			t.Fatalf("expected newsroom name to be `%v` to received: %v", newsroomName, name)
		}
	})

	t.Run("GetOwner", func(t *testing.T) {
		newsroomAddress := createNewsroom("test")
		multisigAddr, err := svc.GetOwner(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}

		if (multisigAddr == common.Address{}) {
			t.Fatalf("invalid multisig address")
		}

	})

	t.Run("GetMultisigMembers", func(t *testing.T) {
		newsroomAddress := createNewsroom("test")
		members, err := svc.GetMultisigMembers(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}

		if len(members) != 1 {
			t.Fatalf("expected multisig to have 1 member")
		}
		if members[0] != ethHelper.Auth.From {
			t.Fatalf("expected multisig owner to be " + ethHelper.Auth.From.String())
		}
	})

	t.Run("ApplyToTCR", func(t *testing.T) {
		newsroomAddress := createNewsroom("test")
		multisigAddr, err := svc.GetOwner(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}

		cvl, err := contract.NewCVLTokenContract(addresses.CVLToken, blockchain)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}

		tokens := big.NewInt(1000)
		tokens = tokens.Mul(tokens, big.NewInt(1e18))

		// fund the multisig with 1k CVL
		_, err = cvl.Transfer(ethHelper.Transact(), multisigAddr, tokens)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		blockchain.Commit()

		// allow the TCR to transferFrom the multisig (needed to apply)
		_, err = svc.AdminApproveTCRTokenTransfer(newsroomAddress, tokens)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		blockchain.Commit()

		// apply to the TCR
		tx, err := svc.AdminApplyToTCR(newsroomAddress, tokens)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		blockchain.Commit()
		// get the receipt to confirm the tx was successful
		receipt, err := ethHelper.Blockchain.(ethereum.TransactionReader).TransactionReceipt(context.Background(), tx)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}

		if receipt.Status == 0 {
			t.Fatalf("transaction failed, expected status of 1")
		}
	})

	t.Run("AdminAddMember", func(t *testing.T) {

		newsroomAddress := createNewsroom("test")
		// just a random address
		newMember := common.HexToAddress("0x8c29E088ddd8233C2B2A35b77ff94f8B208b79b7")
		txHash, err := svc.AdminAddMultisigOwner(newsroomAddress, newMember)
		if err != nil {
			t.Fatalf("not expecting error with AdminAddMultisigOwner %v", err)
		}
		blockchain.Commit()

		receipt, err := ethHelper.Blockchain.(ethereum.TransactionReader).TransactionReceipt(context.Background(), txHash)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		if receipt.Status != 1 {
			t.Fatalf("AddMember transaction failed: %v", receipt.Status)
		}

		members, err := svc.GetMultisigMembers(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		if len(members) != 2 {
			t.Fatalf("expected multisig to have 2 members")
		}
		if members[1] != common.HexToAddress("0x8c29E088ddd8233C2B2A35b77ff94f8B208b79b7") {
			t.Fatalf("expected to find the added address as a member")
		}
	})

	t.Run("AdminRemoveMember", func(t *testing.T) {

		newsroomAddress := createNewsroom("test")
		// just a random address
		owner1 := common.HexToAddress("0x34CB3f586187eB930E17f5c45F607ac78bbce6ae")
		owner2 := common.HexToAddress("0x8c29E088ddd8233C2B2A35b77ff94f8B208b79b7")
		_, err := svc.AdminAddMultisigOwner(newsroomAddress, owner1)
		if err != nil {
			t.Fatalf("not expecting error with AdminAddMultisigOwner %v", err)
		}
		_, err = svc.AdminAddMultisigOwner(newsroomAddress, owner2)
		if err != nil {
			t.Fatalf("not expecting error with AdminAddMultisigOwner %v", err)
		}

		blockchain.Commit()

		members, err := svc.GetMultisigMembers(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		if len(members) != 3 {
			t.Fatalf("expected multisig to have 3 members")
		}

		txHash, err := svc.AdminRemoveMultisigOwner(newsroomAddress, owner2)
		if err != nil {
			t.Fatalf("not expecting error with AdminRemoveMultisigOwner %v", err)
		}
		blockchain.Commit()

		receipt, err := ethHelper.Blockchain.(ethereum.TransactionReader).TransactionReceipt(context.Background(), txHash)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		if receipt.Status != 1 {
			t.Fatalf("RemoveMember transaction failed: %v", receipt.Status)
		}

		members, err = svc.GetMultisigMembers(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		if len(members) != 2 {
			t.Fatalf("expected multisig to have 2 members")
		}
		if members[1] != owner1 {
			t.Fatalf("expected to find the added address as a member")
		}
	})

	t.Run("RenameNewsroom", func(t *testing.T) {

		newsroomAddress := createNewsroom("foo")

		_, err := svc.RenameNewsroom(newsroomAddress, "bar")
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		blockchain.Commit()

		name, err := svc.GetNewsroomName(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}
		if name != "bar" {
			t.Fatalf("did not receive the expected newsroom name")
		}

	})

	t.Run("GetCharter", func(t *testing.T) {
		newsroomAddress := createNewsroom("test newsroom")

		charter, err := svc.GetCharter(newsroomAddress)
		if err != nil {
			t.Fatalf("not expecting an error but received: %v", err)
		}

		if charter.Name != "test newsroom" {
			t.Fatalf("expecting name to be `test newsroom` but is: %v", charter.Name)
		}

		if charter.Tagline != "Stay classy, San Diego" {
			t.Fatalf("expecting tagline to be to be `Stay classy, San Diego` but is: %v", charter.Tagline)
		}
	})

}
