task("query-cid", "Queries the cid based on the research id from the research contract")
    .addParam("researchid", "The research id")
    .setAction(async taskArgs => {
        const contractAddress = process.env.RESEARCH_CONTRACT_ADDRESS;
        const Contract = await ethers.getContractFactory("Research");
        const contract = Contract.attach(contractAddress);
        const cid = await contract.researchIDToCID(parseInt(taskArgs.researchid));
        // converts bytes to string
        console.log(cid);
    });
