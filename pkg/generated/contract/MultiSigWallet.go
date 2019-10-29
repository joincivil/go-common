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

// MultiSigWalletContractABI is the input ABI used to generate the binding from.
const MultiSigWalletContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialOwners\",\"type\":\"address[]\"},{\"name\":\"initialRequired\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRequired\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MultiSigWalletContractBin is the compiled bytecode used for deploying new contracts.
const MultiSigWalletContractBin = `0x60806040523480156200001157600080fd5b50604051620016a6380380620016a68339810160405280516020820151910180519091906000908260328211156200004857600080fd5b818111156200005657600080fd5b8015156200006357600080fd5b8115156200007057600080fd5b600092505b845183101562000149576002600086858151811015156200009257fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff16158015620000e957508451600090869085908110620000d357fe5b90602001906020020151600160a060020a031614155b1515620000f557600080fd5b60016002600087868151811015156200010a57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790556001929092019162000075565b84516200015e90600390602088019062000170565b50505060049190915550620002049050565b828054828255906000526020600020908101928215620001c8579160200282015b82811115620001c85782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019062000191565b50620001d6929150620001da565b5090565b6200020191905b80821115620001d6578054600160a060020a0319168155600101620001e1565b90565b61149280620002146000396000f30060806040526004361061011c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c27811461015e578063173825d91461019257806320ea8d86146101b35780632f54bf6e146101cb5780633411c81c1461020057806354741525146102245780637065cb4814610255578063784547a7146102765780638b51d13f1461028e5780639ace38c2146102a6578063a0e67e2b14610361578063a8abe69a146103c6578063b5dc40c3146103eb578063b77bf60014610403578063ba51a6df14610418578063c01a8c8414610430578063c642747414610448578063d74f8edd146104b1578063dc8452cd146104c6578063e20056e6146104db578063ee22610b14610502575b600034111561015c5760408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a25b005b34801561016a57600080fd5b5061017660043561051a565b60408051600160a060020a039092168252519081900360200190f35b34801561019e57600080fd5b5061015c600160a060020a0360043516610542565b3480156101bf57600080fd5b5061015c6004356106b9565b3480156101d757600080fd5b506101ec600160a060020a0360043516610773565b604080519115158252519081900360200190f35b34801561020c57600080fd5b506101ec600435600160a060020a0360243516610788565b34801561023057600080fd5b50610243600435151560243515156107a8565b60408051918252519081900360200190f35b34801561026157600080fd5b5061015c600160a060020a0360043516610814565b34801561028257600080fd5b506101ec60043561093a565b34801561029a57600080fd5b506102436004356109be565b3480156102b257600080fd5b506102be600435610a2d565b6040518085600160a060020a0316600160a060020a031681526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b8381101561032357818101518382015260200161030b565b50505050905090810190601f1680156103505780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561036d57600080fd5b50610376610aeb565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103b257818101518382015260200161039a565b505050509050019250505060405180910390f35b3480156103d257600080fd5b5061037660043560243560443515156064351515610b4e565b3480156103f757600080fd5b50610376600435610c87565b34801561040f57600080fd5b50610243610e00565b34801561042457600080fd5b5061015c600435610e06565b34801561043c57600080fd5b5061015c600435610e86565b34801561045457600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610243948235600160a060020a0316946024803595369594606494920191908190840183828082843750949750610f519650505050505050565b3480156104bd57600080fd5b50610243610f70565b3480156104d257600080fd5b50610243610f75565b3480156104e757600080fd5b5061015c600160a060020a0360043581169060243516610f7b565b34801561050e57600080fd5b5061015c600435611105565b600380548290811061052857fe5b600091825260209091200154600160a060020a0316905081565b600033301461055057600080fd5b600160a060020a038216600090815260026020526040902054829060ff16151561057957600080fd5b600160a060020a0383166000908152600260205260408120805460ff1916905591505b600354600019018210156106545782600160a060020a03166003838154811015156105c357fe5b600091825260209091200154600160a060020a03161415610649576003805460001981019081106105f057fe5b60009182526020909120015460038054600160a060020a03909216918490811061061657fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610654565b60019091019061059c565b60038054600019019061066790826113a5565b5060035460045411156106805760035461068090610e06565b604051600160a060020a038416907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9090600090a2505050565b3360008181526002602052604090205460ff1615156106d757600080fd5b60008281526001602090815260408083203380855292529091205483919060ff16151561070357600080fd5b600084815260208190526040902060030154849060ff161561072457600080fd5b6000858152600160209081526040808320338085529252808320805460ff191690555187927ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e991a35050505050565b60026020526000908152604090205460ff1681565b600160209081526000928352604080842090915290825290205460ff1681565b6000805b60055481101561080d578380156107d5575060008181526020819052604090206003015460ff16155b806107f957508280156107f9575060008181526020819052604090206003015460ff165b15610805576001820191505b6001016107ac565b5092915050565b33301461082057600080fd5b600160a060020a038116600090815260026020526040902054819060ff161561084857600080fd5b81600160a060020a038116151561085e57600080fd5b600354600454600190910190603282111561087857600080fd5b8181111561088557600080fd5b80151561089157600080fd5b81151561089d57600080fd5b600160a060020a038516600081815260026020526040808220805460ff1916600190811790915560038054918201815583527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff191684179055517ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d9190a25050505050565b600080805b6003548110156109b7576000848152600160205260408120600380549192918490811061096857fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff161561099c576001820191505b6004548214156109af57600192506109b7565b60010161093f565b5050919050565b6000805b600354811015610a2757600083815260016020526040812060038054919291849081106109eb57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610a1f576001820191505b6001016109c2565b50919050565b6000602081815291815260409081902080546001808301546002808501805487516101009582161595909502600019011691909104601f8101889004880284018801909652858352600160a060020a0390931695909491929190830182828015610ad85780601f10610aad57610100808354040283529160200191610ad8565b820191906000526020600020905b815481529060010190602001808311610abb57829003601f168201915b5050506003909301549192505060ff1684565b60606003805480602002602001604051908101604052809291908181526020018280548015610b4357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610b25575b505050505090505b90565b606080600080600554604051908082528060200260200182016040528015610b80578160200160208202803883390190505b50925060009150600090505b600554811015610c0757858015610bb5575060008181526020819052604090206003015460ff16155b80610bd95750848015610bd9575060008181526020819052604090206003015460ff165b15610bff57808383815181101515610bed57fe5b60209081029091010152600191909101905b600101610b8c565b878703604051908082528060200260200182016040528015610c33578160200160208202803883390190505b5093508790505b86811015610c7c578281815181101515610c5057fe5b9060200190602002015184898303815181101515610c6a57fe5b60209081029091010152600101610c3a565b505050949350505050565b606080600080600380549050604051908082528060200260200182016040528015610cbc578160200160208202803883390190505b50925060009150600090505b600354811015610d795760008581526001602052604081206003805491929184908110610cf157fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615610d71576003805482908110610d2c57fe5b6000918252602090912001548351600160a060020a0390911690849084908110610d5257fe5b600160a060020a03909216602092830290910190910152600191909101905b600101610cc8565b81604051908082528060200260200182016040528015610da3578160200160208202803883390190505b509350600090505b81811015610df8578281815181101515610dc157fe5b906020019060200201518482815181101515610dd957fe5b600160a060020a03909216602092830290910190910152600101610dab565b505050919050565b60055481565b333014610e1257600080fd5b600354816032821115610e2457600080fd5b81811115610e3157600080fd5b801515610e3d57600080fd5b811515610e4957600080fd5b60048390556040805184815290517fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a9181900360200190a1505050565b3360008181526002602052604090205460ff161515610ea457600080fd5b6000828152602081905260409020548290600160a060020a03161515610ec957600080fd5b60008381526001602090815260408083203380855292529091205484919060ff1615610ef457600080fd5b6000858152600160208181526040808420338086529252808420805460ff1916909317909255905187927f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef91a3610f4a85611105565b5050505050565b6000610f5e8484846112b5565b9050610f6981610e86565b9392505050565b603281565b60045481565b6000333014610f8957600080fd5b600160a060020a038316600090815260026020526040902054839060ff161515610fb257600080fd5b600160a060020a038316600090815260026020526040902054839060ff1615610fda57600080fd5b600092505b60035483101561106b5784600160a060020a031660038481548110151561100257fe5b600091825260209091200154600160a060020a03161415611060578360038481548110151561102d57fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a0316021790555061106b565b600190920191610fdf565b600160a060020a03808616600081815260026020526040808220805460ff1990811690915593881682528082208054909416600117909355915190917f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9091a2604051600160a060020a038516907ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d90600090a25050505050565b3360008181526002602052604081205490919060ff16151561112657600080fd5b60008381526001602090815260408083203380855292529091205484919060ff16151561115257600080fd5b600085815260208190526040902060030154859060ff161561117357600080fd5b61117c8661093a565b156112ad576000868152602081905260409081902060038101805460ff19166001908117909155815481830154935160028085018054959b50600160a060020a039093169594929391928392859260001991831615610100029190910190911604801561122a5780601f106111ff5761010080835404028352916020019161122a565b820191906000526020600020905b81548152906001019060200180831161120d57829003601f168201915b505091505060006040518083038185875af192505050156112755760405186907f33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed7590600090a26112ad565b60405186907f526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b7923690600090a260038501805460ff191690555b505050505050565b600083600160a060020a03811615156112cd57600080fd5b60055460408051608081018252600160a060020a0388811682526020808301898152838501898152600060608601819052878152808452959095208451815473ffffffffffffffffffffffffffffffffffffffff19169416939093178355516001830155925180519496509193909261134d9260028501929101906113ce565b50606091909101516003909101805460ff191691151591909117905560058054600101905560405182907fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5190600090a2509392505050565b8154818355818111156113c9576000838152602090206113c991810190830161144c565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061140f57805160ff191683800117855561143c565b8280016001018555821561143c579182015b8281111561143c578251825591602001919060010190611421565b5061144892915061144c565b5090565b610b4b91905b8082111561144857600081556001016114525600a165627a7a72305820cfa5fa77ab2e109c37b905dd5e865dfd565faea2624d175ccd2d96d48baa9a470029`

