import React, { useState } from "react";
import Tables from "./tables";
import Content from "./content";


const DynamoDBPage = () => {
    const [tableName, setTableName] = useState('');
    
    return (
        <>
        {tableName === '' && <Tables  onSelect={tableName => setTableName(tableName)} />}
        {tableName !== '' && <Content tableName={tableName} onBack={() => setTableName('')} />}
        </>
    );
}

export default DynamoDBPage;
