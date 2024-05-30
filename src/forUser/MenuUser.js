import React, { useState, useRef } from "react";
import { BsPersonCircle } from "react-icons/bs";
import { UseClickOutside } from "../UseClickOutside";

const MenuUser=({setShow}) =>
{
    const [isOpenb, setOpenb] = useState(false);
    const menuRef=useRef(null);
    
    UseClickOutside(menuRef, ()=>{
        if(isOpenb) setTimeout(()=>setOpenb(false), 50);
    });
    return (
        <header className="header">
            <div className="menu-user-button" onClick={() => setOpenb(!isOpenb)}>
                <BsPersonCircle />
            </div>
            <nav className={`menu ${isOpenb ? "active" : ""}`} ref={menuRef}>
                <ul className="menu-user-list">
                    {/* <li className="menu-user-item"><Link to="/">Мои билеты</Link></li>
                    <li className="menu-user-item"><Link to="/">Избранное</Link></li> */}
                    <li className="menu-user-item"><p><a href="/user-profile">Профиль</a></p></li>
                    <li className="menu-user-item"><p><a href="/log">Выход</a></p></li>
                </ul>
            </nav>
        </header>
    );
}
export default MenuUser;