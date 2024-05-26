import React, { useState } from 'react';
import CourseList from '../components/courseList';
import CourseSearch from '../components/courseSearch';

const Home = () => {
    const [courses, setCourses] = useState([]);

    return (
        <div>
            <CourseSearch setCourses={setCourses} />
            <CourseList courses={courses} />
        </div>
    );
};

export default Home;
