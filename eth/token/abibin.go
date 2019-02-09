package token

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_taxDestination\",\"type\":\"address\"}],\"name\":\"setTaxDestination\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taxDestination\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"defaultBalanceLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"limitOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taxPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_limit\",\"type\":\"uint256\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"setLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"},{\"name\":\"_v\",\"type\":\"uint8\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nonceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nameOf\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_taxDestination\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"TaxDestinationChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `608060405260016002556801158e460913d0000060035534801561002257600080fd5b5060405160208062001cb18339810180604052602081101561004357600080fd5b8101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36101288161012e640100000000026401000000009004565b50610270565b610145610219640100000000026401000000009004565b151561015057600080fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f09eee28d8d70bfad809ce8acadd46ce657b1fa64646b1e4b414e6bbb2eb2c8fa600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b611a3180620002806000396000f3fe608060405234801561001057600080fd5b5060043610610133576000357c0100000000000000000000000000000000000000000000000000000000900480637541f41c116100bf5780639dc29fac1161008e5780639dc29fac14610492578063e69babe8146104e0578063ed2a2d641461059a578063f2fde38b146105f2578063f5c573821461063657610133565b80637541f41c146103995780638da5cb5b146103b75780638e6e3e96146104015780638f32d59b1461047057610133565b806340c10f191161010657806340c10f191461027357806347cf5f04146102c1578063546a2ca4146102df57806370a0823114610337578063715018a61461038f57610133565b80631163c3eb1461013857806318160ddd1461017c5780631f6b6fa71461019a5780632c547b3d14610229575b600080fd5b61017a6004803603602081101561014e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106f3565b005b6101846107cf565b6040518082815260200191505060405180910390f35b610227600480360360c08110156101b057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803560ff1690602001909291905050506107d5565b005b61023161089a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102bf6004803603604081101561028957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108c0565b005b6102c9610a69565b6040518082815260200191505060405180910390f35b610321600480360360208110156102f557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a6f565b6040518082815260200191505060405180910390f35b6103796004803603602081101561034d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610abb565b6040518082815260200191505060405180910390f35b610397610b07565b005b6103a1610bd9565b6040518082815260200191505060405180910390f35b6103bf610bdf565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61046e600480360360a081101561041757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019092919080359060200190929190803560ff169060200190929190505050610c08565b005b610478610d30565b604051808215151515815260200191505060405180910390f35b6104de600480360360408110156104a857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d87565b005b610598600480360360a08110156104f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561053357600080fd5b82018360208201111561054557600080fd5b8035906020019184600183028401116401000000008311171561056757600080fd5b90919293919293908035906020019092919080359060200190929190803560ff169060200190929190505050610ef4565b005b6105dc600480360360208110156105b057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061105d565b6040518082815260200191505060405180910390f35b6106346004803603602081101561060857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110a6565b005b6106786004803603602081101561064c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110c5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106b857808201518184015260208101905061069d565b50505050905090810190601f1680156106e55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106fb610d30565b151561070657600080fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f09eee28d8d70bfad809ce8acadd46ce657b1fa64646b1e4b414e6bbb2eb2c8fa600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60055481565b61088786878787604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140182815260200193505050506040516020818303038152906040528585856111a9565b6108928686866114ef565b505050505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6108c8610d30565b15156108d357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561090f57600080fd5b6109248160055461160d90919063ffffffff16565b60058190555061097f81600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461160d90919063ffffffff16565b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60008383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b60035481565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201549050919050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050919050565b610b0f610d30565b1515610b1a57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60025481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610c77858686604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040528585856111a9565b83600660008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055507fef9c668177207fb68ca5e3894a1efacebb659762b27a737fde58ceebc4f30ad38585604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a15050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b610d8f610d30565b1515610d9a57600080fd5b610daf8160055461162e90919063ffffffff16565b600581905550610e0a81600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461162e90919063ffffffff16565b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef82600083604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b610f6c86878787604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018383808284378083019250505093505050506040516020818303038152906040528585856111a9565b8484600660008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000019190610fbd929190611960565b507f3b0a43ccc1ccd1c76ebb1a8d998fdfe1ded3766582dbbbcdda83889170bec53d868686604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a1505050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6110ae610d30565b15156110b957600080fd5b6110c281611650565b50565b6060600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561119d5780601f106111725761010080835404028352916020019161119d565b820191906000526020600020905b81548152906001019060200180831161118057829003601f168201915b50505050509050919050565b600060197f01000000000000000000000000000000000000000000000000000000000000000260007f01000000000000000000000000000000000000000000000000000000000000000230600160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548860405160200180867effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152600101857effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526001018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c0100000000000000000000000002815260140183815260200182805190602001908083835b6020831015156113485780518252602082019150602081019050602083039250611323565b6001836020036101000a03801982511681845116808217855250505050505090500195505050505050604051602081830303815290604052805190602001209050600060018284878760405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa1580156113e8573d6000803e3d6000fd5b5050506020604051035190508673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515611497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f73656e6465722d616464726573732d646f65732d6e6f742d6d6174636800000081525060200191505060405180910390fd5b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600101919050555050505050505050565b60006064600254830281151561150157fe5b0490506000818303905061153885600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168461174a565b61154385858361174a565b6000600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015414156115ac5760035461159a85610abb565b111515156115a757600080fd5b611606565b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201546115f885610abb565b1115151561160557600080fd5b5b5050505050565b600080828401905083811015151561162457600080fd5b8091505092915050565b600082821115151561163f57600080fd5b600082840390508091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561168c57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561178657600080fd5b6117db81600660008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461162e90919063ffffffff16565b600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555061187681600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461160d90919063ffffffff16565b600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106119a157803560ff19168380011785556119cf565b828001600101855582156119cf579182015b828111156119ce5782358255916020019190600101906119b3565b5b5090506119dc91906119e0565b5090565b611a0291905b808211156119fe5760008160009055506001016119e6565b5090565b9056fea165627a7a72305820537bafeb278dfbb73037b1910e8f4df6f07aa47f7189d147be4daf51a8ed67fc0029`
