package evm

type Config struct {
	Url      string
	Pvk      string
	Address  string
	GasLimit uint64
	Contracts Contract
}

type Contract struct {
	Vault Deployment
	NFT []Deployment
}

type Deployment struct {
	Network string
	Address string
}