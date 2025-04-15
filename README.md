<div align="center">
  <h1>Sacas Network</h1>
</div>

<div align="center">
  <a href="https://github.com/sacasnetwork/sacas/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/sacasnetwork/sacas.svg" />
  </a>
  <a href="https://github.com/sacasnetwork/sacas/blob/main/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/sacasnetwork/sacas.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/sacasnetwork/sacas">
    <img alt="GoDoc" src="https://godoc.org/github.com/sacasnetwork/sacas?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/sacasnetwork/sacas">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/sacasnetwork/sacas"/>
  </a>
</div>

## 📦 About Sacas Network
Sacas Network is a Layer 1 blockchain built using **Cosmos SDK** and **Tendermint**, designed for fast transactions and low gas fees. SACAS is fully **EVM-compatible**, enabling seamless development of smart contracts and decentralized applications (dApps).

## 📦 Key Features
- **Layer 1 Blockchain**
- **Proof of Stake (PoS) Consensus** [ Efficient and secure ]
- **EVM Compatibility** [ Supports Solidity-based smart contracts ]
- **Low Gas Fees** [ Cheaper transactions ]
- **High Transaction Speed** [ Near-instant processing ]
- **ERC-20 & ERC-721 Support** [ Deploy tokens and NFTs on SACAS ]

## 📦 Installation & Running a Node

### System Requirements
- OS: **Ubuntu 20.04 / 22.04**
- RAM: **8GB+**
- Storage: **200GB SSD+**
- Go: **v1.21+**
- Git: **Installed**

### 📦 Clone & Install Dependencies
```sh
# Clone the repository
git clone https://github.com/sacasnetwork/sacas.git
cd sacas

# Install dependencies
make install
```

### 📦 Initialize
```sh
sacasd init "Your Node Name" --chain-id sac_1317-1
sacasd keys add validator
sacasd add-genesis-account $(sacasd keys show validator -a) 1000000000asacas
sacasd gentx validator 1000000asacas --chain-id sac_1317-1
sacasd collect-gentxs
```

### 📦 Start the Node
```sh
sacasd start --chain-id sac_1317-1
```

## 📦 Network & RPC Configuration
- **Mainnet RPC:** [Coming Soon]
- **Testnet RPC:** http://localhost:26657
- **Explorer:** [Coming Soon]

## 📦 Developing Smart Contracts on SACAS
Sacas supports Solidity-based smart contracts. You can use Remix, Hardhat, or Truffle for deployment.

### 📦 Example: Deploying an ERC-20 Token
```solidity
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MyToken is ERC20 {
    constructor() ERC20("My Token", "MTK") {
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }
}
```

Deploy this contract using **Remix** or **Hardhat**.

## 📦 Contributing

1. Fork.
2. Create a feature or bugfix branch.
3. Submit a pull request.

## 📦 Community
- Twitter: [@sacasnetwork](https://twitter.com/sacasnetwork) | [@sacaswallet](https://twitter.com/sacaswallet)
- Discord: [Coming Soon](#)
- Website: [Coming Soon](#)
