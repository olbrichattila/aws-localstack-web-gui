import React, { useState } from 'react';
import Modal from '../modal';
import Button from '../button';
import { newBucket } from '../../api/s3'
import './index.scss';

const NewS3BucketModel = ({ isOpen, onClose, onSaved }) => {
    const [bucketName, setBucketName] = useState('');
    const addBucket = () => {
        newBucket(bucketName).then(() => onSaved())
    }

    return (
        <Modal
            isOpen={isOpen}
            onClose={onClose}
        >
            <div className='newS3Bucket'>
                <label>
                    Bucket name
                    <input
                        type='text'
                        value={bucketName}
                        onChange={(e) => setBucketName(e.target.value)}
                    />

                    <Button label="Create" margin={6} onClick={() => addBucket()} />
                </label>
            </div>
        </Modal>
    );

}

export default NewS3BucketModel;
