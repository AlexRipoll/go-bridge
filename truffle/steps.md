### FRONT
    npm start

### BACK
    go run cmd/main.go

    go run cmd/go-bridgecli/main.go --fn mint 0xe2d56Ed64C3Ea318710AEBC47230a3e8b016ACD6 2
    go run cmd/go-bridgecli/main.go --fn mint 0xe2d56Ed64C3Ea318710AEBC47230a3e8b016ACD6 4
    go run cmd/go-bridgecli/main.go --fn mint 0xe2d56Ed64C3Ea318710AEBC47230a3e8b016ACD6 8

    go run cmd/go-bridgecli/main.go --fn mint 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 1000
    go run cmd/go-bridgecli/main.go --fn mint 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 2000
    go run cmd/go-bridgecli/main.go --fn mint 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 3000
    go run cmd/go-bridgecli/main.go --fn mint 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 4000
    go run cmd/go-bridgecli/main.go --fn mint 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 5000
    go run cmd/go-bridgecli/main.go --fn mint 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 6000
    go run cmd/go-bridgecli/main.go --fn mint 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 7000

    go run cmd/go-bridgecli/main.go --fn transfer 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 0xAa573B03B6E4506C811Eb7Bf1EEa685d224087b4 1000
    go run cmd/go-bridgecli/main.go --fn transfer 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88 0xAa573B03B6E4506C811Eb7Bf1EEa685d224087b4 2000

    go run cmd/go-bridgecli/main.go --fn wallet 0xAa573B03B6E4506C811Eb7Bf1EEa685d224087b4
    go run cmd/go-bridgecli/main.go --fn wallet 0xc239f1E97Fa0B8cCB76b95a48c297C61584Ccf88


### TRUFFLE
    truffle console --network develop
    truffle console --network develop2

    truffle migrate --reset --network develop
    truffle migrate --reset --network develop2

### GANACHE
    ganache --chain.chainId 97 -i 97 -b 1 -m "mimic twin dismiss topic special ranch decline hen soap buzz analyst exact"

