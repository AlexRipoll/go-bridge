### FRONT
    npm start

### BACK
    go run cmd/main.go

    go run cmd/go-bridgecli/main.go --fn mint 0xCd1c0083BD35DDF4221bbC5Ee68fffBE3B8f0DE9 8
    go run cmd/go-bridgecli/main.go --fn wallet 0xCd1c0083BD35DDF4221bbC5Ee68fffBE3B8f0DE9

### TRUFFLE
    truffle migrate --reset --network develop
    truffle migrate --reset --network develop2

    truffle console --network develop
    truffle console --network develop2

### GANACHE
    ganache --chain.chainId 97 -i 97 -b 1 -m "mimic twin dismiss topic special ranch decline hen soap buzz analyst exact"

