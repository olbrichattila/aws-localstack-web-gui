import React, { useEffect, useState } from "react";
import Button from "../../components/button";
import FilterBox from "../../components/filterBox";
import Spacer from "../../components/spacer";
import NewS3BucketModel from "../../components/newS3BucketModel";
import { loadBuckets, delBucket } from "../../api/s3";
import InteractiveTable from "../../components/interactiveTable";

const S3Bucket = ({ onSelectBucket = () => null }) => {
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState('');
    const [newBucketModelIsOpen, setNewBucketModelIsOpen] = useState(false);

    const onEvent = (e) => {
        if (e.name === 'Delete') {
            delBucket(e.i.Name).then(() => loadBuckets().then(buckets => setData(buckets)));
        }

        if (e.name === 'clickable') {
            onSelectBucket(e.i.Name);
        }
    }

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
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: 'CreationDate',
                        asc: true
                    },
                    filterField: 'Name',
                    columns: [
                        {
                            field: 'CreationDate',
                            title: 'Created at',
                            clickable: false,
                        },
                        {
                            field: 'Name',
                            title: 'Bucket Name',
                            clickable: true,
                        },
                    ],
                    events: [
                        'Delete'
                    ],
                }}
                data={data} filter={filter}
                onEvent={e => onEvent(e)}

            />
        </>
    )
}

export default S3Bucket;
