import React, { useState, useEffect } from 'react';
import './sqs.scss';
import { load, save, del, purge, refresh } from '../api/sqs'
import SaveBox from '../components/savebox';
import SendSqsMessageModal from '../components/sendSqsMessageModal';
import SqsTable from '../components/sqsTable';
import FilterBox from '../components/filterBox';
import Modal from '../components/modal';
import Button from '../components/button';
import Spacer from '../components/spacer';
import ReadSqsQueue from '../components/readSqsQueue';

const SqsPage = () => {
    const [data, setData] = useState([]);
    const [watch, setWatch] = useState(-1);
    const [sendQueue, setSendQueue] = useState(-1);
    const [filter, setFilter] = useState('');
    const [newQueueModalOpen, setNewQueueModalOpen] = useState(false);
    const [sqsReadId, setSqsReadId] = useState(-1);

    const refreshQueue = (attributes, idx) => {
        setData(
            [...data.slice(0, idx), { ...data[idx], attributes }, ...data.slice(idx + 1)]
        );
    }

    useEffect(() => {
        const updateTimer = () => {
            if (watch >= 0) {
                refresh(data[watch].url).then(r => refreshQueue(r, watch))
            }
        };

        const timerId = setInterval(updateTimer, 2000);

        return () => {
            clearInterval(timerId);

        };
    }, [watch]);


    useEffect(() => {
        load().then(r => setData(r));
    }, []);

    useEffect(() => {
        if (newQueueModalOpen) {
            setNewQueueModalOpen(false);
        }
    }, [data]);

    return (
        <div>
            <SendSqsMessageModal
                idx={sendQueue}
                queueUrl={sendQueue >= 0 ? data[sendQueue].url : ''}
                isOpen={sendQueue >= 0}
                onClose={() => setSendQueue(-1)}
                onSent={idx => {
                    refresh(data[idx].url).then(r => refreshQueue(r, sendQueue))
                    setSendQueue(-1)
                }}
            />
            <Modal
                isOpen={newQueueModalOpen}
                onClose={() => setNewQueueModalOpen()}
            >
                <SaveBox
                    title='New SQS queue name:'
                    onSubmit={queueName => save(queueName).then(() => load().then(r => setData(r)))}
                />
            </Modal>
            <ReadSqsQueue isOpen={sqsReadId >= 0} onClose={() => setSqsReadId(-1)} queueUrl={sqsReadId >= 0 ? data[sqsReadId].url : ''} />

            <Spacer />
            <Button label="Create new queue" margin={6} onClick={() => {
                setNewQueueModalOpen(true);
                if (watch >= 0) {
                    setWatch(-1);
                }
            }} />

            <FilterBox onSubmit={text => setFilter(text)} />
            <SqsTable
                data={data}
                filter={filter}
                watchIdx={watch}
                onDelete={idx => {
                    del(data[idx].url).then(() => load().then(r => setData(r)));
                    setWatch(-1);
                }
                }

                onPurge={idx => purge(data[idx].url).then(() => load().then(r => setData(r)))}
                onRefresh={idx => refresh(data[idx].url).then(r => refreshQueue(r, idx))}
                onWatchChange={idx => setWatch(idx)}
                onSendMessage={idx => setSendQueue(idx)}
                onReadMessage={idx => setSqsReadId(idx)}
            />
        </div>
    );
};

export default SqsPage;
