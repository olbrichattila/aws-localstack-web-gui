import React, { useState } from "react";
import TopicPage from "./topic";
import SubscriptionPage from "./subscription";

const SnsPage = () => {
    const [topicArn, setTopicArn] = useState(false);

    return (
        <>
            {topicArn === '' && <TopicPage onManageSubs={topicArn => setTopicArn(topicArn)} />}
            {topicArn !== '' && <SubscriptionPage topicArn={topicArn} onBack={() => setTopicArn('')} />}
        </>
    )
}

export default SnsPage;
