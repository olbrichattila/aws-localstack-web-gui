import React, { useState, useEffect } from 'react';
import Button from '../button';
import SortBox from '../sortbox';
import { resort } from '../../helpers';

const initialSortInfo = {
    field: 'LastModified',
    asc: false
}

const S3BucketTable = ({
    data = [],
    filter = '',
    onDelete = () => null,
    onView = () => null,
}) => {
    const [sortInfo, setSortInfo] = useState(initialSortInfo);
    const [sorted, setSorted] = useState([]);

    const sortClick = (field) => {
        const asc = sortInfo.field === field ? !sortInfo.asc : true;
        setSortInfo({ field, asc })
    }

    useEffect(() => {
        setSorted(resort(data, sortInfo.field, sortInfo.asc))
    }, [sortInfo, data])


    return (
        // class sqsTable will have to move to generic stylesheet
        <table className='sqsTable'>
            <thead>
                <tr>
                    <th>
                        <SortBox
                            showArrow={sortInfo.field === 'LastModified'}
                            onClick={() => sortClick('LastModified')}
                            asc={sortInfo.asc}
                            title="Last Modified"
                        />
                    </th>
                    <th>
                        <SortBox
                            showArrow={sortInfo.field === 'Key'}
                            onClick={() => sortClick('Key')}
                            asc={sortInfo.asc}
                            title="Name"
                        />
                    </th>
                    <th>
                        <SortBox
                            showArrow={sortInfo.field === 'Size'}
                            onClick={() => sortClick('Size')}
                            asc={sortInfo.asc}
                            title="Size"
                        />
                    </th>
                    <th colSpan={2}></th>
                </tr>
            </thead>
            <tbody>
                {sorted.map((row, idx) => {
                    if (filter !== '' && !row.Key.toLowerCase().includes(filter.toLocaleLowerCase())) {
                        return null;
                    }

                    return (
                        <tr key={idx}>
                            <td>{row.LastModified}</td>
                            <td >{row.Key}</td>
                            <td >{row.Size}</td>
                            <td className="narrow">
                                <Button onClick={() => onView(row.Key)} label="View" />
                            </td>
                            <td className="narrow">
                                <Button onClick={() => onDelete(row.Key)} label="Delete" />
                            </td>
                        </tr>
                    )
                })}
            </tbody>
        </table>
    );
}

export default S3BucketTable;

