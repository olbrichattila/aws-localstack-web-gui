import React, { useState } from "react";
import TopicPage from "./topic";
import SubscriptionPage from "./subscription";

const SnsPage = () => {
    const [isSubPage, setIsSubPage] = useState(false);

    return (
        <>
            {!isSubPage && <TopicPage onManageSubs={() => setIsSubPage(true)} />}
            {isSubPage && <SubscriptionPage onBack={() => setIsSubPage(false)} />}
        </>
    )
}

export default SnsPage;
