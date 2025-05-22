import { get, post, del } from "./request";

const load = async () => {
    return get("/api/sqs/attributes");
};

const save = async (queueName) => {
    return post("/api/sqs", {
        name: queueName,
        delaySeconds: 5,
        maximumMessageSize: 4096,
    });
};

const saveFifo = async (queueName) => {
    return post("/api/sqs/fifo", {
        messageGroupId: "1234",
        messageDeduplicationId: "5678",
        name: queueName,
        maximumMessageSize: 4096,
    });
};

const delQueue = async (queueUrl) => {
    del("/api/sqs", {
        queueUrl: queueUrl,
    });
};

const purge = async (queueUrl) => {
    return del("/api/sqs/purge", {
        queueUrl: queueUrl,
    });
};

const refresh = async (queueUrl) => {
    return post("/api/sqs/attributes", {
        queueUrl: queueUrl,
    });
};

const sendMessage = async (queueUrl, delaySeconds, messageBody) => {
    return post("/api/sqs/message/send", {
        queueUrl,
        delaySeconds,
        messageBody,
    });
};

const sendFIFOMessage = async (queueUrl, messageGroupId, messageDeduplicationId,  messageBody) => {
    return post("/api/sqs/message/send/fifo", {
        messageGroupId: messageGroupId,
        messageDeduplicationId: messageDeduplicationId,
        queueUrl,
        messageBody,
    });
};

const receiveMessage = async (queueUrl, maxNumberOfMessages) => {
    return post("/api/sqs/message/receive", {
        queueUrl,
        maxNumberOfMessages,
    });
};

export {
    save,
    saveFifo,
    load,
    delQueue,
    purge,
    refresh,
    sendMessage,
    sendFIFOMessage,
    receiveMessage,
};
