import React, { useState } from 'react';
import Modal from '../modal';
import Button from '../button';
import { receiveMessage } from '../../api/sqs';
import { integerOnly } from '../../helpers';
import './index.scss';

const ReadSqsQueue = ({ queueUrl, isOpen, onClose }) => {
    const [text, setText] = useState('5');
    const [messages, setMessages] = useState([]);

    const receive = () => {
        const nr = text === '' ? 1 : parseInt(text)

        receiveMessage(queueUrl, nr > 10 ? 10 : nr).then(m => setMessages(m));
    }

    return (
        <Modal
            isOpen={isOpen}
            onClose={() => {
                setMessages([]);
                onClose();
            }}
        >
            <div className='readSqsQueueModal'>
                <label>
                    Maximum number of messages (fetch max 10):
                    <input
                        type='text'
                        value={text}
                        onChange={(e) => setText(integerOnly(e.target.value))}
                    />

                    <Button label="Receive Messages" margin={6} onClick={() => receive()} />
                </label>
                <hr />
                <h3>{messages.length} messages received</h3>
                <div className='receiveMessageWrapper'>
                    {messages.map((m, idx) => <div key={idx}>{m.Body}</div>)}
                </div>
            </div>
        </Modal>
    );
}

export default ReadSqsQueue;
