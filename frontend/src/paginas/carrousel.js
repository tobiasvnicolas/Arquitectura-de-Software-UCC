import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import 'react-responsive-carousel/lib/styles/carousel.css';
import { Carousel } from 'react-responsive-carousel';
import './carrousel.css';


const CarruselDeCursos = () => {
    const { id } = useParams();
    const [cursos, setCursos] = useState([]);

    const fetchCursos = async () => {
       
            const response = await fetch(`http://localhost:3000/cursos/${id}`);
            const data = await response.json();
            console.log(data);

            if (data && data.busqueda) {
                const cursosConIds = data.busqueda.map((curso, index) => ({
                    ...curso,
                    id: curso.curso_id, // Asigna un id único, empezando desde 1 (si es necesario)
                }));
                setCursos(cursosConIds);
            } else {
                console.error('Error: Estructura de datos incorrecta en la respuesta');
            }
     
    } 

    useEffect(()=>{
        fetchCursos()
    });
       
   

    if (cursos.length === 0) {
        return <div>Cargando...</div>;
    }

    return (
        <Carousel>
            {cursos.map((curso) => (
                <div key={curso.curso_id}>
                    <div className="legend">
                        <h3>{curso.nombre}</h3>
                        <p>{curso.descripcion}</p>
                        <p>{curso.categoria}</p>
                    </div>
                </div>
            ))}
        </Carousel>
    );
};

export default CarruselDeCursos;

