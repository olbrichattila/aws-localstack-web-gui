import React from "react";
import Button from "../../components/button";

const SubscriptionPage = ({onBack = () => null}) => {

    return (
        <div>
            <Button label="Back to topics" margin={6} onClick={() => onBack()} />
        </div>
    )
}

export default SubscriptionPage;
