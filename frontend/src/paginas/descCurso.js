import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { fetchdescCursos } from '../services/api';
import EnrollButton from './EnrollButton';
import './descCurso.css';

const descCurso = () => {
    const { courseId } = useParams();
    const [course, setCourse] = useState(null);

    useEffect(() => {
        fetchdescCursos(courseId).then(data => setCourse(data));
    }, [courseId]);

    if (!course) {
        return <div>Loading...</div>;
    }

    return (
        <div className="desc-curso">
            <h2>{course.title}</h2>
            <p>{course.description}</p>
            <p className="instructor">Instructor: {course.instructor}</p>
            <p className="duration">Duration: {course.duration}</p>
            <p className="requirements">Requirements: {course.requirements}</p>
            <EnrollButton courseId={courseId} className="enroll-button" />
        </div>
    );
};

export default descCurso;