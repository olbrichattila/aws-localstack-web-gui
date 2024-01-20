import React from "react";

const MenuOption = ({children, active = false, onClick = () => null}) => {
    return (
        <li
            className={active ? 'active': ''}
            onClick={() => onClick()}
        >
            {children}
        </li>
    );
}

export default MenuOption;
