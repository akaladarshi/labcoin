require("dotenv").config({ path: './shared/config/.env' })
const {ethers} = require("hardhat")

const PRIVATE_KEY = process.env.PRIVATE_KEY
const wallet = new ethers.Wallet(PRIVATE_KEY, ethers.provider)

async function main() {
    const researchContract = await ethers.getContractFactory("Research");
    const research = await researchContract.attach("0x1baFC45eB9a5552166bd4E986f776fb9F6Bb39f6")

    const labCoinContract = await ethers.getContractFactory("Labcoin");
    const labCoin = await labCoinContract.attach("0xb8c468321368fe8b367914d656b724e31CEe2708");
    const labcoin = await labCoin.connect(wallet)

    // researchers will allow the contract to distribute the labcoins on their behalf.
    const tx = await labcoin.approve(research.target, ethers.parseEther("100"));
    await tx.wait()

    console.log("Transaction: ", tx.hash);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });