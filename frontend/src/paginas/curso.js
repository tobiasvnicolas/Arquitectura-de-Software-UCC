// Curso.js
import React from 'react';

const Curso = ({ curso }) => {
    return (
        <div className="curso">
            <h3>{curso.nombre}</h3>
            <p>{curso.descripcion}</p>
            <p>{curso.categoria}</p>
        </div>
    );
};

export default Curso;