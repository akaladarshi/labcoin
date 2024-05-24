const { ethers,upgrades,network} = require("hardhat");
const fs = require("fs")

const contractAddress = process.env.RESEARCH_CONTRACT_ADDRESS;
// const wallet = new ethers.Wallet(network.config.accounts[0], ethers.provider)

async function main() {
    const contract = await ethers.getContractAt("Research", contractAddress);

    const researches = await contract.getAllResearches();
    console.log(`Total searches: ${researches}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
