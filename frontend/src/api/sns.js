import { get, post, del } from './request';

const load = async () => {
    return get('/api/sns/attributes');
}

const save = async (name) => {
    return post('/api/sns', {name})
}

const deleteTopic = async (name) => {
    return del('/api/sns', {name})
}

const sendHttpMessage = (topicArn, url) => {
    return post(`/api/sns/sub/${encodeURIComponent(topicArn)}`, {url});
    
}

const publish = (topicArn, message) => {
    return post(`/api/sns/sub/${encodeURIComponent(topicArn)}/publish`, {message});
}

export { load, save, deleteTopic, sendHttpMessage, publish };
