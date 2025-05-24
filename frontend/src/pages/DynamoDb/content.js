import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import InteractiveTable from "../../components/interactiveTable";
import { useAppContext } from "../../AppContext";
import Button from "../../components/button";

const DynamoDbContent = () => {
    const { post } = useAppContext();
    const [data, setData] = useState([]);
    const [error, setError] = useState("");
    const navigate = useNavigate();
    const { tableName } = useParams();

    // API call
    const scanTable = async (payload) => {
        return post(`/api/scan_dynamodb`, payload);
    };

    useState(() => {
        const payload = {
            tableName: tableName,
            limit: 1,
            startKey: {},
        };
        scanTable(payload)
            .then((r) => {
                setData(
                    r.items.map((item) => {
                        return { item: JSON.stringify(item) };
                    })
                );
            })
            .catch((err) => setError(err.message ?? "Error fetching data"));
    }, []);

    return (
        <>
            <Button
                margin={6}
                label="Back to table list"
                onClick={() => navigate("/aws/dynamodb")}
            />
            <h3>Table name: {tableName}</h3>
            {error !== "" && (
                <div className="errorLine">{error.message ?? error}</div>
            )}
            <p>
                Note: This feature is under development, insert, delete and
                pagination coming soon,
                <br /> now just list the beginning of the table
            </p>
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: "item",
                        asc: true,
                        clickable: true,
                    },
                    filterField: "item",
                    columns: [
                        {
                            field: "item",
                            title: "Item",
                            clickable: true,
                        },
                    ],
                    // events: [
                    //     'Delete',
                    // ],
                }}
                data={data}
                // filter={filter}
                // onEvent={e => onEvent(e)}
            />
        </>
    );
};

export default DynamoDbContent;
