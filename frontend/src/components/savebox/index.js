import React, { useState } from "react";
import Button from "../button";
import "./index.scss";

const SaveBox = ({ title = 'Name', onSubmit = () => null }) => {
    const [text, setText] = useState('');

    return <div className="saveBox">
        <label>
            {title}
            <input
                type='text'
                value={text}
                onChange={(e) => setText(e.target.value)}
            />
        </label>
        <Button
            onClick={() => onSubmit(text)}
            label="Save"
        />
    </div>
}

export default SaveBox