// DeployMultiSigWalletContract deploys a new Ethereum contract, binding an instance of MultiSigWalletContract to it.
func DeployMultiSigWalletContract(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwners []common.Address, initialRequired *big.Int) (common.Address, *types.Transaction, *MultiSigWalletContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MultiSigWalletContractBin), backend, initialOwners, initialRequired)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiSigWalletContract{MultiSigWalletContractCaller: MultiSigWalletContractCaller{contract: contract}, MultiSigWalletContractTransactor: MultiSigWalletContractTransactor{contract: contract}, MultiSigWalletContractFilterer: MultiSigWalletContractFilterer{contract: contract}}, nil
}

// MultiSigWalletContract is an auto generated Go binding around an Ethereum contract.
type MultiSigWalletContract struct {
	MultiSigWalletContractCaller     // Read-only binding to the contract
	MultiSigWalletContractTransactor // Write-only binding to the contract
	MultiSigWalletContractFilterer   // Log filterer for contract events
}

// MultiSigWalletContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigWalletContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigWalletContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigWalletContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigWalletContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigWalletContractSession struct {
	Contract     *MultiSigWalletContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MultiSigWalletContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigWalletContractCallerSession struct {
	Contract *MultiSigWalletContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MultiSigWalletContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigWalletContractTransactorSession struct {
	Contract     *MultiSigWalletContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MultiSigWalletContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigWalletContractRaw struct {
	Contract *MultiSigWalletContract // Generic contract binding to access the raw methods on
}

// MultiSigWalletContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigWalletContractCallerRaw struct {
	Contract *MultiSigWalletContractCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigWalletContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigWalletContractTransactorRaw struct {
	Contract *MultiSigWalletContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSigWalletContract creates a new instance of MultiSigWalletContract, bound to a specific deployed contract.
func NewMultiSigWalletContract(address common.Address, backend bind.ContractBackend) (*MultiSigWalletContract, error) {
	contract, err := bindMultiSigWalletContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContract{MultiSigWalletContractCaller: MultiSigWalletContractCaller{contract: contract}, MultiSigWalletContractTransactor: MultiSigWalletContractTransactor{contract: contract}, MultiSigWalletContractFilterer: MultiSigWalletContractFilterer{contract: contract}}, nil
}

// NewMultiSigWalletContractCaller creates a new read-only instance of MultiSigWalletContract, bound to a specific deployed contract.
func NewMultiSigWalletContractCaller(address common.Address, caller bind.ContractCaller) (*MultiSigWalletContractCaller, error) {
	contract, err := bindMultiSigWalletContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractCaller{contract: contract}, nil
}

// NewMultiSigWalletContractTransactor creates a new write-only instance of MultiSigWalletContract, bound to a specific deployed contract.
func NewMultiSigWalletContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigWalletContractTransactor, error) {
	contract, err := bindMultiSigWalletContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractTransactor{contract: contract}, nil
}

// NewMultiSigWalletContractFilterer creates a new log filterer instance of MultiSigWalletContract, bound to a specific deployed contract.
func NewMultiSigWalletContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigWalletContractFilterer, error) {
	contract, err := bindMultiSigWalletContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractFilterer{contract: contract}, nil
}

// bindMultiSigWalletContract binds a generic wrapper to an already deployed contract.
func bindMultiSigWalletContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigWalletContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletContract *MultiSigWalletContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletContract.Contract.MultiSigWalletContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletContract *MultiSigWalletContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.MultiSigWalletContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletContract *MultiSigWalletContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.MultiSigWalletContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSigWalletContract *MultiSigWalletContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MultiSigWalletContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSigWalletContract *MultiSigWalletContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSigWalletContract *MultiSigWalletContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWalletContract.Contract.MAXOWNERCOUNT(&_MultiSigWalletContract.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSigWalletContract.Contract.MAXOWNERCOUNT(&_MultiSigWalletContract.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWalletContract.Contract.Confirmations(&_MultiSigWalletContract.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSigWalletContract.Contract.Confirmations(&_MultiSigWalletContract.CallOpts, arg0, arg1)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWalletContract *MultiSigWalletContractSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWalletContract.Contract.GetConfirmationCount(&_MultiSigWalletContract.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) constant returns(uint256 count)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSigWalletContract.Contract.GetConfirmationCount(&_MultiSigWalletContract.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWalletContract *MultiSigWalletContractSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWalletContract.Contract.GetConfirmations(&_MultiSigWalletContract.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) constant returns(address[] _confirmations)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSigWalletContract.Contract.GetConfirmations(&_MultiSigWalletContract.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWalletContract *MultiSigWalletContractCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWalletContract *MultiSigWalletContractSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWalletContract.Contract.GetOwners(&_MultiSigWalletContract.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSigWalletContract.Contract.GetOwners(&_MultiSigWalletContract.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWalletContract *MultiSigWalletContractSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWalletContract.Contract.GetTransactionCount(&_MultiSigWalletContract.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) constant returns(uint256 count)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSigWalletContract.Contract.GetTransactionCount(&_MultiSigWalletContract.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWalletContract *MultiSigWalletContractSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWalletContract.Contract.GetTransactionIds(&_MultiSigWalletContract.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) constant returns(uint256[] _transactionIds)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSigWalletContract.Contract.GetTransactionIds(&_MultiSigWalletContract.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWalletContract.Contract.IsConfirmed(&_MultiSigWalletContract.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSigWalletContract.Contract.IsConfirmed(&_MultiSigWalletContract.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWalletContract.Contract.IsOwner(&_MultiSigWalletContract.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) constant returns(bool)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSigWalletContract.Contract.IsOwner(&_MultiSigWalletContract.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWalletContract *MultiSigWalletContractSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWalletContract.Contract.Owners(&_MultiSigWalletContract.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) constant returns(address)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSigWalletContract.Contract.Owners(&_MultiSigWalletContract.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractSession) Required() (*big.Int, error) {
	return _MultiSigWalletContract.Contract.Required(&_MultiSigWalletContract.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) Required() (*big.Int, error) {
	return _MultiSigWalletContract.Contract.Required(&_MultiSigWalletContract.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MultiSigWalletContract.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWalletContract.Contract.TransactionCount(&_MultiSigWalletContract.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) TransactionCount() (*big.Int, error) {
	return _MultiSigWalletContract.Contract.TransactionCount(&_MultiSigWalletContract.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletContract *MultiSigWalletContractCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _MultiSigWalletContract.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletContract *MultiSigWalletContractSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWalletContract.Contract.Transactions(&_MultiSigWalletContract.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) constant returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSigWalletContract *MultiSigWalletContractCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSigWalletContract.Contract.Transactions(&_MultiSigWalletContract.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.AddOwner(&_MultiSigWalletContract.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.AddOwner(&_MultiSigWalletContract.TransactOpts, owner)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 newRequired) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactor) ChangeRequirement(opts *bind.TransactOpts, newRequired *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.contract.Transact(opts, "changeRequirement", newRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 newRequired) returns()
func (_MultiSigWalletContract *MultiSigWalletContractSession) ChangeRequirement(newRequired *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.ChangeRequirement(&_MultiSigWalletContract.TransactOpts, newRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 newRequired) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactorSession) ChangeRequirement(newRequired *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.ChangeRequirement(&_MultiSigWalletContract.TransactOpts, newRequired)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.ConfirmTransaction(&_MultiSigWalletContract.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.ConfirmTransaction(&_MultiSigWalletContract.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.ExecuteTransaction(&_MultiSigWalletContract.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.ExecuteTransaction(&_MultiSigWalletContract.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.RemoveOwner(&_MultiSigWalletContract.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.RemoveOwner(&_MultiSigWalletContract.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.ReplaceOwner(&_MultiSigWalletContract.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.ReplaceOwner(&_MultiSigWalletContract.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.RevokeConfirmation(&_MultiSigWalletContract.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSigWalletContract *MultiSigWalletContractTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.RevokeConfirmation(&_MultiSigWalletContract.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletContract.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.SubmitTransaction(&_MultiSigWalletContract.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSigWalletContract.Contract.SubmitTransaction(&_MultiSigWalletContract.TransactOpts, destination, value, data)
}

// MultiSigWalletContractConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractConfirmationIterator struct {
	Event *MultiSigWalletContractConfirmation // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractConfirmation)
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
		it.Event = new(MultiSigWalletContractConfirmation)
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
func (it *MultiSigWalletContractConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractConfirmation represents a Confirmation event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletContractConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractConfirmationIterator{contract: _MultiSigWalletContract.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractConfirmation)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "Confirmation", log); err != nil {
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

// MultiSigWalletContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractDepositIterator struct {
	Event *MultiSigWalletContractDeposit // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractDeposit)
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
		it.Event = new(MultiSigWalletContractDeposit)
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
func (it *MultiSigWalletContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractDeposit represents a Deposit event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigWalletContractDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractDepositIterator{contract: _MultiSigWalletContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractDeposit)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// MultiSigWalletContractExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractExecutionIterator struct {
	Event *MultiSigWalletContractExecution // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractExecution)
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
		it.Event = new(MultiSigWalletContractExecution)
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
func (it *MultiSigWalletContractExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractExecution represents a Execution event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletContractExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractExecutionIterator{contract: _MultiSigWalletContract.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractExecution)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "Execution", log); err != nil {
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

// MultiSigWalletContractExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractExecutionFailureIterator struct {
	Event *MultiSigWalletContractExecutionFailure // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractExecutionFailure)
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
		it.Event = new(MultiSigWalletContractExecutionFailure)
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
func (it *MultiSigWalletContractExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractExecutionFailure represents a ExecutionFailure event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletContractExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractExecutionFailureIterator{contract: _MultiSigWalletContract.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractExecutionFailure)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// MultiSigWalletContractOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractOwnerAdditionIterator struct {
	Event *MultiSigWalletContractOwnerAddition // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractOwnerAddition)
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
		it.Event = new(MultiSigWalletContractOwnerAddition)
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
func (it *MultiSigWalletContractOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractOwnerAddition represents a OwnerAddition event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletContractOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractOwnerAdditionIterator{contract: _MultiSigWalletContract.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractOwnerAddition)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
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

// MultiSigWalletContractOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractOwnerRemovalIterator struct {
	Event *MultiSigWalletContractOwnerRemoval // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractOwnerRemoval)
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
		it.Event = new(MultiSigWalletContractOwnerRemoval)
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
func (it *MultiSigWalletContractOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractOwnerRemoval represents a OwnerRemoval event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MultiSigWalletContractOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractOwnerRemovalIterator{contract: _MultiSigWalletContract.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractOwnerRemoval)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
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

// MultiSigWalletContractRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractRequirementChangeIterator struct {
	Event *MultiSigWalletContractRequirementChange // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractRequirementChange)
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
		it.Event = new(MultiSigWalletContractRequirementChange)
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
func (it *MultiSigWalletContractRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractRequirementChange represents a RequirementChange event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MultiSigWalletContractRequirementChangeIterator, error) {

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractRequirementChangeIterator{contract: _MultiSigWalletContract.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractRequirementChange) (event.Subscription, error) {

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractRequirementChange)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "RequirementChange", log); err != nil {
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

// MultiSigWalletContractRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractRevocationIterator struct {
	Event *MultiSigWalletContractRevocation // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractRevocation)
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
		it.Event = new(MultiSigWalletContractRevocation)
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
func (it *MultiSigWalletContractRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractRevocation represents a Revocation event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigWalletContractRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractRevocationIterator{contract: _MultiSigWalletContract.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractRevocation)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// MultiSigWalletContractSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MultiSigWalletContract contract.
type MultiSigWalletContractSubmissionIterator struct {
	Event *MultiSigWalletContractSubmission // Event containing the contract specifics and raw log

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
func (it *MultiSigWalletContractSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigWalletContractSubmission)
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
		it.Event = new(MultiSigWalletContractSubmission)
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
func (it *MultiSigWalletContractSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigWalletContractSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigWalletContractSubmission represents a Submission event raised by the MultiSigWalletContract contract.
type MultiSigWalletContractSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigWalletContractSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigWalletContractSubmissionIterator{contract: _MultiSigWalletContract.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSigWalletContract *MultiSigWalletContractFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultiSigWalletContractSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSigWalletContract.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigWalletContractSubmission)
				if err := _MultiSigWalletContract.contract.UnpackLog(event, "Submission", log); err != nil {
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
