import { useState } from "react";
import Button from "../button";
import Modal from "../modal";
import "./index.scss";

const SaveBox = ({
    isOpen = false,
    numOnly = false,
    title = 'Name',
    onSubmit = () => null,
    onClose = () => null
}) => {
    const [text, setText] = useState('');

    const onTextChange = (value) => {
        if (numOnly) {
            setText(value.replace(/[^0-9]/g, ''))
            return
        }
        setText(value);
    }

    return (
        <Modal
            isOpen={isOpen}
            onClose={() => onClose()}
        >
            <div className="saveBox">
                <label>
                    {title}
                    <input
                        type='text'
                        value={text}
                        onChange={(e) => onTextChange(e.target.value)}
                    />
                </label>
                <Button
                    onClick={() => {
                        onTextChange('');
                        onSubmit(text);
                    }}
                    label="Save"
                />
            </div>
        </Modal>
    );
}

export default SaveBox
