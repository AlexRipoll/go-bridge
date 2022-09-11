abigen:
	abigen --abi truffle/build/NFCustodialVault.abi --pkg contract --type CustosialVault --out blockchain/contract/custodial_vault.go

abigencompile:
	cd truffle/contracts && \
	solc --overwrite --abi --bin --include-path ../node_modules/ --base-path . custodialVault.sol -o ../build && \
	cd ../.. && sleep 3 && \
	abigen --bin truffle/build/NFCustodialVault.bin --abi truffle/build/NFCustodialVault.abi --pkg contract --type CustosialVault --out blockchain/contract/custodial_vault.go

mocks:
	mockgen -source=./blockchain/core/evm/contract.go -destination=./blockchain/core/evm/contract_mock.go -package=evm