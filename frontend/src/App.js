import './App.css';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import NavegadorHome from './componentes/NavegadorHome';
import CarruselDeCursos from './paginas/cursos';
import Home from './paginas/home';


import { useState } from 'react';

function App() {
  const [searchTerm, setSearchTerm] = useState('');

  return (
    <div>
      <Router>
        <NavegadorHome onSearch={setSearchTerm} />
        <Routes>
          <Route exact path="/home" element={<Home />} /> 
          <Route exact path="/cursos/todos" element={<CarruselDeCursos />} /> 

        </Routes>
      </Router>
      <footer />
    </div>
  );
}

export default App;
