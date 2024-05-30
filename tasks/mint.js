task("mint", "Mints LabCoin")
    .addParam("contract", "The LabCoin address")
    .addParam("amount", "The amount to mint")
    .addParam("toaccount", "The account to mint to")
    .setAction(async (taskArgs) => {
        const labCoinContract = await ethers.getContractFactory("Labcoin");
        const labCoin = labCoinContract.attach(taskArgs.contract);

        const amount = ethers.parseEther(taskArgs.amount);
        // tx will be sent be hardhat network account
        // check hardhat.config.js for the account
        const tx = await labCoin.mint(taskArgs.toaccount, amount);
        tx.wait(5)
        console.log("Transaction: ", tx.hash);
    });