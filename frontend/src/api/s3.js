import { get, post, del } from './request';

const loadBuckets = async () => {
    return get('/api/s3/buckets');
}

const newBucket = async (bucketName) => {
    return post('/api/s3/buckets', {
        bucketName,
    });
}

const delBucket = async (bucketName) => {
    return del('/api/s3/buckets', {
        bucketName,
    });
}

const listBucketContent = async (bucketName) => {
    return get(`/api/s3/list/${encodeURIComponent(bucketName)}`)
}

const delObject = async (bucketName, key) => {
    return del('/api/s3/buckets/delete/object', {
        bucketName,
        fileName: key,
    });
}

const upload = async (bucketName, fileName) => {
    return post('/api/s3/buckets/upload', {
        bucketName,
        fileName,
    });
}

export { loadBuckets, newBucket, delBucket, listBucketContent, delObject, upload }
