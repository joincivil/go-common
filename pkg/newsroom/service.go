package newsroom

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joincivil/go-common/pkg/eth"
	"github.com/joincivil/go-common/pkg/generated/contract"
	"github.com/pkg/errors"

	shell "github.com/ipfs/go-ipfs-api"
)

// Service provides methods to create and manage newsrooms
type Service struct {
	eth         *eth.Helper
	ipfs        *shell.Shell
	addresses   eth.DeployerContractAddresses
	cvl         *contract.CVLTokenContract
	cvlABI      abi.ABI
	factory     *contract.CreateNewsroomInGroupContract
	factoryABI  abi.ABI
	tcr         *contract.CivilTCRContract
	tcrABI      abi.ABI
	multisigABI abi.ABI
}

// NewService builds a new newsroom Service
func NewService(eth *eth.Helper, addresses eth.DeployerContractAddresses) (*Service, error) {
	sh := shell.NewShell("https://ipfs.infura.io:5001")

	factory, err := contract.NewCreateNewsroomInGroupContract(addresses.CreateNewsroomInGroup, eth.Blockchain)
	if err != nil {
		return nil, err
	}
	factoryABI, err := abi.JSON(strings.NewReader(contract.NewsroomFactoryABI))
	if err != nil {
		return nil, err
	}

	tcr, err := contract.NewCivilTCRContract(addresses.TCR, eth.Blockchain)
	if err != nil {
		return nil, err
	}
	tcrABI, err := abi.JSON(strings.NewReader(contract.CivilTCRContractABI))
	if err != nil {
		return nil, err
	}

	cvl, err := contract.NewCVLTokenContract(addresses.CVLToken, eth.Blockchain)
	if err != nil {
		return nil, err
	}

	cvlABI, err := abi.JSON(strings.NewReader(contract.CVLTokenContractABI))
	if err != nil {
		return nil, err
	}

	multisigABI, err := abi.JSON(strings.NewReader(contract.MultiSigWalletContractABI))
	if err != nil {
		return nil, err
	}

	return &Service{
		eth:         eth,
		addresses:   addresses,
		factory:     factory,
		factoryABI:  factoryABI,
		tcr:         tcr,
		tcrABI:      tcrABI,
		cvl:         cvl,
		cvlABI:      cvlABI,
		multisigABI: multisigABI,
		ipfs:        sh,
	}, nil
}

// PublishCharter takes a charter and publishes to IPFS
func (s *Service) PublishCharter(charter Charter, pin bool) (string, error) {

	res, err := json.Marshal(charter)
	if err != nil {
		return "", errors.Wrap(err, "Error serializing charter")
	}

	cid, err := s.ipfs.Add(bytes.NewReader(res))
	if err != nil {
		return "", err
	}

	if pin {
		err := s.ipfs.Pin(cid)
		if err != nil {
			return "", err
		}
	}

	return cid, nil
}

// CreateNewsroom creates a new newsroom smart contract
func (s *Service) CreateNewsroom(name string, charterHash string) (common.Hash, error) {

	// set the initial owners on the Multisig to the tx sender
	initialOwners := []common.Address{s.eth.Auth.From}
	charterURI := fmt.Sprintf("ipfs://%v", charterHash)

	// number of confirmations required for the multisig
	initialRequired := big.NewInt(1)

	// convert charterHash to a [32]byte needed in the contract
	var charterHashBytes [32]byte
	copy(charterHashBytes[:], charterHash)

	// create a newsroom using the CreateNewsroomInGroup factory
	tx, err := s.factory.Create(s.eth.Transact(), name, charterURI, charterHashBytes, initialOwners, initialRequired)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "Error creating factory transaction")
	}

	return tx.Hash(), nil
}

// GetNewsroomAddressFromTransaction retrieves the deployed Newsroom address from tx receipt logs
func (s *Service) GetNewsroomAddressFromTransaction(tx common.Hash) (common.Address, error) {
	// get the transaction receipt
	receipt, err := s.eth.Blockchain.(ethereum.TransactionReader).TransactionReceipt(context.Background(), tx)
	if err != nil {
		return common.Address{}, err
	}

	// inspect the transaction receipt logs to find the Newsroom contract address
	var newsroomContract common.Address
	for _, vLog := range receipt.Logs {
		// multisig, newsroom, and factory all create logs, but we are only interested in the NewsroomFactory
		if vLog.Address == s.addresses.NewsroomFactory {
			// struct to unpack the abi into
			event := struct {
				Sender        common.Address
				Instantiation common.Address
			}{}

			// unpack the binary data into the event struct using the "ContractInstantiation" event ABI
			err = s.factoryABI.Unpack(&event, "ContractInstantiation", vLog.Data)
			if err != nil {
				return common.Address{}, err
			}

			if (event.Instantiation != common.Address{}) {
				newsroomContract = event.Instantiation
			}

		}
	}

	return newsroomContract, nil
}

