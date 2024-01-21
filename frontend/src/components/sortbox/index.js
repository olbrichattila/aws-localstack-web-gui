import React from "react";
import "./index.scss";

const SortBox = ({ showArrow = true, title = '', asc = true, onClick = () => null }) => {
    return (
        <div className={`sortBox ${showArrow ? (asc ? 'asc' : 'desc') : ''}`} onClick={() => onClick()}>
            {title}
        </div>
    );
}

export default SortBox;
