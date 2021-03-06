// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultiSigWalletFactoryContractABI is the input ABI used to generate the binding from.
const MultiSigWalletFactoryContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isInstantiation\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"instantiations\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"getInstantiationCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"instantiation\",\"type\":\"address\"}],\"name\":\"ContractInstantiation\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[{\"name\":\"wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MultiSigWalletFactoryContractBin is the compiled bytecode used for deploying new contracts.
const MultiSigWalletFactoryContractBin = `0x608060405234801561001057600080fd5b506119cb806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632f4f3316811461006657806357183c821461009b5780638f838478146100db578063f8f738081461010e575b600080fd5b34801561007257600080fd5b50610087600160a060020a0360043516610165565b604080519115158252519081900360200190f35b3480156100a757600080fd5b506100bf600160a060020a036004351660243561017a565b60408051600160a060020a039092168252519081900360200190f35b3480156100e757600080fd5b506100fc600160a060020a03600435166101b1565b60408051918252519081900360200190f35b34801561011a57600080fd5b50604080516020600480358082013583810280860185019096528085526100bf9536959394602494938501929182918501908490808284375094975050933594506101cc9350505050565b60006020819052908152604090205460ff1681565b60016020528160005260406000208181548110151561019557fe5b600091825260209091200154600160a060020a03169150829050565b600160a060020a031660009081526001602052604090205490565b600082826101d86102e9565b60208082018390526040808352845190830152835182916060830191868201910280838360005b838110156102175781810151838201526020016101ff565b505050509050019350505050604051809103906000f08015801561023f573d6000803e3d6000fd5b50905061024b81610251565b92915050565b600160a060020a038116600081815260208181526040808320805460ff19166001908117909155338085528184528285208054928301815585529383902001805473ffffffffffffffffffffffffffffffffffffffff19168517905580519283529082019290925281517f4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161929181900390910190a150565b6040516116a6806102fa83390190560060806040523480156200001157600080fd5b50604051620016a6380380620016a68339810160405280516020820151910180519091906000908260328211156200004857600080fd5b818111156200005657600080fd5b8015156200006357600080fd5b8115156200007057600080fd5b600092505b845183101562000149576002600086858151811015156200009257fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff16158015620000e957508451600090869085908110620000d357fe5b90602001906020020151600160a060020a031614155b1515620000f557600080fd5b60016002600087868151811015156200010a57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790556001929092019162000075565b84516200015e90600390602088019062000170565b50505060049190915550620002049050565b828054828255906000526020600020908101928215620001c8579160200282015b82811115620001c85782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019062000191565b50620001d6929150620001da565b5090565b6200020191905b80821115620001d6578054600160a060020a0319168155600101620001e1565b90565b61149280620002146000396000f30060806040526004361061011c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c27811461015e578063173825d91461019257806320ea8d86146101b35780632f54bf6e146101cb5780633411c81c1461020057806354741525146102245780637065cb4814610255578063784547a7146102765780638b51d13f1461028e5780639ace38c2146102a6578063a0e67e2b14610361578063a8abe69a146103c6578063b5dc40c3146103eb578063b77bf60014610403578063ba51a6df14610418578063c01a8c8414610430578063c642747414610448578063d74f8edd146104b1578063dc8452cd146104c6578063e20056e6146104db578063ee22610b14610502575b600034111561015c5760408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a25b005b34801561016a57600080fd5b5061017660043561051a565b60408051600160a060020a039092168252519081900360200190f35b34801561019e57600080fd5b5061015c600160a060020a0360043516610542565b3480156101bf57600080fd5b5061015c6004356106b9565b3480156101d757600080fd5b506101ec600160a060020a0360043516610773565b604080519115158252519081900360200190f35b34801561020c57600080fd5b506101ec600435600160a060020a0360243516610788565b34801561023057600080fd5b50610243600435151560243515156107a8565b60408051918252519081900360200190f35b34801561026157600080fd5b5061015c600160a060020a0360043516610814565b34801561028257600080fd5b506101ec60043561093a565b34801561029a57600080fd5b506102436004356109be565b3480156102b257600080fd5b506102be600435610a2d565b6040518085600160a060020a0316600160a060020a031681526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b8381101561032357818101518382015260200161030b565b50505050905090810190601f1680156103505780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561036d57600080fd5b50610376610aeb565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103b257818101518382015260200161039a565b505050509050019250505060405180910390f35b3480156103d257600080fd5b5061037660043560243560443515156064351515610b4e565b3480156103f757600080fd5b50610376600435610c87565b34801561040f57600080fd5b50610243610e00565b34801561042457600080fd5b5061015c600435610e06565b34801561043c57600080fd5b5061015c600435610e86565b34801561045457600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610243948235600160a060020a0316946024803595369594606494920191908190840183828082843750949750610f519650505050505050565b3480156104bd57600080fd5b50610243610f70565b3480156104d257600080fd5b50610243610f75565b3480156104e757600080fd5b5061015c600160a060020a0360043581169060243516610f7b565b34801561050e57600080fd5b5061015c600435611105565b600380548290811061052857fe5b600091825260209091200154600160a060020a0316905081565b600033301461055057600080fd5b600160a060020a038216600090815260026020526040902054829060ff16151561057957600080fd5b600160a060020a0383166000908152600260205260408120805460ff1916905591505b600354600019018210156106545782600160a060020a03166003838154811015156105c357fe5b600091825260209091200154600160a060020a03161415610649576003805460001981019081106105f057fe5b60009182526020909120015460038054600160a060020a03909216918490811061061657fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610654565b60019091019061059c565b60038054600019019061066790826113a5565b5060035460045411156106805760035461068090610e06565b604051600160a060020a038416907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9090600090a2505050565b3360008181526002602052604090205460ff1615156106d757600080fd5b60008281526001602090815260408083203380855292529091205483919060ff16151561070357600080fd5b600084815260208190526040902060030154849060ff161561072457600080fd5b6000858152600160209081526040808320338085529252808320805460ff191690555187927ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e991a35050505050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b6000805b60055481101561080d578380156107d5575060008181526020819052604090206003015460ff16155b806107f957508280156107f9575060008181526020819052604090206003015460ff165b15610805576001820191505b6001016107ac565b5092915050565b33301461082057600080fd5b600160a060020a038116600090815260026020526040902054819060ff161561084857600080fd5b81600160a060020a038116151561085e57600080fd5b600354600454600190910190603282111561087857600080fd5b8181111561088557600080fd5b80151561089157600080fd5b81151561089d57600080fd5b600160a060020a038516600081815260026020526040808220805460ff1916600190811790915560038054918201815583527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff191684179055517ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d9190a25050505050565b600080805b6003548110156109b7576000848152600160205260408120600380549192918490811061096857fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff161561099c576001820191505b6004548214156109af57600192506109b7565b60010161093f565b5050919050565b6000805b600354811015610a2757600083815260016020526040812060038054919291849081106109eb57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610a1f576001820191505b6001016109c2565b50919050565b6000602081815291815260409081902080546001808301546002808501805487516101009582161595909502600019011691909104601f8101889004880284018801909652858352600160a060020a0390931695909491929190830182828015610ad85780601f10610aad57610100808354040283529160200191610ad8565b820191906000526020600020905b815481529060010190602001808311610abb57829003601f168201915b5050506003909301549192505060ff1684565b60606003805480602002602001604051908101604052809291908181526020018280548015610b4357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610b25575b505050505090505b90565b606080600080600554604051908082528060200260200182016040528015610b80578160200160208202803883390190505b50925060009150600090505b600554811015610c0757858015610bb5575060008181526020819052604090206003015460ff16155b80610bd95750848015610bd9575060008181526020819052604090206003015460ff165b15610bff57808383815181101515610bed57fe5b60209081029091010152600191909101905b600101610b8c565b878703604051908082528060200260200182016040528015610c33578160200160208202803883390190505b5093508790505b86811015610c7c578281815181101515610c5057fe5b9060200190602002015184898303815181101515610c6a57fe5b60209081029091010152600101610c3a565b505050949350505050565b606080600080600380549050604051908082528060200260200182016040528015610cbc578160200160208202803883390190505b50925060009150600090505b600354811015610d795760008581526001602052604081206003805491929184908110610cf157fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610d71576003805482908110610d2c57fe5b6000918252602090912001548351600160a060020a0390911690849084908110610d5257fe5b600160a060020a03909216602092830290910190910152600191909101905b600101610cc8565b81604051908082528060200260200182016040528015610da3578160200160208202803883390190505b509350600090505b81811015610df8578281815181101515610dc157fe5b906020019060200201518482815181101515610dd957fe5b600160a060020a03909216602092830290910190910152600101610dab565b505050919050565b60055481565b333014610e1257600080fd5b600354816032821115610e2457600080fd5b81811115610e3157600080fd5b801515610e3d57600080fd5b811515610e4957600080fd5b60048390556040805184815290517fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a9181900360200190a1505050565b3360008181526002602052604090205460ff161515610ea457600080fd5b6000828152602081905260409020548290600160a060020a03161515610ec957600080fd5b60008381526001602090815260408083203380855292529091205484919060ff1615610ef457600080fd5b6000858152600160208181526040808420338086529252808420805460ff1916909317909255905187927f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef91a3610f4a85611105565b5050505050565b6000610f5e8484846112b5565b9050610f6981610e86565b9392505050565b603281565b60045481565b6000333014610f8957600080fd5b600160a060020a038316600090815260026020526040902054839060ff161515610fb257600080fd5b600160a060020a038316600090815260026020526040902054839060ff1615610fda57600080fd5b600092505b60035483101561106b5784600160a060020a031660038481548110151561100257fe5b600091825260209091200154600160a060020a03161415611060578360038481548110151561102d57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555061106b565b600190920191610fdf565b600160a060020a03808616600081815260026020526040808220805460ff1990811690915593881682528082208054909416600117909355915190917f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9091a2604051600160a060020a038516907ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d90600090a25050505050565b3360008181526002602052604081205490919060ff16151561112657600080fd5b60008381526001602090815260408083203380855292529091205484919060ff16151561115257600080fd5b600085815260208190526040902060030154859060ff161561117357600080fd5b61117c8661093a565b156112ad576000868152602081905260409081902060038101805460ff19166001908117909155815481830154935160028085018054959b50600160a060020a039093169594929391928392859260001991831615610100029190910190911604801561122a5780601f106111ff5761010080835404028352916020019161122a565b820191906000526020600020905b81548152906001019060200180831161120d57829003601f168201915b505091505060006040518083038185875af192505050156112755760405186907f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7590600090a26112ad565b60405186907f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923690600090a260038501805460ff191690555b505050505050565b600083600160a060020a03811615156112cd57600080fd5b60055460408051608081018252600160a060020a0388811682526020808301898152838501898152600060608601819052878152808452959095208451815473ffffffffffffffffffffffffffffffffffffffff19169416939093178355516001830155925180519496509193909261134d9260028501929101906113ce565b50606091909101516003909101805460ff191691151591909117905560058054600101905560405182907fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5190600090a2509392505050565b8154818355818111156113c9576000838152602090206113c991810190830161144c565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061140f57805160ff191683800117855561143c565b8280016001018555821561143c579182015b8281111561143c578251825591602001919060010190611421565b5061144892915061144c565b5090565b610b4b91905b8082111561144857600081556001016114525600a165627a7a72305820cfa5fa77ab2e109c37b905dd5e865dfd565faea2624d175ccd2d96d48baa9a470029a165627a7a723058201cb391519aebfa46603763eb00fe5f563e3c6b2e0124724a504168baabc998450029`

