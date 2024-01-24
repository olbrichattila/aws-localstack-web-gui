import { get, post, del } from './request';

const load = async () => {
    return get('/api/sqs/attributes');
}

const save = async (queueName) => {
    return post('/api/sqs', {
        name: queueName,
        delaySeconds: 5,
        maximumMessageSize: 4096
    });
}

const delQueue = async (queueUrl) => {
    del('/api/sqs', {
        "queueUrl": queueUrl
    });
}

const purge = async (queueUrl) => {
    return del('/api/sqs/purge', {
        "queueUrl": queueUrl
    });
}

const refresh = async (queueUrl) => {
    return post('/api/sqs/attributes', {
        "queueUrl": queueUrl
    });
}

const sendMessage = async (queueUrl, delaySeconds, messageBody) => {
    return post('/api/sqs/message/send', {
        queueUrl,
        delaySeconds,
        messageBody
    });
}

const receiveMessage = async (queueUrl, maxNumberOfMessages) => {
    return post('/api/sqs/message/receive', {
        queueUrl,
        maxNumberOfMessages,
    });
}

export { save, load, delQueue, purge, refresh, sendMessage, receiveMessage };
