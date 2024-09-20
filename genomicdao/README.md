# Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, and a script that deploys that contract.

Try running some of the following tasks:

```shell
npm install @openzeppelin/contracts@v4.9.6
```

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.js
```

## Tech debt

### @openzeppelin/contracts
Cannot upgrade `@openzeppelin/contracts` to `v5` because `@openzeppelin/contracts/utils/Counters.sol` is removed in that version.

Solution: https://github.com/OpenZeppelin/openzeppelin-contracts/issues/4233#issuecomment-1561791911
