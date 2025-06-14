import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAppContext } from "../../AppContext";
import FilterBox from "../../components/filterBox";
import Spacer from "../../components/spacer";
import SaveBox from "../../components/savebox";
import SaveSNS from "../../components/savebox/sns";
import SaveFIFOBox from "../../components/savebox/fifo";
import Button from "../../components/button";
import InteractiveTable from "../../components/interactiveTable";

const TopicPage = () => {
    const navigate = useNavigate();
    const { get, post, del } = useAppContext();
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState("");
    const [newTopicModalOpen, setNewTopicModalOpen] = useState(false);
    const [messageArn, setMessageArn] = useState("");
    const [fifoMessageArn, setFIFOMessageArn] = useState("");
    const [error, setError] = useState("");

    // APIS
    const load = async () => {
        return get("/api/sns/attributes");
    };

    const save = async (name) => {
        return post("/api/sns", { name });
    };

    const deleteTopic = async (name) => {
        return del("/api/sns", { name });
    };

    const publish = (topicArn, message) => {
        return post(
            `/api/sns/sub/${encodeURIComponent(topicArn)}/publish`,
            message
        );
    };

    const publishFIFO = (topicArn, message) => {
        return post(
            `/api/sns/sub/${encodeURIComponent(topicArn)}/publish_fifo`,
            message
        );
    };
    // END APIS

    const onEvent = (e) => {
        if (e.name === "Delete") {
            deleteTopic(e.i.TopicArn).then(() =>
                load().then((data) => setData(data))
            );
        }

        if (e.name === "Send Message") {
            if (e.i.TopicArn.endsWith(".fifo")) {
                setFIFOMessageArn(e.i.TopicArn);
            } else {
                setMessageArn(e.i.TopicArn);
            }
        }

        if (e.name === "clickable") {
            navigate(`/aws/sns/${encodeURIComponent(e.i.TopicArn)}`);
        }
    };

    useEffect(() => {
        load()
            .then((data) => setData(data))
            .catch((err) => setError(err.message ?? "Error fetching data"));
    }, []);

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
        <div>
            <SaveBox
                isOpen={newTopicModalOpen}
                onClose={() => setNewTopicModalOpen(false)}
                title="New Topic:"
                onSubmit={(name) =>
                    save(name)
                        .then(() =>
                            load().then((r) => {
                                setData(r);
                                setNewTopicModalOpen(false);
                            })
                        )
                        .catch((error) => {
                            setError(error.message ?? "Error fetching data");
                            setNewTopicModalOpen(false);
                        })
                }
            />

            <SaveSNS
                isOpen={messageArn !== ""}
                onClose={() => setMessageArn("")}
                title="New message:"
                onSubmit={(message) => {
                    publish(messageArn, message)
                        .then(() =>
                            load().then((r) => {
                                setData(r);
                                setMessageArn("");
                            })
                        )
                        .catch((err) => {
                            const errorToDisplay = err.message ?? "Error fetching data"
                            setError("It looks like the message attributes are not following AWS attribute format: (" +errorToDisplay + ")");
                            setMessageArn("");
                        });
                }}
            />

            <SaveFIFOBox
                isOpen={fifoMessageArn !== ""}
                onClose={() => setFIFOMessageArn("")}
                title="New FIFO message:"
                onSubmit={(message) => {
                    publishFIFO(fifoMessageArn, message)
                        .then(() =>
                            load().then((r) => {
                                setData(r);
                                setFIFOMessageArn("");
                            })
                        )
                        .catch((err) => {
                            const errorToDisplay = err.message ?? "Error fetching data"
                            setError("It looks like the message attributes are not following AWS attribute format: (" +errorToDisplay + ")");
                            setFIFOMessageArn("");
                        });
                }}
            />

            <Button
                label="Create new topic"
                margin={6}
                onClick={() => {
                    setNewTopicModalOpen(true);
                }}
            />

            {error !== "" && <div className="errorLine">{error}</div>}

            {/* <Button label="Manage subscriptions" margin={6} onClick={() => onManageSubs()} /> */}

            <FilterBox onSubmit={(text) => setFilter(text)} />
            <Spacer />
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: "TopicArn",
                        asc: true,
                    },
                    filterField: "TopicArn",
                    columns: [
                        {
                            field: "TopicArn",
                            title: "Topic Arn",
                            clickable: true,
                        },
                        {
                            field: "SubscriptionsConfirmed",
                            title: "Subscriptions Confirmed",
                        },
                        {
                            field: "SubscriptionsDeleted",
                            title: "Subscriptions Deleted",
                        },
                        {
                            field: "SubscriptionsPending",
                            title: "Subscriptions Pending",
                        },
                    ],
                    events: ["Delete", "Send Message"],
                }}
                data={data}
                filter={filter}
                onEvent={(e) => onEvent(e)}
            />
        </div>
    );
};

export default TopicPage;
