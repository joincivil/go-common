package token_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/go-common/pkg/eth"
	"github.com/joincivil/go-common/pkg/generated/contract"
	"github.com/joincivil/go-common/pkg/token"
)

func TestDetectTransferRestriction(t *testing.T) {
	helper, err := eth.NewSimulatedBackendHelper()
	if err != nil {
		t.Fatal("could not create simulated backend helper")
	}
	blockchain := helper.Blockchain.(*backends.SimulatedBackend)

	checkTx := func(tx *types.Transaction, err error, key string) {
		if err != nil {
			t.Fatalf("Error with %v: %v", key, err)
		}

		blockchain.Commit()
		rcpt, err := blockchain.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			t.Fatalf("Error with %v TransactionReceipt: %v", key, err)
		}
		if rcpt.Status != 1 {
			t.Fatalf("Error with %v status code: %v", key, rcpt.Status)
		}
	}

	controllerAddress, tx, _, err := deployCivilTokenController(helper)
	checkTx(tx, err, "CivilTokenController")

	tokenAddress, err := deployCVLToken(controllerAddress, helper)
	checkTx(tx, err, "CVLToken")

	svc, err := token.NewCVLTokenService(helper.Blockchain, controllerAddress, tokenAddress)
	if err != nil {
		t.Fatalf("Unable to construct CVLTokenService %v", err)
	}

	genesis := helper.Accounts["genesis"].Address
	newUser := common.HexToAddress("0xd7E6ECF1568C4cA70422807D4a68E5CFF3946f25")
	civilianUser := common.HexToAddress("0xdccd8fC67F20C9084b355303E06fD137AD166c54")
	unlockedUser := common.HexToAddress("0xd289026D1A4e24A55406dEeca118f46460F6458e")

	tx, err = svc.ControllerContract.AddManager(helper.TransactWithGasLimit(), genesis)
	checkTx(tx, err, "AddManager")

	tx, err = svc.ControllerContract.AddToCore(helper.TransactWithGasLimit(), genesis)
	checkTx(tx, err, "AddToCore")

	tx, err = svc.ControllerContract.AddToCivilians(helper.TransactWithGasLimit(), civilianUser)
	checkTx(tx, err, "AddToCivilians1")
	tx, err = svc.ControllerContract.AddToCivilians(helper.TransactWithGasLimit(), unlockedUser)
	checkTx(tx, err, "AddToCivilians2")

	tx, err = svc.ControllerContract.AddToUnlocked(helper.TransactWithGasLimit(), unlockedUser)
	checkTx(tx, err, "AddToUnlocked")

	isCivilian, err := svc.ControllerContract.CivilianList(helper.Call(), genesis)
	if err != nil {
		t.Fatal(err)
	}
	if isCivilian {
		t.Fatal("genesis isCivilian should be false")
	}

	isCore, err := svc.ControllerContract.CoreList(helper.Call(), genesis)
	if err != nil {
		t.Fatal(err)
	}
	if !isCore {
		t.Fatal("genesis isCivilian should be true")
	}

	tx, err = svc.TokenContract.Transfer(helper.TransactWithGasLimit(), newUser, big.NewInt(1))
	checkTx(tx, err, "Transfer to newUser")

	result, err := svc.TokenContract.BalanceOf(&bind.CallOpts{}, newUser)
	if err != nil {
		t.Fatalf("Unable to run BalanceOf %v", err)
	}
	if result.Uint64() != 1 {
		t.Fatal("balanceOf newUser should be 1")
	}

	// core -> newUser should be allowed
	code, err := svc.TokenContract.DetectTransferRestriction(&bind.CallOpts{}, genesis, newUser, big.NewInt(1))
	if err != nil {
		t.Fatalf("Unable to run DetectTransferRestriction %v", err)
	}
	if code != 0 {
		t.Fatalf("genesis->newUser expected code to be 0 but is %v", code)
	}

	// newUser -> civilianUser should fail
	code, err = svc.TokenContract.DetectTransferRestriction(&bind.CallOpts{}, newUser, civilianUser, big.NewInt(1))
	if err != nil {
		t.Fatalf("Unable to run DetectTransferRestriction %v", err)
	}
	if code != 1 {
		t.Fatalf("newUser->civilianUser expected code to be 1 but is %v", code)
	}

	// civilianUser -> newUser should fail since they haven't proved use
	code, err = svc.TokenContract.DetectTransferRestriction(&bind.CallOpts{}, civilianUser, newUser, big.NewInt(1))
	if err != nil {
		t.Fatalf("Unable to run DetectTransferRestriction %v", err)
	}
	if code != 2 {
		t.Fatalf("civilianUser->newUser expected code to be 2 but is %v", code)
	}

	// unlockedUser -> newUser should be allowed
	code, err = svc.TokenContract.DetectTransferRestriction(&bind.CallOpts{}, unlockedUser, newUser, big.NewInt(1))
	if err != nil {
		t.Fatalf("Unable to run DetectTransferRestriction %v", err)
	}
	if code != 0 {
		t.Fatalf("unlockedUser->newUser expected code to be 0 but is %v", code)
	}
}

func deployCVLToken(tokenController common.Address, helper *eth.Helper) (common.Address, error) {
	_initialAmount := &big.Int{}
	_initialAmount, _ = _initialAmount.SetString("1000000000000000000000000", 10)
	_tokenName := "TestCVL"
	_decimalUnits := uint8(18)
	_tokenSymbol := "TESTCVL"

	addr, _, _, err := contract.DeployCVLTokenContract(helper.Auth, helper.Blockchain, _initialAmount, _tokenName, _decimalUnits, _tokenSymbol, tokenController)
	if err != nil {
		return common.Address{}, err
	}

	return addr, nil
}

func deployCivilTokenController(helper *eth.Helper) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	libAddress, _, _, err := contract.DeployMessagesAndCodesContract(helper.Auth, helper.Blockchain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	return eth.DeployContractWithLinks(
		helper.Auth, helper.Blockchain,
		contract.CivilTokenControllerContractABI,
		contract.CivilTokenControllerContractBin,
		map[string]common.Address{"MessagesAndCodes": libAddress},
	)

}
