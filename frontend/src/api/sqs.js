const headers = {
    'Content-Type': 'application/json',
};

const host  = process.env.REACT_APP_API_URL;

const load = async () => {
    try {
        const response = await fetch(`${host}/api/sqs/attributes`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching data:', error.message);
        throw error;
    }
}

const save = async (queueName) => {
    const options = {
        method: 'POST',
        headers,
        body: JSON.stringify(
            {
                name: queueName,
                delaySeconds: 5,
                maximumMessageSize: 4096
            }
        )
    };

    try {
        const response = await fetch(`${host}/api/sqs`, options);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching data:', error.message);
        throw error;
    }
}


const del = async (queueUrl) => {
    const options = {
        method: 'DELETE',
        headers,
        body: JSON.stringify(
            {
                "queueUrl": queueUrl
            }
        )
    };

    try {
        const response = await fetch(`${host}/api/sqs`, options);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching data:', error.message);
        throw error;
    }
}

const purge = async (queueUrl) => {
    const options = {
        method: 'DELETE',
        headers,
        body: JSON.stringify(
            {
                "queueUrl": queueUrl
            }
        )
    };

    try {
        const response = await fetch(`${host}/api/sqs/purge`, options);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching data:', error.message);
        throw error;
    }
}

const refresh = async (queueUrl) => {
    const options = {
        method: 'POST',
        headers,
        body: JSON.stringify(
            {
                "queueUrl": queueUrl
            }
        )
    };

    try {
        const response = await fetch(`${host}/api/sqs/attributes`, options);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching data:', error.message);
        throw error;
    }
}

export { save, load, del, purge, refresh };
