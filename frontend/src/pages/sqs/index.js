import React, { useState, useEffect } from "react";
import { load, save, saveFifo, delQueue, purge, refresh } from "../../api/sqs";
import SaveBox from "../../components/savebox";
import SendSqsMessageModal from "../../components/sendSqsMessageModal";
import FilterBox from "../../components/filterBox";
import Button from "../../components/button";
import Spacer from "../../components/spacer";
import ReadSqsQueue from "../../components/readSqsQueue";
import "./index.scss";
import InteractiveTable from "../../components/interactiveTable";

const SqsPage = () => {
    const [data, setData] = useState([]);
    const [watch, setWatch] = useState("");
    const [sendQueue, setSendQueue] = useState("");
    const [filter, setFilter] = useState("");
    const [newQueueModalOpen, setNewQueueModalOpen] = useState(false);
    const [newFIFOQueueModalOpen, setNewFIFOQueueModalOpen] = useState(false);
    const [sqsReadUrl, setSqsReadUrl] = useState("");
    const [error, setError] = useState("");

    const onEvent = (e) => {
        switch (e.name) {
            case "Delete":
                delQueue(e.i.url).then(() => load().then((r) => setData(r)));
                setWatch("");
                break;
            case "Purge":
                purge(e.i.url).then(() => load().then((r) => setData(r)));
                break;
            case "Refresh":
                refresh(e.i.url).then((r) => refreshQueue(r, e.i.url));
                break;
            case "Watch":
                setWatch(e.w ? e.i.url : "");
                break;
            case "Send Message":
                setSendQueue(e.i.url);
                break;
            case "Read Message":
                setSqsReadUrl(e.i.url);
                break;
            default:
                console.error("Invalid event type" + e.name);
        }
    };

    const refreshQueue = (attributes, url) => {
        const idx = data.findIndex((item) => item.url === url);
        if (idx !== -1) {
            setData([
                ...data.slice(0, idx),
                { ...data[idx], attributes },
                ...data.slice(idx + 1),
            ]);
        }
    };

    useEffect(() => {
        let timerId = -1;
        if (watch !== "") {
            timerId = setInterval(() => {
                refresh(watch).then((r) => refreshQueue(r, watch));
            }, 2000);
            refresh(watch).then((r) => refreshQueue(r, watch));
        }

        return () => {
            if (timerId !== -1) {
                clearInterval(timerId);
            }
        };
    }, [watch]);

    useEffect(() => {
        load()
            .then((r) => setData(r))
            .catch((err) => setError(err.message ?? "Error fetching data"));
    }, []);

    useEffect(() => {
        if (newQueueModalOpen) {
            setNewQueueModalOpen(false);
        }
        if (newFIFOQueueModalOpen) {
            setNewFIFOQueueModalOpen(false)
        }
        
    }, [data]);

    useEffect(() => {
        let timeoutId = -1;
        if (error !== "") {
            timeoutId = setTimeout(() => {
                setError("");
            }, 6000);
        }

        return () => {
            if (timeoutId !== -1) {
                clearTimeout(timeoutId);
            }
        };
    }, [error]);

    return (
        <>
            <SendSqsMessageModal
                queueUrl={sendQueue}
                isOpen={sendQueue !== ""}
                onClose={() => setSendQueue("")}
                onSent={(url) => {
                    refresh(url).then((r) => refreshQueue(r, url));
                    setSendQueue("");
                }}
            />

            <SaveBox
                isOpen={newQueueModalOpen}
                onClose={() => setNewQueueModalOpen(false)}
                title="Name:"
                onSubmit={(queueName) =>
                    save(queueName)
                        .then(() => load().then((r) => setData(r)))
                        .catch((error) => {
                            setError(error.message ?? "Cannot fetch data");
                            setNewQueueModalOpen(false);
                        })
                }
            />

            <SaveBox
                isOpen={newFIFOQueueModalOpen}
                onClose={() => setNewFIFOQueueModalOpen(false)}
                title="Name (FIFO):"
                onSubmit={(queueName) =>
                    saveFifo(queueName)
                        .then(() => load().then((r) => setData(r)))
                        .catch((error) => {
                            setError(error.message ?? "Cannot fetch data");
                            setNewFIFOQueueModalOpen(false);
                        })
                }
            />

            <ReadSqsQueue
                isOpen={sqsReadUrl !== ""}
                onClose={() => setSqsReadUrl("")}
                queueUrl={sqsReadUrl}
            />

            <Button
                label="Create new queue"
                margin={6}
                onClick={() => {
                    setNewQueueModalOpen(true);
                    if (watch !== "") {
                        setWatch("");
                    }
                }}
            />

            <Button
                label="Create new FIFO queue"
                margin={6}
                onClick={() => {
                    setNewFIFOQueueModalOpen(true);
                    if (watch !== "") {
                        setWatch("");
                    }
                }}
            />

            {error !== "" && <div className="errorLine">{error}</div>}

            <FilterBox onSubmit={(text) => setFilter(text)} />
            <Spacer />

            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: "TopicArn",
                        asc: true,
                    },
                    filterField: "url",
                    columns: [
                        {
                            field: "url",
                            title: "Queue Url",
                            clickable: false,
                        },
                        {
                            field: "attributes.ApproximateNumberOfMessages",
                            title: "Messages",
                            clickable: false,
                        },
                        {
                            field: "attributes.ApproximateNumberOfMessagesNotVisible",
                            title: "Messages Not Visible",
                            clickable: false,
                        },
                        {
                            field: "attributes.ApproximateNumberOfMessagesDelayed",
                            title: "Messages Delayed",
                            clickable: false,
                        },
                    ],
                    events: [
                        "Delete",
                        "Purge",
                        "Refresh",
                        "Watch",
                        "Send Message",
                        "Read Message",
                    ],
                    watchButton: "Watch",
                }}
                data={data}
                filter={filter}
                onEvent={(e) => onEvent(e)}
                watch={watch}
            />
        </>
    );
};

export default SqsPage;
