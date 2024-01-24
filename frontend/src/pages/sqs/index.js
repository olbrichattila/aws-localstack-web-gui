import React, { useState, useEffect } from 'react';
import { load, save, delQueue, purge, refresh } from '../../api/sqs'
import SaveBox from '../../components/savebox';
import SendSqsMessageModal from '../../components/sendSqsMessageModal';
import SqsTable from '../../components/sqsTable';
import FilterBox from '../../components/filterBox';
import Modal from '../../components/modal';
import Button from '../../components/button';
import Spacer from '../../components/spacer';
import ReadSqsQueue from '../../components/readSqsQueue';
import './index.scss';

const SqsPage = () => {
    const [data, setData] = useState([]);
    const [watch, setWatch] = useState('');
    const [sendQueue, setSendQueue] = useState('');
    const [filter, setFilter] = useState('');
    const [newQueueModalOpen, setNewQueueModalOpen] = useState(false);
    const [sqsReadUrl, setSqsReadUrl] = useState('');

    const refreshQueue = (attributes, url) => {
        const idx = data.findIndex(item => item.url === url);
        if (idx !== -1) {
            setData(
                [...data.slice(0, idx), { ...data[idx], attributes }, ...data.slice(idx + 1)]
            );
        }
    }

    useEffect(() => {
        const updateTimer = () => {
            if (watch !== '') {
                refresh(watch).then(r => refreshQueue(r, watch))
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
        <>
            <SendSqsMessageModal
                queueUrl={sendQueue}
                isOpen={sendQueue !== ''}
                onClose={() => setSendQueue('')}
                onSent={url => {
                    refresh(url).then(r => refreshQueue(r, url))
                    setSendQueue('')
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
            <ReadSqsQueue isOpen={sqsReadUrl !== ''} onClose={() => setSqsReadUrl('')} queueUrl={sqsReadUrl} />

            <Button label="Create new queue" margin={6} onClick={() => {
                setNewQueueModalOpen(true);
                if (watch !== '') {
                    setWatch('');
                }
            }} />

            <FilterBox onSubmit={text => setFilter(text)} />
            <Spacer />
            <SqsTable
                data={data}
                filter={filter}
                watchUrl={watch}
                onDelete={url => {
                    delQueue(url).then(() => load().then(r => setData(r)));
                    setWatch('');
                }
                }

                onPurge={url => purge(url).then(() => load().then(r => setData(r)))}
                onRefresh={url => refresh(url).then(r => refreshQueue(r, url))}
                onWatchChange={url => setWatch(url)}
                onSendMessage={url => setSendQueue(url)}
                onReadMessage={url => setSqsReadUrl(url)}
            />
        </>
    );
};

export default SqsPage;
