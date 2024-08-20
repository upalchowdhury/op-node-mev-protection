require("@nomiclabs/hardhat-waffle");

module.exports = {
    solidity: "0.8.24",
    networks: {
        ganache: {
            url: "http://127.0.0.1:8545",
            accounts: [
                "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", // Replace with private keys from Ganache
                "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
            ]
        }
    }
};
