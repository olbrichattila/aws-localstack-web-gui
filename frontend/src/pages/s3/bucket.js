import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAppContext } from "../../AppContext";
import Button from "../../components/button";
import FilterBox from "../../components/filterBox";
import Spacer from "../../components/spacer";
import InteractiveTable from "../../components/interactiveTable";
import SaveBox from "../../components/savebox";

const S3Bucket = () => {
    const { get, post, del } = useAppContext();
    const navigate = useNavigate();
    const [data, setData] = useState([]);
    const [filter, setFilter] = useState("");
    const [newBucketModelIsOpen, setNewBucketModelIsOpen] = useState(false);
    const [error, setError] = useState("");

    // API calls

    const loadBuckets = async () => {
        return get("/api/s3/buckets");
    };

    const newBucket = async (bucketName) => {
        return post("/api/s3/buckets", {
            bucketName,
        });
    };

    const delBucket = async (bucketName) => {
        return del("/api/s3/buckets", {
            bucketName,
        });
    };

    const onEvent = (e) => {
        if (e.name === "Delete") {
            delBucket(e.i.Name).then(() =>
                loadBuckets()
                    .then((buckets) => setData(buckets))
                    .catch((err) =>
                        setError(err.message ?? "Error fetching data")
                    )
            );
        }

        if (e.name === "clickable") {
            navigate(`/aws/s3/${encodeURIComponent(e.i.Name)}`);
        }
    };

    useEffect(() => {
        loadBuckets()
            .then((buckets) => setData(buckets))
            .catch((err) => setError(err.message ?? "Error fetching data"));
    }, []);

    useEffect(() => {
        let timeoutId = -1;
        if (error !== "") {
            timeoutId = setTimeout(() => {
                setError("");
            }, 6000);
        }

        return () => {
            if (timeoutId !== -1) {
                clearTimeout(timeoutId);
            }
        };
    }, [error]);

    return (
        <>
            <SaveBox
                title="New Bucket:"
                isOpen={newBucketModelIsOpen}
                onClose={() => setNewBucketModelIsOpen(false)}
                onSubmit={(name) => {
                    newBucket(name)
                        .then(() =>
                            loadBuckets().then((buckets) => {
                                setData(buckets);
                                setNewBucketModelIsOpen(false);
                            })
                        )
                        .catch((error) => {
                            setNewBucketModelIsOpen(false);
                            setError(error.message ?? "Error fetching data");
                        });
                }}
            />
            <Button
                label="Create new bucket"
                margin={6}
                onClick={() => setNewBucketModelIsOpen(true)}
            />
            {error !== "" && <div className="errorLine">{error}</div>}
            <FilterBox onSubmit={(text) => setFilter(text)} />
            <Spacer />
            <InteractiveTable
                structInfo={{
                    initialSort: {
                        field: "CreationDate",
                        asc: true,
                    },
                    filterField: "Name",
                    columns: [
                        {
                            field: "CreationDate",
                            title: "Created at",
                            clickable: false,
                        },
                        {
                            field: "Name",
                            title: "Bucket Name",
                            clickable: true,
                        },
                    ],
                    events: ["Delete"],
                }}
                data={data}
                filter={filter}
                onEvent={(e) => onEvent(e)}
            />
        </>
    );
};

export default S3Bucket;
