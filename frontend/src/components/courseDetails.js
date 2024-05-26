import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { fetchCourseDetails } from '../services/api';
import EnrollButton from './EnrollButton';

const CourseDetail = () => {
    const { courseId } = useParams();
    const [course, setCourse] = useState(null);

    useEffect(() => {
        fetchCourseDetails(courseId).then(data => setCourse(data));
    }, [courseId]);

    if (!course) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h2>{course.title}</h2>
            <p>{course.description}</p>
            <p>Instructor: {course.instructor}</p>
            <p>Duration: {course.duration}</p>
            <p>Requirements: {course.requirements}</p>
            <EnrollButton courseId={courseId} />
        </div>
    );
};

export default CourseDetail;
