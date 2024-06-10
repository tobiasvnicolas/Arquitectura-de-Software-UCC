import React, { useState } from 'react';
import { Link, useNavigate  } from 'react-router-dom';
import './NavegadorHome.css';

const NavegadorHome = ({ onSearch }) => {
  const [palabra, setSearchTerm] = useState('');
  const navigate = useNavigate();

  const handleSearchChange = (event) => {
    setSearchTerm(event.target.value);
  };

  const handleSearchSubmit = (event) => {
    event.preventDefault();
    onSearch(palabra);
    navigate(`/cursos/buscar?palabra=${palabra}`);
    setSearchTerm(''); 
    // Limpiar el término de búsqueda después de buscar
  };

  return (
    <nav className="navbar">
      <h2>CodeWave Learning</h2>
      <ul className="navlinks">
        <li><Link to="/">Home</Link></li>
        <li><Link to="/cursos/todos">Cursos</Link></li>
        <li>
          <form onSubmit={handleSearchSubmit}>
            <input
              type="text"
              value={palabra}
              onChange={handleSearchChange}
              placeholder="Buscar cursos..."
            />
            <button type="submit">Buscar</button>
          </form>
        </li>
      </ul>
    </nav>
  );
};

export default NavegadorHome;
