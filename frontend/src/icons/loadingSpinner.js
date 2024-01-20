import React from 'react';
import { FaSpinner } from 'react-icons/fa';
import './loadingSpinner.scss';

const LoadingSpinner = ({onClick}) => {
  return (
    <div className="loading-spinner" onClick={() => onClick()}s>
      <FaSpinner className="spinner-icon" />
    </div>
  );
};

export default LoadingSpinner;