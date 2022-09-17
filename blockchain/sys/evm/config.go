package evm

type Config struct {
	//Url      string
	Pvk      string
	Address  string
	GasLimit uint64
	//Contracts Contract
	Networks []Network
}

type Network struct {
	Name string
	RpcUrl string
	Contract []Contract
}

type Contract struct {
	Type string
	Name string
	Address string
}

var c = Config{
	Pvk:       "",
	Address:   "",
	GasLimit:  0,
	Networks:  []Network{
		{
			Name:     "ethereum",
			RpcUrl:   "...",
			Contract: []Contract{
				{
					Type: "Custodian",
					Name:    "NFTCustodialVault",
					Address: "...",
				},
			},
		},
		{
			Name:     "polygon",
			RpcUrl:   "...",
			Contract: []Contract{
				{
					Type: "minter",
					Name:    "...",
					Address: "...",
				},
			},
		},
		{
			Name:     "hedera",
			RpcUrl:   "...",
			Contract: []Contract{
				{
					Type: "minter",
					Name:    "...",
					Address: "...",
				},
			},
		},
	},
}