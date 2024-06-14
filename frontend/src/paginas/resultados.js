import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import './resultados.css';

function useQuery() {
  return new URLSearchParams(useLocation().search);
}

function Resultados({ searchTerm }) {
  const [cursos, setCursos] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    if (searchTerm) {
      setLoading(true);
      setError(null);
      fetch(`http://localhost:8080/cursos/buscar?palabra=${encodeURIComponent(searchTerm)}`)
        .then(response => {
          if (!response.ok) {
            throw new Error('Error en la solicitud');
          }
          return response.json();
        })
        .then(data => {
          if (data && Array.isArray(data.results)) {
            setCursos(data.results);
          } else {
            setCursos([]);
            setError('No se encontraron resultados o estructura de datos inválida');
          }
        })
        .catch(error => {
          console.error('Error fetching courses:', error);
          setError('Error al obtener los datos');
        })
        .finally(() => setLoading(false));
    } else {
      // Si no hay término de búsqueda, podrías querer limpiar los resultados o manejarlo de otra manera
      setCursos([]);
    }
  }, [searchTerm]);

  return (
    <div className="Resultadosfondo">
      {loading ? (
        <img
          style={{ width: '36px' }}
          src="https://raw.githubusercontent.com/Codelessly/FlutterLoadingGIFs/master/packages/circular_progress_indicator_square_large.gif"
          alt="Loading"
        />
      ) : error ? (
        <p>{error}</p>
      ) :(
        <div className="course-list">
          {cursos.length > 0 ? (
            cursos.map(curso => (
              <div key={curso.curso_id} className="course-list-item">
                <div>
                  <p><strong>{curso.nombre}</strong></p>
                  <p>{curso.descripcion}</p>
                  <p>{curso.categoria}</p>
                </div>
                <hr />
              </div>
            ))
          ) : (
            <p>No se encontraron cursos.</p>
          )}
        </div>
      )}
    </div>
  );
}

export default Resultados;
