import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { fetchdescCursos } from '../services/api';
import EnrollButton from './EnrollButton';
import './descCurso.css';

const descCurso = () => {
    const { curso_id } = useParams();
    const [curso, setCurso] = useState(null);

    useEffect(() => {
        fetchdescCursos(curso_id).then(data => setCurso(data));
    }, [curso_id]);

    if (!curso) {
        return <div>Cargando...</div>;
    }

    return (
        <div className="desc-curso">
            <p>{curso.curso_id}</p>
            <p>{curso.nombre}</p>
            <p>{curso.descripcion}</p>
            <p>{curso.categoria}</p>
            <EnrollButton CursoID={CursoID} className="enroll-button" />
        </div>
    );
};

export default descCurso;