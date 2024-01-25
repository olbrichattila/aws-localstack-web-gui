import React, { useState } from "react";
import Button from "../button";
import Modal from "../modal";
import "./index.scss";

const SaveBox = ({
    isOpen = false,
    title = 'Name',
    onSubmit = () => null,
    onClose = () => null
}) => {
    const [text, setText] = useState('');

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
                        onChange={(e) => setText(e.target.value)}
                    />
                </label>
                <Button
                    onClick={() => {
                        setText('');
                        onSubmit(text);
                    }}
                    label="Save"
                />
            </div>
        </Modal>
    );
}

export default SaveBox
