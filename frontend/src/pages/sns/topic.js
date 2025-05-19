import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { deleteTopic, load, save, publish } from "../../api/sns";
import FilterBox from "../../components/filterBox";
import Spacer from "../../components/spacer";
import SaveBox from "../../components/savebox";
import Button from "../../components/button";
import InteractiveTable from "../../components/interactiveTable";

const TopicPage = () => {
    const navigate = useNavigate();
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState('');
    const [newTopicModalOpen, setNewTopicModalOpen] = useState(false);
    const [messageArn, setMessageArn] = useState('');
    const [error, setError] = useState('');

    const onEvent = (e) => {
        if (e.name === 'Delete') {
            deleteTopic(e.i.TopicArn).then(() => load().then(data => setData(data)))
        }

        if (e.name === 'Send Message') {
            setMessageArn(e.i.TopicArn);
        }

        if (e.name === 'clickable') {
            navigate(`/aws/sns/${encodeURIComponent(e.i.TopicArn)}`);
        }
    }

    useEffect(() => {
        load().then(data => setData(data)).catch(err => setError(err.message ?? 'Error fetching data'));
    }, []);

    useEffect(() => {
        let timeoutId = -1;
        if (error !== '') {
            timeoutId = setTimeout(() => {
                setError('');
            }, 6000);
        }

        return () => {
            if (timeoutId !== -1) {
                clearTimeout(timeoutId);
            }
        };
    }, [error])

    return (
        <div>
            <SaveBox
                isOpen={newTopicModalOpen}
                onClose={() => setNewTopicModalOpen(false)}
                title='New Topic:'
                onSubmit={name => save(name).then(() => load().then(r => {
                    setData(r);
                    setNewTopicModalOpen(false);
                })).catch(error => {
                    setError(error.message ?? 'Error fetching data');
                    setNewTopicModalOpen(false);
                })}
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

            {error !== '' && <div className="errorLine">{error}</div>}

            {/* <Button label="Manage subscriptions" margin={6} onClick={() => onManageSubs()} /> */}

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
                            title: 'Topic Arn',
                            clickable: true,
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