// DeployMultiSigWalletFactoryContract deploys a new Ethereum contract, binding an instance of MultiSigWalletFactoryContract to it.
func DeployMultiSigWalletFactoryContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MultiSigWalletFactoryContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletFactoryContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiSigWalletFactoryContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSigWalletFactoryContract{MultiSigWalletFactoryContractCaller: MultiSigWalletFactoryContractCaller{contract: contract}, MultiSigWalletFactoryContractTransactor: MultiSigWalletFactoryContractTransactor{contract: contract}, MultiSigWalletFactoryContractFilterer: MultiSigWalletFactoryContractFilterer{contract: contract}}, nil
}

// MultiSigWalletFactoryContract is an auto generated Go binding around an Ethereum contract.
type MultiSigWalletFactoryContract struct {
	MultiSigWalletFactoryContractCaller     // Read-only binding to the contract
	MultiSigWalletFactoryContractTransactor // Write-only binding to the contract
	MultiSigWalletFactoryContractFilterer   // Log filterer for contract events
}

// MultiSigWalletFactoryContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigWalletFactoryContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletFactoryContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigWalletFactoryContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletFactoryContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigWalletFactoryContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletFactoryContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigWalletFactoryContractSession struct {
	Contract     *MultiSigWalletFactoryContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MultiSigWalletFactoryContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigWalletFactoryContractCallerSession struct {
	Contract *MultiSigWalletFactoryContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// MultiSigWalletFactoryContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigWalletFactoryContractTransactorSession struct {
	Contract     *MultiSigWalletFactoryContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// MultiSigWalletFactoryContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigWalletFactoryContractRaw struct {
	Contract *MultiSigWalletFactoryContract // Generic contract binding to access the raw methods on
}

// MultiSigWalletFactoryContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigWalletFactoryContractCallerRaw struct {
	Contract *MultiSigWalletFactoryContractCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigWalletFactoryContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigWalletFactoryContractTransactorRaw struct {
	Contract *MultiSigWalletFactoryContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSigWalletFactoryContract creates a new instance of MultiSigWalletFactoryContract, bound to a specific deployed contract.
func NewMultiSigWalletFactoryContract(address common.Address, backend bind.ContractBackend) (*MultiSigWalletFactoryContract, error) {
	contract, err := bindMultiSigWalletFactoryContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletFactoryContract{MultiSigWalletFactoryContractCaller: MultiSigWalletFactoryContractCaller{contract: contract}, MultiSigWalletFactoryContractTransactor: MultiSigWalletFactoryContractTransactor{contract: contract}, MultiSigWalletFactoryContractFilterer: MultiSigWalletFactoryContractFilterer{contract: contract}}, nil
}

// NewMultiSigWalletFactoryContractCaller creates a new read-only instance of MultiSigWalletFactoryContract, bound to a specific deployed contract.
func NewMultiSigWalletFactoryContractCaller(address common.Address, caller bind.ContractCaller) (*MultiSigWalletFactoryContractCaller, error) {
	contract, err := bindMultiSigWalletFactoryContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletFactoryContractCaller{contract: contract}, nil
}

// NewMultiSigWalletFactoryContractTransactor creates a new write-only instance of MultiSigWalletFactoryContract, bound to a specific deployed contract.
func NewMultiSigWalletFactoryContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigWalletFactoryContractTransactor, error) {
	contract, err := bindMultiSigWalletFactoryContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletFactoryContractTransactor{contract: contract}, nil
}

