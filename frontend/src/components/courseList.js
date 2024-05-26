import React, { useEffect, useState } from 'react';
import { fetchCourses } from '../services/api';

const CourseList = () => {
    const [courses, setCourses] = useState([]);

    useEffect(() => {
        fetchCourses().then(data => setCourses(data));
    }, []);

    return (
        <div>
            <h2>Available Courses</h2>
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

export default CourseList;