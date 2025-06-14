import { useState } from "react";
import Button from "../button";
import Modal from "../modal";
import "./index.scss";

const initialSaveSNSState = {
    message: "",
    messageAttributes: "",
};

const SaveSNS = ({
    isOpen = false,
    title = "Name",
    onSubmit = () => null,
    onClose = () => null,
}) => {
    const [snsState, setSnsState] = useState(initialSaveSNSState);
    const [errors, setErrors] = useState([]);

    const validateAndSend = () => {
        const vErrors = [];

        if (snsState.message === "") {
            vErrors.push("Message is required");
        }

        let attrs = snsState.messageAttributes != "" ? snsState.messageAttributes :  null;
        let dispatch = { ...snsState, messageAttributes: attrs };
        if (attrs) {
            try {
                const parsed = JSON.parse(attrs);
                dispatch.messageAttributes = parsed;
            } catch (error) {
                dispatch.messageAttributes = null;
                vErrors.push("If message attributes are provided, it must be a valid JSON");
            }
        }

        if (vErrors.length > 0) {
            setErrors(vErrors);
            return;
        }

        onSubmit(dispatch);
        setSnsState(initialSaveSNSState);
        setErrors([]);
    };

    return (
        <Modal isOpen={isOpen} onClose={() => onClose()}>
            <div className="saveBox vertical">
                <label>
                    Message
                    <textarea
                        type="text"
                        value={snsState.message}
                        onChange={(e) =>
                            setSnsState({
                                ...snsState,
                                message: e.target.value,
                            })
                        }
                    />
                </label>

                <label>
                    Message Attributes (Optional)
                    <textarea
                        type="text"
                        value={snsState.messageAttributes}
                        onChange={(e) =>
                            setSnsState({
                                ...snsState,
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

export default SaveSNS;
