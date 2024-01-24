import { get, post } from './request';

const getSettings = async () => {
    return get('/api/settings');
}

const saveSettings = async (request) => {
    return post('/api/settings', request)
}

export { getSettings, saveSettings };
