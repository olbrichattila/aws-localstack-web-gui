import React, { useState } from "react";
import { deleteTopic, load, save, publish } from "../../api/sns";
import FilterBox from "../../components/filterBox";
import Spacer from "../../components/spacer";
import SaveBox from "../../components/savebox";
import Button from "../../components/button";
import InteractiveTable from "../../components/interactiveTable";

const TopicPage = ({ onManageSubs = () => null }) => {
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState('');
    const [newTopicModalOpen, setNewTopicModalOpen] = useState(false);
    const [messageArn, setMessageArn] = useState('');

    const onEvent = (e) => {

        if (e.name === 'Delete') {
            deleteTopic(e.i.TopicArn).then(() => load().then(data => setData(data)))
        }

        if (e.name === 'Send Message') {
            setMessageArn(e.i.TopicArn);
        }
    }

    useState(() => {
        load().then(data => setData(data));
    }, []);

    return (
        <div>
            <SaveBox
                isOpen={newTopicModalOpen}
                onClose={() => setNewTopicModalOpen(false)}
                title='New Topic:'
                onSubmit={name => save(name).then(() => load().then(r => {
                    setData(r);
                    setNewTopicModalOpen(false);
                }))}
            />

            <SaveBox
                isOpen={messageArn !== ''}
                onClose={() => setMessageArn('')}
                title='New message:'
                onSubmit={message => publish(messageArn, message).then(() => load().then(r => {
                    setData(r);
                    setMessageArn('');
                }))}
            />

            <Button label="Create new topic" margin={6} onClick={() => {
                setNewTopicModalOpen(true);
            }} />

            <Button label="Manage subscriptions" margin={6} onClick={() => onManageSubs()} />

            <FilterBox onSubmit={text => setFilter(text)} />
            <Spacer />
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: 'TopicArn',
                        asc: true
                    },
                    filterField: 'TopicArn',
                    columns: [
                        {
                            field: 'TopicArn',
                            title: 'Topic Arn'
                        },
                        {
                            field: 'SubscriptionsConfirmed',
                            title: 'Subscriptions Confirmed'
                        },
                        {
                            field: 'SubscriptionsDeleted',
                            title: 'Subscriptions Deleted'
                        },
                        {
                            field: 'SubscriptionsPending',
                            title: 'Subscriptions Pending'
                        },
                    ],
                    events: [
                        'Delete',
                        'Send Message'
                    ],
                }}
                data={data}
                filter={filter}
                onEvent={e => onEvent(e)}
            />
        </div>
    )
}

export default TopicPage;