abigen:
	abigen --abi truffle/build/ERC721CustodialVault.abi --bin truffle/build/ERC721CustodialVault.bin --pkg contract --type CustosialVault --out ./contract/custodial_vault.go && \
	abigen --abi truffle/build/ERC721Token.abi --bin truffle/build/ERC721Token.bin --pkg contract --type erc721Token --out ./contract/erc721Token.go

compile:
	solc --overwrite --abi --bin --include-path ./truffle/node_modules/ --base-path . ./truffle/contracts/ERC721CustodialVault.sol ./truffle/contracts/ERC721Token.sol -o ./truffle/build

abigencompile:
	cd truffle/contracts && \
	solc --overwrite --abi --bin --include-path ../node_modules/ --base-path . custodialVault.sol erc721Token.sol -o ../build && \
	cd ../.. && sleep 3 && \
	abigen --bin truffle/build/CustodialVault.bin --abi truffle/build/CustodialVault.abi --pkg contract --type CustosialVault --out ./contract/custodial_vault.go && \
	abigen --bin truffle/build/erc721Token.bin --abi truffle/build/erc721Token.abi --pkg contract --type erc721Token --out ./contract/nft.go

mocks:
	mockgen -source=./core/evm/contract.go -destination=./core/evm/contract_mock.go -package=evm -aux_files=github.com/AlexRipoll/go-bridge/core/evm=./core/evm/geth.go