import React, { useEffect, useState } from "react";
import FilterBox from "../../components/filterBox";
import Spacer from "../../components/spacer";
import InteractiveTable from "../../components/interactiveTable";
import { listTables, deleteTable, createTable } from "../../api/dynamoDb";
import Button from "../../components/button";
import DynamoDbCreateTable from "../../components/dynamoDbCreateTable";

const Tables = ({ onSelect = () => null }) => {
    const [filter, setFilter] = useState('');
    const [data, setData] = useState([]);
    const [error, setError] = useState('');
    const [createTableModelStatus, setCreateTableModelStatus] = useState(false);

    const onEvent = (e) => {
        if (e.name === 'Delete') {
            del(e.i.tableName);
        }

        if (e.name === 'clickable') {
            onSelect(e.i.tableName);
        }
    }

    const del = (tableName) => {
        deleteTable(tableName).then(() => load()).catch(e => setError(e));
    }

    const load = () => {
        listTables('tab', 100).then(talbeList => setData(talbeList.map(tableName => { return { tableName } })))
    }

    const save = (payload) => {
        createTable(payload).then(() => load()).catch(e => setError(e));
        setCreateTableModelStatus(false)
    }

    useEffect(() => {
        let timeoutId = -1;
        if (error !== '') {
            timeoutId = setTimeout(() => {
                setError('');
            }, 4000);
        }

        return () => {
            if (timeoutId !== -1) {
                clearTimeout(timeoutId);
            }
        };
    }, [error])

    useEffect(() => {
        load();
    }, []);

    return (
        <>
            <DynamoDbCreateTable
                isOpen={createTableModelStatus}
                onClose={() => setCreateTableModelStatus(false)}
                onSave={payload => save(payload)}

            />
            <Button margin={6} label="Create new table" onClick={() => setCreateTableModelStatus(true)} />
            {error !== '' && <div className="errorLine">{error.message ?? error}</div>}
            <FilterBox onSubmit={text => setFilter(text)} />
            <Spacer />
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: 'tableName',
                        asc: true,
                        clickable: true,
                    },
                    filterField: 'tableName',
                    columns: [
                        {
                            field: 'tableName',
                            title: 'Table Name',
                            clickable: true,
                        },
                    ],
                    events: [
                        'Delete',
                    ],
                }}
                data={data}
                filter={filter}
                onEvent={e => onEvent(e)}
            />
        </>
    );
}

export default Tables;
