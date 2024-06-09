import React, { useState, useEffect } from 'react';
import { Carousel } from 'react-responsive-carousel';
import Curso from './curso';
import axios from 'axios';
import 'react-responsive-carousel/lib/styles/carousel.min.css';
import './carrousel.css';

const CarruselDeCursos = () => {
    const [cursos, setCursos] = useState([]);

    useEffect(() => {
        const fetchCursos = async () => {
            try {
                const response = await axios.get('http://localhost:8080/prueba');
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
        <Carousel showThumbs={false}>
            {cursos.map(curso => (
                <Curso key={curso.id} curso={curso} />
            ))}
        </Carousel>
    );
};

export default CarruselDeCursos;