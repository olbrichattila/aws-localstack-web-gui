import React, { useState } from "react";
import Button from "../button";
import "./index.scss";

const FilterBox = ({ onSubmit = () => null }) => {
    const [text, setText] = useState('');

    return <div className="filterBox">
        <input
            type='text'
            value={text}
            onChange={(e) => setText(e.target.value)}
        />
        <Button
            onClick={() => onSubmit(text)}
            label="Filter result"
        />
    </div>
}

export default FilterBox;
