import React, { useState } from "react";
import Field from "./field";
import Button from "../button";
import { initialFieldState } from "./initialState";
import './fieldEditor.scss';

const FieldEditor = ({ onSave = () => null }) => {
    const [fields, setFields] = useState([]);

    const newField = () => {
        setFields([...fields, initialFieldState])
    }

    const deleteField = (idx) => {
        setFields([...fields.slice(0, idx), ...fields.slice(idx + 1)])
    }

    const onFieldChange = (idx, state) => {
        setFields([...fields.slice(0, idx), { ...state }, ...fields.slice(idx + 1)]);
    }

    return (
        <div className="fieldEditor">
            <div className="fieldWrapper">
                {fields.map((field, idx) => <Field
                    key={idx}
                    index={idx}
                    field={field}
                    onDelete={idx => deleteField(idx)}
                    onChange={(idx, state) => onFieldChange(idx, state)}
                />)}

            </div>
            <Button margin={6} label="Add new field" onClick={() => newField()} />
            <Button margin={6} onClick={() => onSave(fields)} label="Create table" />
        </div>
    )
}

export default FieldEditor;