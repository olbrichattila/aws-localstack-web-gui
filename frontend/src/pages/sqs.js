import React, { useState, useEffect } from 'react';
import './sqs.scss';
import Button from '../components/button';
import { load, save, del, purge, refresh } from '../api/sqs'
import LoadingSpinner from '../icons/loadingSpinner';
import SaveBox from '../components/savebox';

const SqsPage = () => {
  const [data, setData] = useState([]);
  const [watch, setWatch] = useState(-1);

  const refreshQueue = (attributes, idx) => {
    setData(
      [...data.slice(0, idx), {...data[idx], attributes}, ...data.slice(idx + 1)]
    );
  }

  useEffect(() => {
    const updateTimer = () => {
      if (watch >= 0) {
        refresh(data[watch].url).then(r => refreshQueue(r, watch))
      }
    };

    const timerId = setInterval(updateTimer, 2000);

    return () => {
      clearInterval(timerId);
      
    };
  }, [watch]);

  useEffect(() => {
    load().then(r => setData(r));
  }, []);

  return (
    <div>
      <SaveBox
        title='New SQS queue name:'
        onSubmit={queueName => save(queueName).then(() => load().then(r => setData(r)))}
      />
      <h1>Data from JSON API</h1>
      <table className='sqsTable'>
        <thead>
          <tr>
            <th>Queue Name</th>
            <th className="narrow">Messages</th>
            <th className="narrow">Messages NotVisible</th>
            <th className="narrow">Messages Delayed</th>
            <th className="narrow">Delete</th>
            <th className="narrow">Purge</th>
            <th className="narrow">Refresh</th>
            <th className="narrow"></th>
          </tr>
        </thead>
        <tbody>
        {data.map((item, idx) => (
          <tr className="sqsRow" key={idx}>
            <td>{item.url.split('/').pop()}</td>
            <td>{item.attributes.ApproximateNumberOfMessages}</td>
            <td>{item.attributes.ApproximateNumberOfMessagesNotVisible}</td>
            <td>{item.attributes.ApproximateNumberOfMessagesDelayed}</td>
            <td>{<Button onClick={() => del(data[idx].url).then(() => load().then(r => setData(r)))} label="Delete" />}</td>
            <td>{<Button onClick={() => purge(data[idx].url).then(() => load().then(r => setData(r)))} label="Purge" />}</td>
            <td>{<Button onClick={() => refresh(data[idx].url).then(r => refreshQueue(r, idx))} label="Refresh" />}</td>
            <td>
              {watch !== idx ? <Button label="Watch" onClick={() => setWatch(idx)}/> : <LoadingSpinner onClick={() => setWatch(-1)} />}
            </td>
          </tr>
        ))}
        </tbody>
      </table>
    </div>
  );
};

export default SqsPage;
