
import React from 'react';
import './index.scss'; // Import your modal styles

const Modal = ({ isOpen, onClose, children }) => {
  const overlayStyle = {
    display: isOpen ? 'block' : 'none',
  };

  return (
    <div className="modal-overlay" style={overlayStyle} onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <span className="close-button" onClick={onClose}></span>
        {children}
      </div>
    </div>
  );
};

export default Modal;
