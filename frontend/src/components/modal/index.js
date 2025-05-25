import { useEffect, useRef, useState } from "react";
import "./index.scss"; // Import your modal styles

const Modal = ({ isOpen, onClose, children }) => {
    const modalRef = useRef(null);
    const overlayStyle = {
        display: isOpen ? "block" : "none",
    };

    useEffect(() => {
        if (modalRef.current === null) return;

        var timer = setTimeout(() => {
            if (isOpen) {
                modalRef.current.classList.add("visible");
            } else {
                modalRef.current.classList.remove("visible");
            }
        }, 10);

        return () => {
            clearTimeout(timer);
        };
    }, [isOpen]);

    return (
        <div
            className="modal-overlay"
            style={overlayStyle}
            onClick={onClose}
            ref={modalRef}
        >
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
                <span className="close-button" onClick={onClose}></span>
                {children}
            </div>
        </div>
    );
};

export default Modal;
