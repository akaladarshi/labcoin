const express = require("express");
const { ethers } = require("hardhat");
const { ResearchFormData, StoreCID } = require('./models');
require("dotenv").config({ path: "../shared/config/.env" });

const app = express();
const port = process.env.EXPRESS_SERVER_PORT || 3000;
const contractAddress = process.env.RESEARCH_CONTRACT_ADDRESS;
const privateKey = process.env.PRIVATE_KEY;

app.use(express.json());

app.get("/", (req, res) => {
    res.send("API is running");
});

app.get("/allformdetails", async (req, res) => {
    try {
        console.log("Received request to fetch all form details");

        const contract = await ethers.getContractAt("Research", contractAddress);

        const formDetails = await contract.getAllFormsDetails();

        // Convert contract data to ResearchFormData instances
        const formDataArray = formDetails.map(data => ResearchFormData.fromContractData(data))

        // Serialize the response to JSON
        res.status(200).json(formDataArray.map(formData => formData.toJSON()));
    } catch (error) {
        console.error(error);
        res.status(500).json({ error: "An error occurred" });
    }
});

// POST endpoint to receive and parse JSON request
app.post('/store-cid', async (req, res) => {
    try {
        const {research_id, name, cid} = req.body;

        // Validate the request data
        if (typeof research_id !== 'number' || typeof name !== 'string' || typeof cid !== 'string') {
            return res.status(400).json({error: 'Invalid request data'});
        }

        // Create an instance of StoreCID
        const storeCID = new StoreCID(research_id, cid, name);

        // Log the parsed data (here you would send the transaction to the contract)
        console.log('Received data:', storeCID);

        //create a new wallet instance
        const wallet = new ethers.Wallet(privateKey, ethers.provider)
        const contract  = await ethers.getContractAt("Research", contractAddress)
        const researchContract = await contract.connect(wallet)

        const tx = await researchContract.storeCID(storeCID.researchID, storeCID.cid)
        const receipt = await tx.wait()

        console.log("Transaction hash:", tx.hash);

        // Respond with the parsed data
        res.status(200).json({ tx_hash: tx.hash });
    } catch (error) {
        console.error('Error processing request:', error);
        res.status(500).json({error: 'Internal Server Error'});
    }
});


app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});