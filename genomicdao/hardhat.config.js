require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    solidity: "0.8.19",
    networks: {
        LIFENetworkLocal: {
            url: `http://127.0.0.1:9650/ext/bc/LIFENetwork/rpc`,
            accounts: [process.env.LIFENETWORK_LOCAL_PRIVATE_KEY],
            chainId: 9999
        },
        AvalancheFUJI: {
            url: `https://api.avax-test.network/ext/bc/C/rpc`,
            accounts: [process.env.LIFENETWORK_LOCAL_PRIVATE_KEY],
            chainId: 43113
        }
    }
};
