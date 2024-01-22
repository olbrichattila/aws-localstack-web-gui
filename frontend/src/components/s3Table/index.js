import React, { useState, useEffect } from 'react';
import Button from '../button';
import SortBox from '../sortbox';
import { resort } from '../../helpers';

const initialSortInfo = {
    field: 'CreationDate',
    asc: true
}

const S3Table = ({ data = [], filter = '', onDelete = () => null, onSelectBucket = () => null }) => {
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
                            showArrow={sortInfo.field === 'CreationDate'}
                            onClick={() => sortClick('CreationDate')}
                            asc={sortInfo.asc}
                            title="Created at"
                        />
                    </th>
                    <th>
                        <SortBox
                            showArrow={sortInfo.field === 'Name'}
                            onClick={() => sortClick('Name')}
                            asc={sortInfo.asc}
                            title="Name"
                        />
                    </th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                {sorted.map((row, idx) => {
                    if (filter !== '' && !row.Name.toLowerCase().includes(filter.toLocaleLowerCase())) {
                        return null;
                    }

                    return (
                        <tr key={idx}>
                            <td>{row.CreationDate}</td>
                            <td className="clickable" onClick={() => onSelectBucket(row.Name)} >{row.Name}</td>
                            <td className="narrow">
                                <Button onClick={() => onDelete(row.Name)} label="Delete" />
                            </td>
                        </tr>
                    )
                })}

            </tbody>
        </table>
    );
}

export default S3Table;

