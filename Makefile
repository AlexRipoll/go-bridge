abigen:
	abigen --abi truffle/build/NFTCustodialVault.abi --bin truffle/build/NFTCustodialVault.bin --pkg contract --type CustosialVault --out blockchain/contract/custodial_vault.go && \
	abigen --abi truffle/build/NFT.abi --bin truffle/build/NFT.bin --pkg contract --type NFT --out blockchain/contract/nft.go

compile:
	solc --overwrite --abi --bin --include-path ./truffle/node_modules/ --base-path . ./truffle/contracts/custodialVault.sol ./truffle/contracts/NFT.sol -o ./truffle/build

abigencompile:
	cd truffle/contracts && \
	solc --overwrite --abi --bin --include-path ../node_modules/ --base-path . custodialVault.sol NFT.sol -o ../build && \
	cd ../.. && sleep 3 && \
	abigen --bin truffle/build/NFTCustodialVault.bin --abi truffle/build/NFTCustodialVault.abi --pkg contract --type CustosialVault --out blockchain/contract/custodial_vault.go && \
	abigen --bin truffle/build/NFT.bin --abi truffle/build/NFT.abi --pkg contract --type NFT --out blockchain/contract/nft.go

mocks:
	mockgen -source=./blockchain/core/evm/contract.go -destination=./blockchain/core/evm/contract_mock.go -package=evm -aux_files=github.com/AlexRipoll/go-bridge/blockchain/core/evm=./blockchain/core/evm/geth.go