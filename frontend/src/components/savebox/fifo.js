import { useState } from "react";
import Button from "../button";
import Modal from "../modal";
import "./index.scss";

const initialSaveFIFOState = {
    messageGroupId: "",
    messageDeduplicationId: "",
    message: "",
    messageAttributes: "",
};

const SaveFIFOBox = ({
    isOpen = false,
    title = "Name",
    onSubmit = () => null,
    onClose = () => null,
}) => {
    const [fifoState, setFifoState] = useState(initialSaveFIFOState);
    const [errors, setErrors] = useState([]);

    const validateAndSend = () => {
        const vErrors = [];
        if (fifoState.messageGroupId === "") {
            vErrors.push("Message Group ID is required");
        }

        if (fifoState.messageDeduplicationId === "") {
            vErrors.push("Message Deduplication ID is required");
        }

        if (fifoState.message === "") {
            vErrors.push("Message is required");
        }
        
        let attrs = fifoState.messageAttributes != "" ? fifoState.messageAttributes :  null;
        
        let dispatch = { ...fifoState, messageAttributes: attrs };
        if (attrs) {
            try {
                const parsed = JSON.parse(attrs);
                dispatch.messageAttributes = parsed;
            } catch (error) {
                dispatch.messageAttributes = null;
                vErrors.push(
                    "If message attributes are provided, it must be a valid JSON"
                );
            }
        }

        if (vErrors.length > 0) {
            setErrors(vErrors);
            return;
        }

        onSubmit(dispatch);
        setFifoState(initialSaveFIFOState);
        setErrors([]);
    };

    return (
        <Modal isOpen={isOpen} onClose={() => onClose()}>
            <div className="saveBox vertical">
                <label>
                    MessageGroup ID
                    <input
                        type="text"
                        value={fifoState.messageGroupId}
                        onChange={(e) =>
                            setFifoState({
                                ...fifoState,
                                messageGroupId: e.target.value,
                            })
                        }
                    />
                </label>
                <label>
                    Message Deduplication Id
                    <input
                        type="text"
                        value={fifoState.messageDeduplicationId}
                        onChange={(e) =>
                            setFifoState({
                                ...fifoState,
                                messageDeduplicationId: e.target.value,
                            })
                        }
                    />
                </label>
                <label>
                    {title}
                    <textarea
                        type="text"
                        value={fifoState.message}
                        onChange={(e) =>
                            setFifoState({
                                ...fifoState,
                                message: e.target.value,
                            })
                        }
                    />
                </label>
                <label>
                    Message Attributes (Optional)
                    <textarea
                        type="text"
                        value={fifoState.messageAttributes}
                        onChange={(e) =>
                            setFifoState({
                                ...fifoState,
                                messageAttributes: e.target.value,
                            })
                        }
                    />
                </label>
                {errors.map((e) => (
                    <div className="errorLine">{e}</div>
                ))}
                <Button onClick={() => validateAndSend()} label="Save" />
            </div>
        </Modal>
    );
};

export default SaveFIFOBox;
