const API_URL = 'http://localhost:3000/api';

export const fetchCourses = async () => {
    const response = await fetch(`${API_URL}/courses`);
    return response.json();
};

export const searchCourses = async (query) => {
    const response = await fetch(`${API_URL}/courses/search?query=${query}`);
    return response.json();
};

export const fetchCourseDetails = async (courseId) => {
    const response = await fetch(`${API_URL}/courses/${courseId}`);
    return response.json();
};

export const enrollInCourse = async (courseId) => {
    const response = await fetch(`${API_URL}/courses/${courseId}/enroll`, { method: 'POST' });
    return response.json();
};

export const fetchMyCourses = async () => {
    const response = await fetch(`${API_URL}/my-courses`);
    return response.json();
};