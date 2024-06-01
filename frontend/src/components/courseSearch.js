import React, { useState } from 'react';
import { searchCourses } from '../services/api';

const CourseSearch = ({ setCourses }) => {
    const [query, setQuery] = useState('');

    const handleSearch = (e) => {
        e.preventDefault();
        searchCourses(query).then(data => setCourses(data));
    };

    return (
        <form onSubmit={handleSearch}>
            <input 
                type="text" 
                value={query} 
                onChange={(e) => setQuery(e.target.value)} 
                placeholder="Search courses..."
            />
            <button type="submit">Search</button>
        </form>
    );
};

export default CourseSearch;