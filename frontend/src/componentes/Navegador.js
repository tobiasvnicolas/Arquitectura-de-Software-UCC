import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import   './Navegador.css';

const Navegador = ({ onSearch }) => {
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
      <ul className="navlinks">
        <li><Link to="/">Home</Link></li>
        <li><Link to="/Cursos">Cursos</Link></li>
        <li><Link to="/cursos/:id">Registrarse</Link></li>
        <li><Link to="/IniciarSesion">Iniciar Sesión</Link></li>
        <li>
          <form onSubmit={handleSearchSubmit}>
            <input 
              type="text" 
              placeholder="Search..." 
              value={searchTerm} 
              onChange={handleSearchChange} 
            />
            <button type="submit">Search</button>
          </form>
        </li>
      </ul>
    </nav>
  );
};

export default Navegador;