import React, { useState } from 'react';
import Modal from '../modal';
import FieldEditor from './fieldEditor';
import './index.scss';

const DynamoDbCreateTable = ({
    isOpen = false,
    onClose = () => null,
    onSave = () => null
}) => {
    const [tableName, setTableName] = useState('');
    const [errors, setErrors] = useState([]);

    const save = (fields) => {
        const payload = {
            name: tableName,
            fields,
        };

        const validationErrors = validate(payload);

        if (validationErrors.length === 0) {
            setErrors([]);
            onSave(payload);
            return
        }

        setErrors(validationErrors);
    }

    const validate = (payload) => {
        const errors = [];
        if (payload.name === '') {
            errors.push('Table name is required');
        }

        const fieldErrors = payload.fields.reduce((prev, current, idx) => {
            if (current.attributeName === '') {
                prev.push(`Attribute name is missing, row: ${idx}`)

            }
            if (current.attributeType === '') {
                prev.push(`Attribute type is missing, row: ${idx}`)
            }
            if (current.keyType === '') {
                prev.push(`Key type is missing, row: ${idx}`)
            }

            return prev;
        }, []);

        return [...errors, ...fieldErrors];
    }


    return (
        <Modal
            isOpen={isOpen}
            onClose={() => onClose()}
        >
            <div className='dynamoDbCreateTableWrapper'>
                <label>
                    Table Name:
                    <input
                        type='text'
                        value={tableName}
                        onChange={(e) => setTableName(e.target.value)}
                    />
                </label>
                <hr />
                <label>Fields:</label>
                {errors.length > 0 && errors.map(e => <div className='errorLine'>{e}</div>)}
                <FieldEditor onSave={fields => save(fields)} />
            </div>
        </Modal>
    )

}

export default DynamoDbCreateTable;