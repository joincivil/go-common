package token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/go-common/pkg/generated/contract"
)

// CVLTokenService is a set of helpers to interact with the CVLToken smart contract and token controller
type CVLTokenService struct {
	backend            bind.ContractBackend
	controllerAddress  common.Address
	ControllerContract *contract.CivilTokenControllerContract
	tokenAddress       common.Address
	TokenContract      *contract.CVLTokenContract
}

// NewCVLTokenService constructs a new instance of CVLTokenService
func NewCVLTokenService(backend bind.ContractBackend, controllerAddress common.Address, tokenAddress common.Address) (*CVLTokenService, error) {
	tokenContract, err := contract.NewCVLTokenContract(tokenAddress, backend)
	if err != nil {
		return nil, err
	}

	controllerContract, err := contract.NewCivilTokenControllerContract(controllerAddress, backend)
	if err != nil {
		return nil, err
	}

	return &CVLTokenService{
		backend,
		controllerAddress,
		controllerContract,
		tokenAddress,
		tokenContract,
	}, nil
}
