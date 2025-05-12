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
            const json = await response.json();
            const message = json.errors ? JSON.stringify(json) : 'Network response was not ok';

            throw new Error(message);
        }

        const data = await response.json();
        return data;
    } catch (error) {
        throw error;
    }
}

export { get, post, del, request };
