import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { useAppContext } from "../../AppContext";
import InteractiveTable from "../../components/interactiveTable";
import Button from "../../components/button";

const SNSListPage = () => {
    const { get } = useAppContext();
    const { portNum } = useParams();
    const [data, setData] = useState([]);
    const [error, setError] = useState("");

    //API
    const getCapturedMessages = async (port) => {
        return get(`/api/sns/listener/${port}/get`);
    };

    const purgeMessages = (portNum) => {
        get(`/api/sns/listener/${portNum}/purge`)
            .then(() => {
                getCapturedMessages(portNum)
                    .then((data) => setData(data))
                    .catch((err) =>
                        setError(err.message ?? "Error fetching data")
                    );
            })
            .catch((err) => setError(err.message ?? "Error fetching data"));
    };

    useEffect(() => {
        getCapturedMessages(portNum)
            .then((data) => setData(data))
            .catch((err) => setError(err.message ?? "Error fetching data"));
    }, []);

    return (
        <div>
            <Button
                label="Refresh"
                margin={6}
                onClick={() => {
                    getCapturedMessages(portNum)
                        .then((data) => setData(data))
                        .catch((err) =>
                            setError(err.message ?? "Error fetching data")
                        );
                }}
            />
            <Button label="Purge" margin={6} onClick={() => purgeMessages(portNum)} />
            {error !== "" && <div className="errorLine">{error}</div>}

            {data && data.requests && (
                <InteractiveTable
                    structInfo={{
                        initialSort: {
                            field: "SubscribeURL",
                            asc: true,
                        },
                        filterField: "SubscribeURL",
                        columns: [
                            {
                                field: "Type",
                                title: "Type",
                                clickable: false,
                            },
                            {
                                field: "SubscribeURL",
                                title: "SubscribeURL",
                                clickable: false,
                            },
                            {
                                field: "Token",
                                title: "Token",
                                clickable: false,
                            },
                            {
                                field: "TopicArn",
                                title: "TopicArn",
                                clickable: false,
                            },
                            {
                                field: "Message",
                                title: "Message",
                                clickable: false,
                            },
                            {
                                field: "MessageAttributes",
                                title: "MessageAttributes",
                                clickable: false,
                            },
                            {
                                field: "MessageID",
                                title: "MessageID",
                                clickable: false,
                            },
                            {
                                field: "Timestamp",
                                title: "Timestamp",
                                clickable: false,
                            },
                            {
                                field: "SignatureVersion",
                                title: "SignatureVersion",
                                clickable: false,
                            },
                            {
                                field: "Signature",
                                title: "Signature",
                                clickable: false,
                            },
                            {
                                field: "SigningCertURL",
                                title: "SigningCertURL",
                                clickable: false,
                            },
                        ],
                    }}
                    data={data.requests}
                    filter={""}
                    //onEvent={(e) => console.log(e)}
                />
            )}
        </div>
    );
};

export default SNSListPage;
