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
.addParam("title", "The title of the research")
.addParam("description", "The description of the research")
.addParam("timeduration", "The time duration of the research")
.addParam("datasetcount", "The data set count of the research")
.addParam("link", "The link of the research")
.addParam("datalink", "The data link of the research")
.setAction(async taskArgs => {
    const contractAddr = taskArgs.contract;
    const Contract = await ethers.getContractFactory("Research");
    const contract = Contract.attach(contractAddr);
    const researchData = {
        title: taskArgs.title,
        description: taskArgs.description,
        timeDuration: parseInt(taskArgs.timeduration),
        dataSetCount: parseInt(taskArgs.datasetcount),
        researcher: "",
        data: {
            link: taskArgs.link,
            dataLink: taskArgs.datalink
        }
    };
    const tx = await contract.registerResearch(
        researchData.title,
        researchData.description,
        researchData.data.link,
        researchData.data.dataLink,
        researchData.dataSetCount,
        researchData.timeDuration,
    );
    tx.wait(5)
    console.log("Transaction: ", tx.hash);
});