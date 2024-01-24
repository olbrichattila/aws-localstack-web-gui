import React, { useEffect, useState } from "react";
import Button from "../../components/button";
import FilterBox from "../../components/filterBox";
import S3Table from "../../components/s3Table";
import Spacer from "../../components/spacer";
import NewS3BucketModel from "../../components/newS3BucketModel";
import { loadBuckets, delBucket } from "../../api/s3";

const S3Bucket = ({onSelectBucket = () => null}) => {
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState('');
    const [newBucketModelIsOpen, setNewBucketModelIsOpen] = useState(false);
    
    useEffect(() => {
        loadBuckets().then(buckets => setData(buckets));
    }, []);

    return (
        <>
            <NewS3BucketModel
                isOpen={newBucketModelIsOpen}
                onClose={() => setNewBucketModelIsOpen(false)}
                onSaved={() => {
                    loadBuckets().then(buckets => setData(buckets));
                    setNewBucketModelIsOpen(false);

                }}
            />
            <Button label="Create new bucket" margin={6} onClick={() => setNewBucketModelIsOpen(true)} />
            <FilterBox onSubmit={text => setFilter(text)} />
            <Spacer />
            <S3Table
                data={data} filter={filter}
                onDelete={bucketName => delBucket(bucketName).then(() => loadBuckets().then(buckets => setData(buckets)))}
                onSelectBucket={bucket => onSelectBucket(bucket)}
            />
        </>
    )
}

export default S3Bucket;
