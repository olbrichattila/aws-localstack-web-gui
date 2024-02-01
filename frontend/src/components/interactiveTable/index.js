import React, { useEffect, useState } from 'react';
import LoadingSpinner from '../../icons/loadingSpinner';
import Button from '../button';
import SortBox from '../sortbox';
import { resort, valueByPath } from '../../helpers';
import './index.scss';

const InteractiveTable = ({
    structInfo = {},
    data = [],
    filter = '',
    watch = '',
    onEvent = () => null,
}) => {
    const [sortInfo, setSortInfo] = useState(structInfo.initialSort);
    const [sorted, setSorted] = useState([]);

    const sortClick = (field) => {
        const asc = sortInfo.field === field ? !sortInfo.asc : true;
        setSortInfo({ field, asc })
    }

    useEffect(() => {
        setSorted(resort(data, sortInfo.field, sortInfo.asc))
    }, [sortInfo, data])

    return (
        <table className='sqsTable'>
            <thead>
                <tr>
                    {structInfo.columns.map((struct, colidx) =>
                        <th key={`col_${colidx}`}>
                            <SortBox
                                showArrow={sortInfo.field === struct.field}
                                onClick={() => sortClick(struct.field)}
                                asc={sortInfo.asc}
                                title={struct.title}
                            />
                        </th>
                    )}
                    <th colSpan={structInfo.events.length}></th>
                </tr>
            </thead>
            <tbody>
                {sorted.map((item, idx) => {
                    const filterValue = valueByPath(item, structInfo.filterField);
                    if (filter !== '' && !filterValue.toLowerCase().includes(filter.toLocaleLowerCase())) {
                        return null;
                    }

                    return (
                        <tr className="sqsRow" key={`sorted_${idx}`}>
                            {structInfo.columns.map((struct, idx) => {
                                return struct.clickable ?
                                    <td key={`clickable_${idx}`} className="clickable" onClick={() => onEvent({ name: 'clickable', i: item, w: true })}>{valueByPath(item, struct.field)}</td> :
                                    <td key={`clickable_${idx}`}>{valueByPath(item, struct.field)}</td>
                            }
                            )}

                            {structInfo.events.map((e, idx) => <td className="narrow">
                                {
                                    filterValue === watch && structInfo.watchButton === e ?
                                        <LoadingSpinner key={`event_${idx}`} onClick={() => onEvent({ name: e, i: item, w: false })} /> :
                                        <Button key={idx} onClick={() => onEvent({ name: e, i: item, w: true })} label={e} />
                                }
                            </td>)}
                        </tr>
                    )
                })}
            </tbody>
        </table>
    );
}

export default InteractiveTable;
