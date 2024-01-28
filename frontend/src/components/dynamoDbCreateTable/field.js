import React from "react";
import { initialFieldState } from "./initialState";
import Button from "../button";
import './field.scss';

const Field = ({
    index = 0,
    field = initialFieldState,
    onChange = () => null,
    onDelete = () => null
}) => {
    return (
        <div className="fieldEditorField">
            <input
                type='text'
                value={field.attributeName}
                onChange={(e) => onChange(index, { ...field, attributeName: e.target.value })}
            />
            <select 
                onChange={(e) => onChange(index, { ...field, attributeType: e.target.value })}
                value={field.attributeType}
            >
                <option value="">Select attribute type</option>
                <option value="S">S</option>
                <option value="N">N</option>
                <option value="B">B</option>
            </select>

            <select
                onChange={(e) => onChange(index, { ...field, keyType: e.target.value })}
                value={field.keyType}
            >
                <option value="">Select key type</option>
                <option value="HASH">HASH</option>
                <option value="RANGE">RANGE</option>
            </select>
            <Button label="delete" onClick={() => onDelete(index)} />
        </div>
    )
}

export default Field;