// NewMultiSigWalletFactoryContractFilterer creates a new log filterer instance of MultiSigWalletFactoryContract, bound to a specific deployed contract.
func NewMultiSigWalletFactoryContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigWalletFactoryContractFilterer, error) {
	contract, err := bindMultiSigWalletFactoryContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletFactoryContractFilterer{contract: contract}, nil
}

// bindMultiSigWalletFactoryContract binds a generic wrapper to an already deployed contract.
func bindMultiSigWalletFactoryContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletFactoryContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletFactoryContract.Contract.MultiSigWalletFactoryContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletFactoryContract.Contract.MultiSigWalletFactoryContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletFactoryContract.Contract.MultiSigWalletFactoryContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletFactoryContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletFactoryContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletFactoryContract.Contract.contract.Transact(opts, method, params...)
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) constant returns(uint256)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractCaller) GetInstantiationCount(opts *bind.CallOpts, creator common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletFactoryContract.contract.Call(opts, out, "getInstantiationCount", creator)
	return *ret0, err
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) constant returns(uint256)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractSession) GetInstantiationCount(creator common.Address) (*big.Int, error) {
	return _MultiSigWalletFactoryContract.Contract.GetInstantiationCount(&_MultiSigWalletFactoryContract.CallOpts, creator)
}

// GetInstantiationCount is a free data retrieval call binding the contract method 0x8f838478.
//
// Solidity: function getInstantiationCount(address creator) constant returns(uint256)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractCallerSession) GetInstantiationCount(creator common.Address) (*big.Int, error) {
	return _MultiSigWalletFactoryContract.Contract.GetInstantiationCount(&_MultiSigWalletFactoryContract.CallOpts, creator)
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) constant returns(address)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractCaller) Instantiations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiSigWalletFactoryContract.contract.Call(opts, out, "instantiations", arg0, arg1)
	return *ret0, err
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) constant returns(address)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractSession) Instantiations(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _MultiSigWalletFactoryContract.Contract.Instantiations(&_MultiSigWalletFactoryContract.CallOpts, arg0, arg1)
}

