import React, { useEffect, useState } from 'react';
import LoadingSpinner from '../../icons/loadingSpinner';
import Button from '../button';
import './index.scss';
import { resort } from '../../helpers';
import SortBox from '../sortbox';

const initialSortInfo = {
    field: 'attributes.ApproximateNumberOfMessages',
    asc: true
}

const SqsTable = ({
    data = [],
    filter = '',
    watchIdx = -1,
    onWatchChange = () => null,
    onDelete = () => null,
    onPurge = () => null,
    onRefresh = () => null,
    onSendMessage = () => null,
    onReadMessage = () => null,
}) => {
    const [watch, setWatch] = useState(watchIdx);
    const [sortInfo, setSortInfo] = useState(initialSortInfo);
    const [sorted, setSorted] = useState([]);

    const sortClick = (field) => {
        const asc = sortInfo.field === field ? !sortInfo.asc : true;
        setSortInfo({ field, asc })
        if (watch >= 0) {
            setWatch(-1);
        }
    }

    const watchChange = (idx) => {
        setWatch(idx);
        onWatchChange(idx)
    }

    useEffect(() => {
        setSorted(resort(data, sortInfo.field, sortInfo.asc))
    }, [sortInfo, data])


    useEffect(() => {
        if (watch >= 0) {
            setWatch(-1);
        }
    }, [filter])

    useEffect(() => {
        setWatch(watchIdx);
    }, [watchIdx])
    
    return (
        <table className='sqsTable'>
            <thead>
                <tr>
                    <th>
                        <SortBox
                            showArrow={sortInfo.field === 'url'}
                            onClick={() => sortClick('url')}
                            asc={sortInfo.asc}
                            title="Queue Name"
                        />
                    </th>
                    <th className="narrow">
                        <SortBox
                            showArrow={sortInfo.field === 'attributes.ApproximateNumberOfMessages'}
                            onClick={() => sortClick('attributes.ApproximateNumberOfMessages')}
                            asc={sortInfo.asc}
                            title="Messages"
                        />
                    </th>
                    <th className="narrow">
                        <SortBox
                            showArrow={sortInfo.field === 'attributes.ApproximateNumberOfMessagesNotVisible'}
                            onClick={() => sortClick('attributes.ApproximateNumberOfMessagesNotVisible')}
                            asc={sortInfo.asc}
                            title="Messages NotVisible"
                        />
                    </th>
                    <th className="narrow">
                        <SortBox
                            showArrow={sortInfo.field === 'attributes.ApproximateNumberOfMessagesDelayed'}
                            onClick={() => sortClick('attributes.ApproximateNumberOfMessagesDelayed')}
                            asc={sortInfo.asc}
                            title="Messages Delayed"
                        />
                    </th>
                    <th colSpan={6}></th>
                </tr>
            </thead>
            <tbody>
                {sorted.map((item, idx) => {
                    const bucketName = item.url.split('/').pop();

                    if (filter !== '' && !bucketName.toLowerCase().includes(filter.toLocaleLowerCase())) {
                        return null;
                    }

                    return (
                        <tr className="sqsRow" key={idx}>
                            <td>{bucketName}</td>
                            <td>{item.attributes.ApproximateNumberOfMessages}</td>
                            <td>{item.attributes.ApproximateNumberOfMessagesNotVisible}</td>
                            <td>{item.attributes.ApproximateNumberOfMessagesDelayed}</td>
                            <td className="narrow"><Button onClick={() => onDelete(idx)} label="Delete" /></td>
                            <td className="narrow"><Button onClick={() => onPurge(idx)} label="Purge" /></td>
                            <td className="narrow"><Button onClick={() => onRefresh(idx)} label="Refresh" /></td>
                            <td className="narrow">
                                {watch !== idx ? <Button label="Watch" onClick={() => watchChange(idx)} /> : <LoadingSpinner onClick={() => watchChange(-1)} />}
                            </td>
                            <td className="narrow"><Button onClick={() => onSendMessage(idx)} label="Send Message" /></td>
                            <td className="narrow"><Button onClick={() => onReadMessage(idx)} label="Read Message" /></td>
                        </tr>
                    )
                })}
            </tbody>
        </table>
    );
}

export default SqsTable;


