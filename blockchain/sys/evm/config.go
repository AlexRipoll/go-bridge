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
