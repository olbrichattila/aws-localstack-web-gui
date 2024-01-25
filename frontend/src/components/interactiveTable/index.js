import React, { useEffect, useState } from 'react';
import LoadingSpinner from '../../icons/loadingSpinner';
import Button from '../button';
import './index.scss';
import { resort, valueByPath } from '../../helpers';
import SortBox from '../sortbox';

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
                    {structInfo.columns.map(struct =>
                        <th>
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
                {sorted.map(item => {
                    const filterValue = valueByPath(item, structInfo.filterField);
                    if (filter !== '' && !filterValue.toLowerCase().includes(filter.toLocaleLowerCase())) {
                        return null;
                    }

                    return (
                        <tr className="sqsRow" key={item.TopicArn}>
                            {structInfo.columns.map(struct => {
                                return struct.clickable ?
                                    <td className="clickable" onClick={() => onEvent({ name: 'clickable', i: item, w: true })}>{valueByPath(item, struct.field)}</td> :
                                    <td>{valueByPath(item, struct.field)}</td>
                            }
                            )}

                            {structInfo.events.map(e => <td className="narrow">
                                {
                                    filterValue === watch && structInfo.watchButton === e ?
                                        <LoadingSpinner onClick={() => onEvent({ name: e, i: item, w: false })} /> :
                                        <Button onClick={() => onEvent({ name: e, i: item, w: true })} label={e} />
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
