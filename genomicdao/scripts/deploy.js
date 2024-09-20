const {ethers} = require("hardhat");

const SMGeneNFT = "GeneNFT"
const SMPostCovidStrokePrevention = "PostCovidStrokePrevention"
const SMController = "Controller"

async function main() {
    // 1. Deploy the NFT contract
    const nft = await ethers.deployContract(SMGeneNFT)
    console.log(`${SMGeneNFT} deployed to: ${nft.target}`);

    // 2. Deploy the Token contract
    const pcspToken = await ethers.deployContract(SMPostCovidStrokePrevention)
    console.log(`${SMPostCovidStrokePrevention} deployed to: ${pcspToken.target}`);

    // 3. Deploy the Controller contract
    const controller = await ethers.deployContract("Controller", [nft.target, pcspToken.target]);
    await nft.transferOwnership(controller.target)
    await pcspToken.transferOwnership(controller.target)
    console.log(`${SMController} deployed to: ${controller.target}`);
}

main().then(() => {
    console.log("\nFinished!!!");
    process.exit(0);
}).catch((error) => {
    console.error(error);
    process.exit(1);
});
