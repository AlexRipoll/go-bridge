abigen:
	abigen --abi truffle/build/NFCustodialVault.abi --pkg contract --type CustosialVault --out blockchain/contract/custodial_vault.go

abigencompile:
	cd truffle/contracts && \
	solc --abi --include-path ../node_modules/ --base-path . custodialVault.sol -o ../build && \
	cd ../.. && sleep 3 && \
	abigen --abi truffle/build/NFCustodialVault.abi --pkg contract --type CustosialVault --out blockchain/contract/custodial_vault.go