import React, { useEffect, useState } from 'react';
import { fetchMyCourses } from '../services/api';

const MyCourses = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        fetchMyCourses().then(data => setCourses(data));
    }, []);

    return (
        <div>
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

export default MyCourses;
