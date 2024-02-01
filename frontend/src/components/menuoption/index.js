import React from "react";
import { Link, useLocation } from "react-router-dom";
import "./index.scss";

const MenuOption = ({ children,  to = '/' }) => {
    const location = useLocation();
    const regexEscapedPathname = to.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
    const regex = new RegExp(`^${regexEscapedPathname}.*$`);
    const active = regex.test(location.pathname);

    return (
        <li className={active ? 'active' : ''}>
            <Link className="menuOptionLink" to={to}>
                {children} 
            </Link>
        </li>
    );
}

export default MenuOption;
