const headers = {
    'Content-Type': 'application/json',
};

const host  = process.env.REACT_APP_API_URL;

const loadBuckets = async () => {
    try {
        const response = await fetch(`${host}/api/s3/buckets`);
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

const newBucket = async (bucketName) => {
    const options = {
        method: 'POST',
        headers,
        body: JSON.stringify(
            {
                bucketName,
            }
        )
    };

    try {
        const response = await fetch(`${host}/api/s3/buckets`, options);
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

const delBucket = async (bucketName) => {
    const options = {
        method: 'DELETE',
        headers,
        body: JSON.stringify(
            {
                bucketName,
            }
        )
    };

    try {
        const response = await fetch(`${host}/api/s3/buckets`, options);
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

const listBucketContent = async (bucketName) => {
    try {
        const response = await fetch(`${host}/api/s3/list/${encodeURIComponent(bucketName)}`);
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

const delObject = async (bucketName, key) => {
    const options = {
        method: 'DELETE',
        headers,
        body: JSON.stringify(
            {
                bucketName,
                fileName: key,
            }
        )
    };

    try {
        const response = await fetch(`${host}/api/s3/buckets/delete/object`, options);
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

const upload = async (bucketName, fileName) => {
    const options = {
        method: 'POST',
        headers,
        body: JSON.stringify(
            {
                bucketName,
                fileName,
            }
        )
    };

    try {
        const response = await fetch(`${host}/api/s3/buckets/upload`, options);
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

export { loadBuckets, newBucket, delBucket, listBucketContent, delObject, upload }
