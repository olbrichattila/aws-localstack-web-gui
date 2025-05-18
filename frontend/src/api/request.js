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
            let message = 'Whoops, something went wrong'

            if (json.error) {
                 message = json.error
            } else if (json.errors) {
                 message = json.errors
            }

            throw new Error(message);
        }

        const data = await response.json();
        return data;
    } catch (error) {
        throw error;
    }
}

export { get, post, del, request };
