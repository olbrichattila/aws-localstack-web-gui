import React, { useEffect, useState } from 'react';
import Modal from '../modal';
import Button from '../button';
import './index.scss';

const FileUploadModal = ({ isOpen, onClose, onUploaded = () => null }) => {
    const [file, setFile] = useState(null);
    const [error, setError] = useState('');

    useEffect(() => {
        if (error !== '') {
            setError('');
        }
    }, [isOpen]);

    useEffect(() => {
        const timeoutId = setTimeout(() => {
            if (error !== '') {
                setError('');
            }
        }, 4000);
        return () => {
            clearTimeout(timeoutId);
        };
    }, [error])


    const handleFileUpload = () => {
        const formData = new FormData();
        formData.append('file', file);

        fetch(`http://localhost:8080/api/s3/file_upload`, {
//        fetch(`${process.env.REACT_APP_API_URL}/api/s3/file_upload`, {
            method: 'POST',
            body: formData,
        })
            .then(response => {
                if (response.ok) {
                    onUploaded(file.name);
                } else {
                    setError(response.statusText ?? 'unknown error');
                }
            })
            .catch(error => {
                setError('Error uploading file');
            });
    };

    return (
        <Modal
            isOpen={isOpen}
            onClose={onClose}
        >
            <div className='fileUploadModal'>
                <input type="file" onChange={e => setFile(e.target.files[0])} />
                {error !== '' && <div className='fileUploadError'>Error uploading file<br />{error}</div>}
                <Button
                    label="Upload File"
                    onClick={() => handleFileUpload()}
                />
            </div>
        </Modal>
    );
};

export default FileUploadModal;
