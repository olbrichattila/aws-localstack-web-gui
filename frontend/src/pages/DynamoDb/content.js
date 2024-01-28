import React from "react";
import Button from "../../components/button";

const Content = ({
    tableName = '',
    onBack = () => null,
}) => {

    return (
        <>
            <h1>{tableName}</h1>
            <Button margin={6} label="Back to table list" onClick={() => onBack()} />
            <h2>Dynamo DB content page is Under development</h2>
        </>
    );
}

export default Content;
