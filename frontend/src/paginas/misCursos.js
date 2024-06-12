import React, { useEffect, useState } from 'react';
import { fetchmisCursos } from '../services/api';
import './misCursos.css';

const MisCursos = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        fetchmisCursos().then(data => setCourses(data));
    }, []);

    if (courses.length === 0) {
        return <div className="loading">Loading...</div>;
    }

    return (
        <div className="mis-cursos">
            <h2>My Courses</h2>
            <ul>
                {courses.map(course => (
                    <li key={course.id}>
                        <a href={`/course/${course.id}`}>{course.title}</a>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default MisCursos;