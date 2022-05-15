# Decentralized File Archive(DFA)

## _A blockchain (Ethereum) based file sharing system_

DFA is an Ethereum blockchain based file sharing system, which makes file sharing decentralized and immutable. The permission of accessing a file is recorded on Ethereum, and the original file is stored on IPFS.

The file can be traded on this system. A file can be traded by ForForToken, which is maintained by a smart contract. Users can send and withdraw ETH through this contract.

## Prerequisite
DFA requires IPFS and [Ganache](https://goethereumbook.org/en/client-setup/) environment.

After the installation, set up the environment by:
```sh
ipfs daemon
ganache-cli -m "much repair shock carbon improve miss forget sock include bullet interest solution"
```
You may need 2 additional terminals to run these commonds seprately.
Then, you can trun on DFA by
```sh
go run server.go
```