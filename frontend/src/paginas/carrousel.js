import { useParams } from 'react-router-dom';
import React, { useEffect, useState } from 'react';

function Course() {
  const { id } = useParams();
  const [ curso, setCourse ] = useState({});

  useEffect(() => {
    if (id == null) {
      return
    }
    console.log(`Fetching data from http://localhost:8080/cursos/${id}`)
    // Fetch data from the API
    fetch(`http://localhost:8080/cursos/${id}`)
      .then(response => response.json())
      .then(data => setCourse(data))
      .catch(error => console.error('Error fetching courses:', error));
  }, [curso]);

  return (
    <>
      <div className="Courses">
        {curso != null ? (
          <>
            <div key={curso.curso_id} className="Course">
                <>
                  <div className="Course-details">
                    <h1 className="Course-title">{curso.nombre}</h1>
                    <p className="Course-description">{curso.descripcion}</p>
                    <p className="Course-category"><strong>{curso.categoria}</strong></p>
                  </div>
                </>
            </div>
            <form><input type="submit" value="Subscribe"/></form>
          </>
        ) : (
          <p>Loading course...</p>
        )}
      </div>
    </>
  );
}

export default Course;