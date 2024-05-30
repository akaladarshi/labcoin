require("hardhat-deploy")
require("hardhat-deploy-ethers")
require("dotenv").config({ path: './shared/config/.env' })

const { networkConfig } = require("../helper-hardhat-config")

const PRIVATE_KEY = process.env.PRIVATE_KEY
const wallet = new ethers.Wallet(PRIVATE_KEY, ethers.provider)

module.exports = async ({ deployments }) => {
    const {deploy} = deployments;
    console.log("Wallet Ethereum Address:", wallet.address)
    const chainId = network.config.chainId
    const labCoinContract = await deploy("Labcoin", {
        from: wallet.address,
        args: [wallet.address],
        log: true,
    })

    const researchContract = await deploy("Research", {
        from: wallet.address,
        args: [labCoinContract.address],
        log: true,
    })

}