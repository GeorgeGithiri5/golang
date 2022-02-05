import React from "react"
import { Link } from "react-router-dom"

const Navbar = ()=>{
    return(
        <div>
            <nav className="navbar navbar-expand navbar-dark bg-dark mb-4">
                <div className="container-fluid">
                <Link to='/' className="navbar-brand">Home</Link>
                <div>
                    <ul className="navbar-nav me-auto mb-2 mb-md-0">
                    <li className="nav-item">
                        <Link to='/login' className="nav-link" aria-current="page">Login</Link>
                    </li>
                    <li className="nav-item">
                        <Link to='/register' className="nav-link" aria-current="page">Register</Link>
                    </li>
                    </ul>
                </div>
                </div>
            </nav>
        </div>
    )
}

export default Navbar