const get = async (path) => {
    return request('GET', path);
}

const post = async (path, payload) => {
    return request('POST', path, payload);
}

const del = async (path, payload) => {
    return request('DELETE', path, payload);
}

const request = async (type, path, payload = null) => {
    const options = {
        method: type,
        headers: {
            'Content-Type': 'application/json',
        }
    };

    if (payload) {
        options.body = JSON.stringify(payload);
    }

    try {
        const response = await fetch(`${process.env.REACT_APP_API_URL}${path}`, options);
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

export { get, post, del, request };
