const axios = require('axios');
const fs = require('fs');
require('dotenv').config();

const JOTFORM_API_KEY = process.env.JOTFORM_API_KEY;
const FORM_ID = 'YOUR_FORM_ID';

async function fetchFormData() {
    try {
        const response = await axios.get(`https://api.jotform.com/form/${FORM_ID}/submissions`, {
            params: {
                apiKey: JOTFORM_API_KEY
            }
        });

        const formData = response.data;

        // Save data to a local file
        fs.writeFileSync('formData.json', JSON.stringify(formData, null, 2));
        console.log('Data has been saved to formData.json');
    } catch (error) {
        console.error('Error fetching form data:', error);
    }
}

fetchFormData();