// Instantiations is a free data retrieval call binding the contract method 0x57183c82.
//
// Solidity: function instantiations(address , uint256 ) constant returns(address)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractCallerSession) Instantiations(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _MultiSigWalletFactoryContract.Contract.Instantiations(&_MultiSigWalletFactoryContract.CallOpts, arg0, arg1)
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) constant returns(bool)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractCaller) IsInstantiation(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletFactoryContract.contract.Call(opts, out, "isInstantiation", arg0)
	return *ret0, err
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) constant returns(bool)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractSession) IsInstantiation(arg0 common.Address) (bool, error) {
	return _MultiSigWalletFactoryContract.Contract.IsInstantiation(&_MultiSigWalletFactoryContract.CallOpts, arg0)
}

// IsInstantiation is a free data retrieval call binding the contract method 0x2f4f3316.
//
// Solidity: function isInstantiation(address ) constant returns(bool)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractCallerSession) IsInstantiation(arg0 common.Address) (bool, error) {
	return _MultiSigWalletFactoryContract.Contract.IsInstantiation(&_MultiSigWalletFactoryContract.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0xf8f73808.
//
// Solidity: function create(address[] _owners, uint256 _required) returns(address wallet)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractTransactor) Create(opts *bind.TransactOpts, _owners []common.Address, _required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletFactoryContract.contract.Transact(opts, "create", _owners, _required)
}

// Create is a paid mutator transaction binding the contract method 0xf8f73808.
//
// Solidity: function create(address[] _owners, uint256 _required) returns(address wallet)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractSession) Create(_owners []common.Address, _required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletFactoryContract.Contract.Create(&_MultiSigWalletFactoryContract.TransactOpts, _owners, _required)
}

// Create is a paid mutator transaction binding the contract method 0xf8f73808.
//
// Solidity: function create(address[] _owners, uint256 _required) returns(address wallet)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractTransactorSession) Create(_owners []common.Address, _required *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletFactoryContract.Contract.Create(&_MultiSigWalletFactoryContract.TransactOpts, _owners, _required)
}

// MultiSigWalletFactoryContractContractInstantiationIterator is returned from FilterContractInstantiation and is used to iterate over the raw logs and unpacked data for ContractInstantiation events raised by the MultiSigWalletFactoryContract contract.
type MultiSigWalletFactoryContractContractInstantiationIterator struct {
	Event *MultiSigWalletFactoryContractContractInstantiation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MultiSigWalletFactoryContractContractInstantiationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletFactoryContractContractInstantiation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MultiSigWalletFactoryContractContractInstantiation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MultiSigWalletFactoryContractContractInstantiationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletFactoryContractContractInstantiationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletFactoryContractContractInstantiation represents a ContractInstantiation event raised by the MultiSigWalletFactoryContract contract.
type MultiSigWalletFactoryContractContractInstantiation struct {
	Sender        common.Address
	Instantiation common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterContractInstantiation is a free log retrieval operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractFilterer) FilterContractInstantiation(opts *bind.FilterOpts) (*MultiSigWalletFactoryContractContractInstantiationIterator, error) {

	logs, sub, err := _MultiSigWalletFactoryContract.contract.FilterLogs(opts, "ContractInstantiation")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletFactoryContractContractInstantiationIterator{contract: _MultiSigWalletFactoryContract.contract, event: "ContractInstantiation", logs: logs, sub: sub}, nil
}

// WatchContractInstantiation is a free log subscription operation binding the contract event 0x4fb057ad4a26ed17a57957fa69c306f11987596069b89521c511fc9a894e6161.
//
// Solidity: event ContractInstantiation(address sender, address instantiation)
func (_MultiSigWalletFactoryContract *MultiSigWalletFactoryContractFilterer) WatchContractInstantiation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletFactoryContractContractInstantiation) (event.Subscription, error) {

	logs, sub, err := _MultiSigWalletFactoryContract.contract.WatchLogs(opts, "ContractInstantiation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletFactoryContractContractInstantiation)
				if err := _MultiSigWalletFactoryContract.contract.UnpackLog(event, "ContractInstantiation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
