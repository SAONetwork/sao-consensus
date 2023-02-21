### Installation and Setup
#### Install Go 1.19.1

Currently, SAO Network uses Go 1.19.1 to compile the code.

Install [Go 1.19.1](https://go.dev/doc/install) by following instructions there.

Verify the installation by typing `go version` in your terminal.

```
$ go version
go version go1.19.1 darwin/amd64
```



#### Build Consensus Node 

```bash
$ git clone https://github.com/SaoNetwork/sao-consensus.git
$ git checkout testnet1
$ make 
$ which saod
```

#### Faucet

In order to get testnet tokens use [https://faucet.testnet.sao.network/](https://faucet.testnet.sao.network/)

#### Join Testnet

use saod to initialize your node 

```
$ saod init sao-testnet --chain-id=sao-test-1
```

add peer to config.toml

```
$ cd $HOME/.sao/config
$ vi config.toml
```

download genesis file

```
$ cd $HOME/.osmosisd/config
$ wget https://github.com/saonetwork/networks/raw/main/sao-test-1/genesis.json
```

run node

```
$ saod start
```



## License

Copyright © SAO Network, Inc. All rights reserved.

Licensed under the [Apache v2 License](LICENSE.md).