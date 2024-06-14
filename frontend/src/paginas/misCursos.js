import React, { useEffect, useState } from 'react';
import { fetchmisCursos } from '../services/api';
import axios from 'axios';
import './misCursos.css';

const MisCursos = () => {
    const { id } = useParams();
    const [cursos, setCursos] = useState([]);

    useEffect(() => {
        const fetchCursos = async () => {
            try {
                const response = await axios.get('http://localhost:8080/inscripcion/${id}');
                setCursos(response.data);
            } catch (error) {
                console.error('Error fetching cursos:', error);
            }
        };

        fetchCursos();
        //fetchmisCursos().then(data => setCourses(data));
        //esto era lo original
    }, [id]);

    if (cursos.length === 0) {
        return <div className="loading">Loading...</div>;
    }

    return (
        <div className="mis-cursos">
            <h2>My Courses</h2>
            <ul>
                {cursos.map(curso => (
                    <li key={curso.curso_id}>
                        <a href={`/inscripcion/${curso.curso_id}`}>{curso.nombre}</a>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default MisCursos;