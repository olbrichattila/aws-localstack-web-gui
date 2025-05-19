import React from "react";
import { useNavigate, useParams } from "react-router-dom";
import Button from "../../components/button";

const DynamoDbContent = () => {
    const navigate = useNavigate();
    const { tableName } = useParams();

    return (
        <>
            <Button margin={6} label="Back to table list" onClick={() => navigate('/aws/dynamodb')} />
            <h3>Table name: {tableName}</h3>
            <h2>Dynamo DB content page is Under development</h2>
        </>
    );
}

export default DynamoDbContent;
