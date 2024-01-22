import React, { useState } from "react";
import S3Bucket from "./bucket";
import S3BucketContent from "./content";

const S3Page = () => {
    const [bucketName, setBucketName] = useState('');

    return (
        <>
            {bucketName === '' && <S3Bucket onSelectBucket={bucket => setBucketName(bucket)} />}
            {bucketName !== '' && <S3BucketContent bucketName={bucketName} onBack={() => setBucketName('')} />}
        </>
    )
}

export default S3Page;
