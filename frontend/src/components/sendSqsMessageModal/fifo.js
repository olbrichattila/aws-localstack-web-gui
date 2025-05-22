import React, { useEffect, useState } from "react";
import Modal from "../modal";
import { sendFIFOMessage } from "../../api/sqs";
import { integerOnly } from "../../helpers";
import Button from "../button";
import "./index.scss";

const initialInputParams = {
    messageGroupId: "",
    messageDeduplicationId: "",
    messageBody: "",
};

const SendFIFOSqsMessageModal = ({ isOpen, onClose, onSent, queueUrl }) => {
    const [inputParams, setInputParams] = useState(initialInputParams);
    const [errors, setErrors] = useState([]);

    const validate = () => {
        const errors = [];
        if (!queueUrl) errors.push("Queue not selected");
        if (!inputParams.messageGroupId)
            errors.push(
                "MessageGroup ID is required"
            );
        if (!inputParams.messageDeduplicationId)
            errors.push(
                "Message Deduplication ID is required"
            );

        if (!inputParams.messageBody) errors.push("Message body is required");

        return errors;
    };

    const send = () => {
        const validationErrors = validate();
        if (validationErrors.length > 0) {
            setErrors(validationErrors);
            return;
        }

        sendFIFOMessage(
            queueUrl,
            inputParams.messageGroupId,
            initialInputParams.messageDeduplicationId,
            inputParams.messageBody
        ).then(() => {
            onSent(queueUrl);
            setInputParams(initialInputParams);
        });
    };

    const onCloseClick = () => {
        setInputParams(initialInputParams);
        onClose();
    };

    useEffect(() => {
        const timeoutId = setTimeout(() => {
            if (errors.length > 0) {
                setErrors([]);
            }
        }, 6000);
        return () => {
            clearTimeout(timeoutId);
        };
    }, [errors]);

    return (
        <Modal isOpen={isOpen} onClose={() => onCloseClick()}>
            <div className="sendSqsMessageModalWrapper">
                <label>
                    <span>MessageGroup ID:</span>
                    <input
                        type="text"
                        value={inputParams.messageGroupId}
                        onChange={(e) =>
                            setInputParams({
                                ...inputParams,
                                messageGroupId: e.target.value,
                            })
                        }
                    />
                </label>
                <label>
                    <span>Message Deduplication ID:</span>
                    <input
                        type="text"
                        value={inputParams.messageDeduplicationId}
                        onChange={(e) =>
                            setInputParams({
                                ...inputParams,
                                messageDeduplicationId: e.target.value,
                            })
                        }
                    />
                </label>
                <label>
                    <span>Message Body:</span>
                    <textarea
                        type="text"
                        value={inputParams.messageBody}
                        onChange={(e) =>
                            setInputParams({
                                ...inputParams,
                                messageBody: e.target.value,
                            })
                        }
                    />
                </label>
                {errors.length > 0 && (
                    <ul>
                        {errors.map((e) => (
                            <li>{e}</li>
                        ))}
                    </ul>
                )}
                <Button label="Send" onClick={() => send()} />
            </div>
        </Modal>
    );
};

export default SendFIFOSqsMessageModal;
