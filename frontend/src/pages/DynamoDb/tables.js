import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAppContext } from "../../AppContext";
import FilterBox from "../../components/filterBox";
import Spacer from "../../components/spacer";
import InteractiveTable from "../../components/interactiveTable";
import Button from "../../components/button";
import DynamoDbCreateTable from "../../components/dynamoDbCreateTable";

const DynamoDbTables = () => {
    const { get, post, del } = useAppContext();
    const navigate = useNavigate();
    const [filter, setFilter] = useState("");
    const [data, setData] = useState([]);
    const [error, setError] = useState("");
    const [createTableModelStatus, setCreateTableModelStatus] = useState(false);
    const [tableListProps, setTableListProps] = useState({
        exclusiveStartTableName: "",
        limit: 2,
    });

    // APIS
    const listTables = async (prefix, limit) => {
        const suffix = prefix === "" ? "" : `/${encodeURIComponent(prefix)}`;
        return get(`/api/dynamodb-list/${limit}${suffix}`);
    };

    const deleteTable = async (tableName) => {
        return del(`/api/dynamodb/${encodeURIComponent(tableName)}`);
    };

    const createTable = async (payload) => {
        return post(`/api/dynamodb`, payload);
    };

    const onEvent = (e) => {
        if (e.name === "Delete") {
            delTable(e.i.tableName);
        }

        if (e.name === "clickable") {
            navigate(`/aws/dynamodb/${encodeURIComponent(e.i.tableName)}`);
        }
    };

    const delTable = (tableName) => {
        deleteTable(tableName)
            .then(() => load())
            .catch((e) => setError(e));
    };

    const load = () => {
        listTables(
            tableListProps.exclusiveStartTableName,
            parseInt(tableListProps.limit)
        )
            .then((talbeList) =>
                setData(
                    talbeList.map((tableName) => {
                        return { tableName };
                    })
                )
            )
            .catch((err) => setError(err.message ?? "Error fetching data"));
    };

    const save = (payload) => {
        createTable(payload)
            .then(() => load())
            .catch((e) => setError(e));
        setCreateTableModelStatus(false);
    };

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

    useEffect(() => {
        load();
    }, [tableListProps]);

    return (
        <>
            <DynamoDbCreateTable
                isOpen={createTableModelStatus}
                onClose={() => setCreateTableModelStatus(false)}
                onSave={(payload) => save(payload)}
            />
            <Button
                margin={6}
                label="Create new table"
                onClick={() => setCreateTableModelStatus(true)}
            />
            {error !== "" && (
                <div className="errorLine">{error.message ?? error}</div>
            )}
            <FilterBox onSubmit={(text) => setFilter(text)} />
            <Spacer />
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: "tableName",
                        asc: true,
                        clickable: true,
                    },
                    filterField: "tableName",
                    columns: [
                        {
                            field: "tableName",
                            title: "Table Name",
                            clickable: true,
                        },
                    ],
                    events: ["Delete"],
                }}
                data={data}
                filter={filter}
                onEvent={(e) => onEvent(e)}
            />
            {tableListProps.exclusiveStartTableName !== "" && (
                <Button
                    margin={6}
                    label="First page"
                    onClick={() =>
                        setTableListProps({
                            ...tableListProps,
                            exclusiveStartTableName: "",
                        })
                    }
                />
            )}
            {data.length > 0 && (
                <Button
                    margin={6}
                    label="Next page"
                    onClick={() =>
                        setTableListProps({
                            ...tableListProps,
                            exclusiveStartTableName:
                                data[data.length - 1].tableName,
                        })
                    }
                />
            )}
        </>
    );
};

export default DynamoDbTables;
