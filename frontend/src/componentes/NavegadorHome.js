import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import   './NavegadorHome.css';

const NavegadorHome = ({ onSearch }) => {
  const [searchTerm, setSearchTerm] = useState('');

  const handleSearchChange = (event) => {
    setSearchTerm(event.target.value);
  };

  const handleSearchSubmit = (event) => {
    event.preventDefault();
    onSearch(searchTerm);
  };

  return (
    <nav className="navbar">
      <h2 >CodeWave Learning</h2>
      <ul className="navlinks">
        <li><Link to="/home">Home</Link></li>
        <li><Link to="/cursos/:id">Registrarse</Link></li>
        <li><Link to="/cursos/todos">Iniciar Sesión</Link></li>
        </ul>
    </nav>
  );
};

export default NavegadorHome;