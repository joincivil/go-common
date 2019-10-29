// Code generated by 'gen/libgen'  DO NOT EDIT.
// IT SHOULD NOT BE EDITED BY HAND AS ANY CHANGES MAY BE OVERWRITTEN
// Please reference 'gen/libgen' for more details
// File was generated at 2019-10-29 14:14:40.825227 +0000 UTC
package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const MessagesAndCodesContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EMPTY_MESSAGE_ERROR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CODE_UNASSIGNED_ERROR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CODE_RESERVED_ERROR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_code\",\"type\":\"uint8\"},{\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"addMessage\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"autoAddMessage\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_code\",\"type\":\"uint8\"}],\"name\":\"removeMessage\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"self\",\"type\":\"string\"},{\"name\":\"_code\",\"type\":\"uint8\"},{\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"updateMessage\",\"outputs\":[{\"name\":\"code\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

const MessagesAndCodesContractBin = `0x610a88610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600436106100995763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416636ef3054c811461009e57806379bb4acc14610117578063832a10ef14610194578063d7118140146101f7578063d87d728914610215578063ef0067941461021d578063f7feed761461027b575b600080fd5b8180156100aa57600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610101948235946024803560ff16953695946064949201919081908401838280828437509497506102839650505050505050565b6040805160ff9092168252519081900360200190f35b61011f610456565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610159578181015183820152602001610141565b50505050905090810190601f1680156101865780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b8180156101a057600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610101948235946024803560ff169536959460649492019190819084018382808284375094975061047b9650505050505050565b81801561020357600080fd5b5061010160043560ff602435166105de565b61011f6107cc565b81801561022957600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261010195833595369560449491939091019190819084018382808284375094975061082c9650505050505050565b61011f6108d9565b600061028e82610939565b60408051808201909152601e8152600080516020610a3d8339815191526020820152901561033d5760405160e560020a62461bcd0281526004018080602001828103825283818151815260200191508051906020019080838360005b838110156103025781810151838201526020016102ea565b50505050905090810190601f16801561032f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50610348848461093e565b60408051606081018252602b81527f476976656e20636f646520697320616c726561647920706f696e74696e67207460208201527f6f2061206d6573736167650000000000000000000000000000000000000000009181019190915290156103f55760405160e560020a62461bcd028152600401808060200182810382528381815181526020019150805190602001908083836000838110156103025781810151838201526020016102ea565b5060ff8316600090815260208581526040909120835161041792850190610968565b505050600191820180549283018155600090815260209081902090830401805460ff808416601f9095166101000a948502940219169290921790915590565b60408051808201909152601e8152600080516020610a3d833981519152602082015281565b600061048682610939565b60408051808201909152601e8152600080516020610a3d833981519152602082015290156104f95760405160e560020a62461bcd028152600401808060200182810382528381815181526020019150805190602001908083836000838110156103025781810151838201526020016102ea565b50610504848461093e565b606060405190810160405280602681526020017f476976656e20636f646520646f6573206e6f7420706f696e7420746f2061206d81526020017f65737361676500000000000000000000000000000000000000000000000000008152509015156105b35760405160e560020a62461bcd028152600401808060200182810382528381815181526020019150805190602001908083836000838110156103025781810151838201526020016102ea565b5060ff831660009081526020858152604090912083516105d592850190610968565b50919392505050565b60008060006105ed858561093e565b606060405190810160405280602681526020017f476976656e20636f646520646f6573206e6f7420706f696e7420746f2061206d81526020017f657373616765000000000000000000000000000000000000000000000000000081525090151561069c5760405160e560020a62461bcd028152600401808060200182810382528381815181526020019150805190602001908083836000838110156103025781810151838201526020016102ea565b50600091505b8360ff16856001018360ff168154811015156106ba57fe5b60009182526020918290209181049091015460ff601f9092166101000a900416146106ea576001909101906106a2565b50805b60018501546000190160ff8216101561078257846001018160010160ff1681548110151561071757fe5b90600052602060002090602091828204019190069054906101000a900460ff16856001018260ff1681548110151561074b57fe5b90600052602060002090602091828204019190066101000a81548160ff021916908360ff16021790555080806001019150506106ed565b600185018054906107979060001983016109e6565b50604080516020818101808452600080845260ff891681529189905292902090516107c29290610968565b5092949350505050565b606060405190810160405280602681526020017f476976656e20636f646520646f6573206e6f7420706f696e7420746f2061206d81526020017f657373616765000000000000000000000000000000000000000000000000000081525081565b600061083782610939565b60408051808201909152601e8152600080516020610a3d833981519152602082015290156108aa5760405160e560020a62461bcd028152600401808060200182810382528381815181526020019150805190602001908083836000838110156103025781810151838201526020016102ea565b50600090505b6108ba838261093e565b156108c7576001016108b0565b6108d2838284610283565b5092915050565b606060405190810160405280602b81526020017f476976656e20636f646520697320616c726561647920706f696e74696e67207481526020017f6f2061206d65737361676500000000000000000000000000000000000000000081525081565b511590565b60ff1660009081526020919091526040812054600260001961010060018416150201909116041190565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109a957805160ff19168380011785556109d6565b828001600101855582156109d6579182015b828111156109d65782518255916020019190600101906109bb565b506109e2929150610a1f565b5090565b815481835581811115610a1a57601f016020900481601f01602090048360005260206000209182019101610a1a9190610a1f565b505050565b610a3991905b808211156109e25760008155600101610a25565b9056004d6573736167652063616e6e6f7420626520656d70747920737472696e670000a165627a7a72305820c89600f5a4f3ddc5e354cffd8975d4ec3b2924110071062bd475c31dd8cc1acd0029`

func DeployMessagesAndCodesContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesAndCodesContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessagesAndCodesContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, contract, nil
}