// GetNewsroomName retrieves the name of the newsroom at the provided address
func (s *Service) GetNewsroomName(newsroomAddress common.Address) (string, error) {
	newsroom, err := contract.NewNewsroomContract(newsroomAddress, s.eth.Blockchain)
	if err != nil {
		return "", err
	}
	name, err := newsroom.NewsroomContractCaller.Name(nil)
	if err != nil {
		return "", err
	}

	return name, nil
}

// GetOwner retrieves the address of the owner of the Newsroom
func (s *Service) GetOwner(newsroomAddress common.Address) (common.Address, error) {
	newsroom, err := contract.NewNewsroomContract(newsroomAddress, s.eth.Blockchain)
	if err != nil {
		return common.Address{}, err
	}
	owner, err := newsroom.NewsroomContractCaller.Owner(nil)
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}

// getMultisigForNewsroom returns a Multisig contract instance and the address of the Newsroom owner
func (s *Service) getMultisigForNewsroom(newsroomAddress common.Address) (*contract.MultiSigWalletContract, common.Address, error) {
	owner, err := s.GetOwner(newsroomAddress)
	if err != nil {
		return nil, common.Address{}, err
	}

	multisig, err := contract.NewMultiSigWalletContract(owner, s.eth.Blockchain)
	if err != nil {
		return nil, common.Address{}, err
	}

	return multisig, owner, nil
}

// GetMultisigMembers returns the Addresses on the Multisig associated with the provided newsroom address
func (s *Service) GetMultisigMembers(newsroomAddress common.Address) ([]common.Address, error) {
	multisig, _, err := s.getMultisigForNewsroom(newsroomAddress)
	if err != nil {
		return nil, err
	}

	return multisig.GetOwners(nil)
}

// AdminAddMultisigOwner adds an owner to the newsroom multisig
func (s *Service) AdminAddMultisigOwner(newsroomAddress common.Address, newOwner common.Address) (common.Hash, error) {
	multisig, multisigAddr, err := s.getMultisigForNewsroom(newsroomAddress)
	if err != nil {
		return common.Hash{}, err
	}

	// we can't use `multisig.AddOwner` directly, since the function requires it comes from the multisig
	// instead, we build the []byte for the tx data and then use `multisig.SubmitTransaction` to execute it
	data, err := s.multisigABI.Pack("addOwner", newOwner)
	if err != nil {
		return common.Hash{}, err
	}

	// amounf of ETH to send with the tx
	value := big.NewInt(0)

	// submit the tx to the multisig
	tx, err := multisig.SubmitTransaction(s.eth.Transact(), multisigAddr, value, data)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// AdminRemoveMultisigOwner removes the address from the multisig
func (s *Service) AdminRemoveMultisigOwner(newsroomAddress common.Address, owner common.Address) (common.Hash, error) {
	multisig, multisigAddr, err := s.getMultisigForNewsroom(newsroomAddress)
	if err != nil {
		return common.Hash{}, err
	}

	// we can't use `multisig.RemoveOwner` directly, since the function requires it comes from the multisig
	// instead, we build the []byte for the tx data and then use `multisig.SubmitTransaction` to execute it
	data, err := s.multisigABI.Pack("removeOwner", owner)
	if err != nil {
		return common.Hash{}, err
	}

	// amounf of ETH to send with the tx
	value := big.NewInt(0)

	// submit the tx to the multisig
	tx, err := multisig.SubmitTransaction(s.eth.Transact(), multisigAddr, value, data)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// AdminApproveTCRTokenTransfer approves the specified number of tokens to be transferred by the TCR contract
// this needs to be called before `AdminApplyToTCR` otherwise the tx will fail
func (s *Service) AdminApproveTCRTokenTransfer(newsroomAddress common.Address, tokens *big.Int) (common.Hash, error) {
	multisig, _, err := s.getMultisigForNewsroom(newsroomAddress)
	if err != nil {
		return common.Hash{}, err
	}

	// produce the []byte for the `cvltoken.approve` function call
	approveData, err := s.cvlABI.Pack("approve", s.addresses.TCR, tokens)
	if err != nil {
		return common.Hash{}, err
	}

	// submit the transaction to the multisig to execute, targeting the CVLToken address
	tx, err := multisig.SubmitTransaction(s.eth.Transact(), s.addresses.CVLToken, big.NewInt(0), approveData)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

// AdminApplyToTCR applies the newsroom to the TCR
func (s *Service) AdminApplyToTCR(newsroomAddress common.Address, tokens *big.Int) (common.Hash, error) {
	multisig, _, err := s.getMultisigForNewsroom(newsroomAddress)
	if err != nil {
		return common.Hash{}, err
	}

	// produce the []byte for the `tcr.apply` function call
	data, err := s.tcrABI.Pack("apply", newsroomAddress, tokens, "")
	if err != nil {
		return common.Hash{}, err
	}

	// submit the transaction to the multisig to execute, targeting the tcr address
	tcrTx, err := multisig.SubmitTransaction(s.eth.Transact(), s.addresses.TCR, big.NewInt(0), data)
	if err != nil {
		return common.Hash{}, err
	}

	return tcrTx.Hash(), nil
}
