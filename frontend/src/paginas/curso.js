// Curso.js
import React from 'react';


const Curso = ({ curso }) => {
    return (

        <div className="Curso">
            <br/><br/><br/>
            <h3>{curso.nombre}</h3>
            <p>{curso.descripcion}</p>
            <p><b>Categoria:</b> {curso.categoria}</p>
            <br/>
            <button type="submit">Inscribirme</button>
            <br/><br/> <br/>

        </div>
    
    );
};

export default Curso;