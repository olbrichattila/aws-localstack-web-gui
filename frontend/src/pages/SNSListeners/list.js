import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import InteractiveTable from "../../components/interactiveTable";
import Button from "../../components/button";
import { getCapturedMessages } from "../../api/sns";

const SNSListPage = () => {
    const { portNum } = useParams();
    const [data, setData] = React.useState([]);
    const [error, setError] = useState("");

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
                    .catch((err) => setError(err.message ?? "Error fetching data"));
                }}
            />
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
                    filter={''}
                    onEvent={(e) => console.log(e)}
                />
            )}
        </div>
    );
};

export default SNSListPage;
