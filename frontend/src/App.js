import './App.css';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Navegador from './componentes/Navegador';
import Cursos from './paginas/cursos';
import Course from './paginas/carrousel';

import { useState } from 'react';

function App() {
  const [searchTerm, setSearchTerm] = useState('');

  return (
    <div>
      <Router>
        <Navegador onSearch={setSearchTerm} />
        <Routes>
          <Route exact path="/cursos" element={<Cursos />} />
          <Route exact path="/cursos/:id" element={<Course />} /> 

        </Routes>
      </Router>
      <footer />
    </div>
  );
}

export default App;
