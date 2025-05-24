import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { useAppContext } from "../../AppContext";
import Button from "../../components/button";
import SaveBox from "../../components/savebox";
import FilterBox from "../../components/filterBox";
import InteractiveTable from "../../components/interactiveTable";
import Spacer from "../../components/spacer";

const SubscriptionPage = () => {
    const { get, post, del } = useAppContext();
    const { topicArn } = useParams();
    const navigate = useNavigate();
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState("");
    const [newSubModelVisible, setNewSubModelVisible] = useState(false);
    const [error, setError] = useState("");

    // APIS
    const sendHttpMessage = (topicArn, url) => {
        return post(`/api/sns/sub/${encodeURIComponent(topicArn)}`, { url });
    };
    
    const subList = async (topicArn) => {
        return get(`/api/sns/sub/${encodeURIComponent(topicArn)}`);
    };

    const deleteSub = async (subArn) => {
        return del(`/api/sns/sub/${encodeURIComponent(subArn)}`);
    };


    const onEvent = (e) => {
        if (e.name === "Delete") {
            deleteSub(e.i.SubscriptionArn).then(() =>
                subList(topicArn).then((data) => setFilteredData(data))
            );
        }
    };

    const setFilteredData = (data) => {
        setData(data.filter((item) => item.TopicArn === topicArn));
    };

    useEffect(() => {
        subList(topicArn)
            .then((data) => setFilteredData(data))
            .catch((err) => setError(err.message ?? "Error fetching data"));
    }, [topicArn]);

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
                isOpen={newSubModelVisible}
                onClose={() => setNewSubModelVisible(false)}
                title="New Subscription:"
                onSubmit={(message) =>
                    sendHttpMessage(topicArn, message)
                        .then(() =>
                            subList(topicArn).then((r) => {
                                setFilteredData(r);
                                setNewSubModelVisible(false);
                            })
                        )
                        .catch((error) => {
                            setError(error.message ?? "Cannot fetch data");
                            setNewSubModelVisible(false);
                        })
                }
            />

            <Button
                label="Back to topics"
                margin={6}
                onClick={() => navigate("/aws/sns")}
            />
            <Button
                label="Create new HTTP subscription"
                margin={6}
                onClick={() => {
                    setNewSubModelVisible(true);
                }}
            />

            <h3>Topic Arn: {topicArn}</h3>
            {error !== "" && <div className="errorLine">{error}</div>}

            <FilterBox onSubmit={(text) => setFilter(text)} />
            <Spacer />
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: "Endpoint",
                        asc: true,
                    },
                    filterField: "Endpoint",
                    columns: [
                        {
                            field: "Endpoint",
                            title: "Endpoint",
                            clickable: true,
                        },
                        {
                            field: "SubscriptionArn",
                            title: "Subscription Arn",
                            clickable: true,
                        },
                        {
                            field: "Owner",
                            title: "Owner",
                        },
                    ],
                    events: ["Delete"],
                }}
                data={data}
                filter={filter}
                onEvent={(e) => onEvent(e)}
            />
        </div>
    );
};

export default SubscriptionPage;
