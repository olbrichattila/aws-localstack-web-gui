import React from 'react';
import './index.scss';

const Button = ({ label, onClick, margin = 0 }) => {
  return (
    <button className="custom-button" onClick={onClick} style={{margin}}>
      {label}
    </button>
  );
};

export default Button;
