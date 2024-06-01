import React from 'react';
import { enrollInCourse } from '../services/api';

const EnrollButton = ({ courseId }) => {
    const handleEnroll = () => {
        enrollInCourse(courseId).then(response => {
            if (response.success) {
                alert('Enrolled successfully!');
            }
        });
    };

    return <button onClick={handleEnroll}>Enroll</button>;
};

export default EnrollButton;
