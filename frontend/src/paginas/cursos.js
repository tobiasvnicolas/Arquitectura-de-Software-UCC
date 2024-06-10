import React, { useState, useEffect } from 'react';
import { Carousel } from 'react-responsive-carousel';
import Curso from './curso';
import axios from 'axios';
import 'react-responsive-carousel/lib/styles/carousel.min.css';
import './cursos.css';

const CarruselDeCursos = () => {
    const [cursos, setCursos] = useState([]);

    useEffect(() => {
        const fetchCursos = async () => {
            try {
                const response = await axios.get('http://localhost:8080/cursos/todos');
                setCursos(response.data);
            } catch (error) {
                console.error('Error fetching cursos:', error);
            }
        };

        fetchCursos();
    }, []);

    if (cursos.length === 0) {
        return <div>Cargando cursos...</div>;
    }

    return (
        <div className="Courses">  
            <div className="Course"> 
                <Carousel showThumbs={false} className="carousel">
                    {cursos.map(curso => (
                        <Curso key={curso.curso_id} curso={curso} />
                        
                    ))}
                    
                </Carousel>

            </div> 
        </div>

    );
};

export default CarruselDeCursos;