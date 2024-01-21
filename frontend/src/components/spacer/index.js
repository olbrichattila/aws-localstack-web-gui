import React from "react";
import "./index.scss";

const Spacer = ({space = 12}) => {
    return <div className="spacer" style={{paddingTop: space}} />
}

export default Spacer;
