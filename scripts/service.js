const express = require("express");
const { ethers, upgrades, network } = require("hardhat");
const {fromContractData} = require("./models");
require("dotenv").config();


// Function to handle BigInt serialization
const bigIntReplacer = (key, value) => {
    return typeof value === "bigint" ? value.toString() : value;
};

// Function to convert BigInt values to strings within the form details array
const convertBigIntToString = (formDetails) => {
    return formDetails.map((form) => {
        return form.map((value) => (typeof value === "bigint" ? value.toString() : value));
    });
};


const app = express();
const port = process.env.EXPRESS_SERVER_PORT || 3000;

app.use(express.json());

app.get("/", (req, res) => {
    res.send("API is running");
});

app.get("/allformdetails", async (req, res) => {
    try {
        const contractAddress = process.env.RESEARCH_CONTRACT_ADDRESS;
        const contract = await ethers.getContractAt("Research", contractAddress);

        const formDetails = await contract.getAllFormsDetails();

        // Convert contract data to ResearchFormData instances
        const formDataArray = formDetails.map(data => fromContractData(data))

        // Serialize the response to JSON
        res.status(200).json(formDataArray.map(formData => formData.toJSON()));
    } catch (error) {
        console.error(error);
        res.status(500).json({ error: "An error occurred" });
    }
});

app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});