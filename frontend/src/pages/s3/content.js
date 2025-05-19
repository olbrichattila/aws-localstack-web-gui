import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { delObject, listBucketContent, upload } from '../../api/s3';
import FilterBox from '../../components/filterBox';
import Spacer from '../../components/spacer';
import Button from '../../components/button';
import FileUploadModal from '../../components/fileUploadModal';
import InteractiveTable from '../../components/interactiveTable';
import { handleOpenS3Object } from '../../helpers';

const S3BucketContent = () => {
    const navigate = useNavigate()
    const { bucketName } = useParams();
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState('');
    const [uploadFileModalVisible, setUploadFileMOdalVisible] = useState(false);

    const onEvent = (e) => {
        if (e.name === 'Delete') {
            delObject(bucketName, e.i.Key).then(() => listBucketContent(bucketName).then(content => setData(content)));
        }

        if (e.name === 'View') {
            handleOpenS3Object(bucketName, e.i.Key);

        }
    }

    useEffect(() => {
        if (bucketName === '') {
            setData([]);
            return;
        }
        listBucketContent(bucketName).then(content => setData(content));
    }, [bucketName]);

    return (
        <>

            <FileUploadModal
                isOpen={uploadFileModalVisible}
                onClose={() => setUploadFileMOdalVisible(false)}
                onUploaded={fileName => {
                    upload(bucketName, fileName).then(() => listBucketContent(bucketName).then(content => setData(content)));
                    setUploadFileMOdalVisible(false);

                }} />
            <Button label="<< Back" margin={6} onClick={() => navigate('/aws/s3')} />
            <Button label="Upload file" margin={6} onClick={() => setUploadFileMOdalVisible(true)} />
            <h3>Bucket name: {bucketName}</h3>
            <FilterBox onSubmit={text => setFilter(text)} />
            <Spacer />
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: 'LastModified',
                        asc: false
                    },
                    filterField: 'Key',
                    columns: [
                        {
                            field: 'LastModified',
                            title: 'Last Modified',
                            clickable: false,
                        },
                        {
                            field: 'Key',
                            title: 'Name',
                            clickable: false,
                        },
                        {
                            field: 'Size',
                            title: 'File Size',
                            clickable: false,
                        },
                    ],
                    events: [
                        'View',
                        'Delete'
                    ],
                }}
                data={data}
                filter={filter}
                onEvent={e => onEvent(e)}
            />
        </>
    )
}

export default S3BucketContent;
