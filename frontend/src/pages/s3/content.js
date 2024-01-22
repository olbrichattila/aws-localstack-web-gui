import React, { useEffect, useState } from 'react';
import { delObject, listBucketContent, upload } from '../../api/s3';
import S3BucketTable from '../../components/s3BucketTable';
import FilterBox from '../../components/filterBox';
import Spacer from '../../components/spacer';
import Button from '../../components/button';
import FileUploadModal from '../../components/fileUploadModal';
import { handleOpenS3Object } from '../../helpers';

const S3BucketContent = ({ bucketName = '', onBack = () => null }) => {
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState('');
    const [uploadFileModalVisible, setUploadFileMOdalVisible] = useState(false);

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
            <Button label="<< Back" margin={6} onClick={() => onBack()} />
            <Button label="Upload file" margin={6} onClick={() => setUploadFileMOdalVisible(true)} />
            <FilterBox onSubmit={text => setFilter(text)} />
            <Spacer />
            <S3BucketTable
                data={data}
                filter={filter}
                onDelete={key => delObject(bucketName, key).then(() => listBucketContent(bucketName).then(content => setData(content)))}
                onView={fileName => handleOpenS3Object(bucketName, fileName)}
            />
        </>
    )
}

export default S3BucketContent;
