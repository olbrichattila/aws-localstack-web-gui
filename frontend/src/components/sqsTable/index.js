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
    watchUrl = '',
    onWatchChange = () => null,
    onDelete = () => null,
    onPurge = () => null,
    onRefresh = () => null,
    onSendMessage = () => null,
    onReadMessage = () => null,
}) => {
    const [watchUrlState, setWatchUrlState] = useState(watchUrl);
    const [sortInfo, setSortInfo] = useState(initialSortInfo);
    const [sorted, setSorted] = useState([]);

    const sortClick = (field) => {
        const asc = sortInfo.field === field ? !sortInfo.asc : true;
        setSortInfo({ field, asc })
        if (watchUrlState !== '') {
            setWatchUrlState('');
        }
    }

    const watchChange = (url) => {
        setWatchUrlState(url);
        onWatchChange(url)
    }

    useEffect(() => {
        setSorted(resort(data, sortInfo.field, sortInfo.asc))
    }, [sortInfo, data])


    useEffect(() => {
        if (watchUrlState !== '') {
            setWatchUrlState('');
        }
    }, [filter])

    useEffect(() => {
        setWatchUrlState(watchUrl);
    }, [watchUrl])
   
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
                {sorted.map(item  => {
                    const bucketName = item.url.split('/').pop();

                    if (filter !== '' && !bucketName.toLowerCase().includes(filter.toLocaleLowerCase())) {
                        return null;
                    }

                    return (
                        <tr className="sqsRow" key={item.url}>
                            <td>{bucketName}</td>
                            <td>{item.attributes.ApproximateNumberOfMessages}</td>
                            <td>{item.attributes.ApproximateNumberOfMessagesNotVisible}</td>
                            <td>{item.attributes.ApproximateNumberOfMessagesDelayed}</td>
                            <td className="narrow"><Button onClick={() => onDelete(item.url)} label="Delete" /></td>
                            <td className="narrow"><Button onClick={() => onPurge(item.url)} label="Purge" /></td>
                            <td className="narrow"><Button onClick={() => onRefresh(item.url)} label="Refresh" /></td>
                            <td className="narrow">
                                {watchUrlState !== item.url ? <Button label="Watch" onClick={() => watchChange(item.url)} /> : <LoadingSpinner onClick={() => watchChange('')} />}
                            </td>
                            <td className="narrow"><Button onClick={() => onSendMessage(item.url)} label="Send Message" /></td>
                            <td className="narrow"><Button onClick={() => onReadMessage(item.url)} label="Read Message" /></td>
                        </tr>
                    )
                })}
            </tbody>
        </table>
    );
}

export default SqsTable;
