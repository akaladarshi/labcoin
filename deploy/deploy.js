require("hardhat-deploy")
require("hardhat-deploy-ethers")
require("dotenv").config({ path: './shared/config/.env' })
const {ethers} = require("hardhat")

const PRIVATE_KEY = process.env.PRIVATE_KEY
const wallet = new ethers.Wallet(PRIVATE_KEY, ethers.provider)
module.exports = async ({ deployments }) => {
    const {deploy} = deployments;
    console.log("Wallet Ethereum Address:", wallet.address)

    // const labCoinContract = await ethers.getContractFactory("Labcoin");
    // const labCoin = await labCoinContract.connect(wallet)
    // const labcoin = await labCoin.deploy(wallet.address);
    // await labcoin.waitForDeployment()
    // console.log("Labcoin Contract Address:", labcoin.target)

    const researchContract = await ethers.getContractFactory("Research");
    const res = await researchContract.connect(wallet)
    const research = await res.deploy("0xb8c468321368fe8b367914d656b724e31CEe2708");
    await research.waitForDeployment()
    console.log("Research Contract Address:", research.target)
}