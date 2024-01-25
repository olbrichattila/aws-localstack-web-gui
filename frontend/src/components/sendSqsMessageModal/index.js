import React, { useEffect, useState } from 'react';
import Modal from '../modal';
import { sendMessage } from '../../api/sqs';
import { integerOnly } from '../../helpers';
import Button from '../button';
import './index.scss'

const initialInputParams = {
    delaySeconds: '5',
    messageBody: ''
}

const SendSqsMessageModal = ({
    isOpen,
    onClose,
    onSent,
    queueUrl,
}) => {
    const [inputParams, setInputParams] = useState(initialInputParams);
    const [errors, setErrors] = useState([]);

    const validate = () => {
        const errors = [];
        if (!queueUrl) errors.push('Queue not selected');
        if (!inputParams.delaySeconds) errors.push('Delay in seconds is required');
        if (!inputParams.messageBody) errors.push('Message body is required');

        return errors;
    }

    const send = () => {
        const validationErrors = validate();
        if (validationErrors.length > 0) {
            setErrors(validationErrors);
            return;
        }

        sendMessage(queueUrl, inputParams.delaySeconds, inputParams.messageBody)
            .then(() => {
                onSent(queueUrl);
                setInputParams(initialInputParams);
            });
    }

    const onCloseClick = () => {
        setInputParams(initialInputParams);
        onClose();
    }

    useEffect(() => {
        const timeoutId = setTimeout(() => {
            if (errors.length > 0) {
                setErrors([]);
            }
        }, 4000);
        return () => {
            clearTimeout(timeoutId);
        };
    }, [errors])

    return (
        <Modal isOpen={isOpen} onClose={() => onCloseClick()}>
            <div className='sendSqsMessageModalWrapper'>
                <label>
                    <span>Delay Seconds:</span>
                    <input
                        type='text'
                        value={inputParams.delaySeconds}
                        onChange={(e) => setInputParams({ ...inputParams, delaySeconds: integerOnly(e.target.value) })}
                    />
                </label>
                <label>
                    <span>Message Body:</span>
                    <textarea
                        type='text'
                        value={inputParams.messageBody}
                        onChange={(e) => setInputParams({ ...inputParams, messageBody: e.target.value })}
                    />
                </label>
                {errors.length > 0 && <ul>{errors.map(e => <li>{e}</li>)}</ul>}
                <Button label="Send" onClick={() => send()} />
            </div>
        </Modal>)
}

export default SendSqsMessageModal;
