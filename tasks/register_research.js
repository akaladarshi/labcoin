// const Web3 = require('web3');
// const web3 = new Web3(Web3.givenProvider || 'http://localhost:8545');
//
// const contractAddress = 'YOUR_CONTRACT_ADDRESS';
// const contractABI = [ /* ABI from the compiled contract */ ];
//
// const researchRegistry = new web3.eth.Contract(contractABI, contractAddress);
//
// document.getElementById('researchForm').addEventListener('submit', async (e) => {
//     e.preventDefault();
//     const accounts = await web3.eth.requestAccounts();
//     const researchData = {
//         title: document.getElementById('title').value,
//         description: document.getElementById('description').value,
//         timeDuration: parseInt(document.getElementById('timeDuration').value), // TODO: convert into unix timestamp
//         dataSetCount: parseInt(document.getElementById('dataSetCount').value),
//         data: {
//             link: document.getElementById('formLink').value,
//             dataLink: document.getElementById('dataLink').value
//         }
//     };
//
//     await researchRegistry.methods.registerResearch(researchData).send({ from: accounts[0] });
// });

task("register-research", "Calls the registerResearch function of the contract")
    .addParam("contract", "The address of the contract that you want to interact with")
    .addParam("labcoincontract", "The address of the labcoin contract")
    .addParam("title", "The title of the research")
    .addParam("description", "The description of the research")
    .addParam("timeduration", "The time duration of the research")
    .addParam("maxdatasetcount", "The max data set count for the research")
    .addParam("formlink", "The link of the form research")
    .addParam("spreadsheetid", "The id of the research spreadsheet")
    .addParam("sheetid", "The id of sheet of the research spreadsheet")
    .addParam("amount", "The amount to distribute to the research participants")
    .setAction(async taskArgs => {
        const contractAddr = taskArgs.contract;
        const ResearchContract = await ethers.getContractFactory("Research");
        const researchContract = ResearchContract.attach(contractAddr);

        const labCoinContract = await ethers.getContractFactory("Labcoin");
        const labCoin = labCoinContract.attach(taskArgs.labcoincontract);

        // researchers will allow the contract to distribute the labcoins on their behalf.
        await labCoin.approve(researchContract.target, ethers.parseEther("100"));

        const amount= ethers.parseEther(taskArgs.amount);
        const tx = await researchContract.registerResearch(
            taskArgs.title,
            taskArgs.description,
            taskArgs.formlink,
            taskArgs.spreadsheetid,
            parseInt(taskArgs.sheetid),
            parseInt(taskArgs.maxdatasetcount),
            parseInt(taskArgs.timeduration),
            {value: amount },
        );
        tx.wait(5)
        console.log("Transaction: ", tx.hash);
    });