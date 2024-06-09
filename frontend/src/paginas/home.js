const courses = [
    { id: 1, title: "Curso de HTML", instructor: "Juan Pérez", duration: "4 semanas", requirements: "Ninguno", description: "Aprende HTML desde cero." },
    { id: 2, title: "Curso de JavaScript", instructor: "María Gómez", duration: "6 semanas", requirements: "HTML básico", description: "Aprende JavaScript y desarrolla interactividad en la web." },
    { id: 3, title: "Curso de C++", instructor: "Carlos Sánchez", duration: "8 semanas", requirements: "Programación básica", description: "Aprende C++ y sus aplicaciones." },
    { id: 4, title: "Curso de GoLang", instructor: "Ana López", duration: "5 semanas", requirements: "Programación intermedia", description: "Aprende GoLang para el desarrollo de software eficiente." },
    { id: 5, title: "Curso de Python", instructor: "Pedro Jiménez", duration: "7 semanas", requirements: "Ninguno", description: "Aprende Python y sus múltiples aplicaciones." },
];

let enrolledCourses = [];

document.addEventListener('DOMContentLoaded', () => {
    loadCourses();
    loadMyCourses();
});

function loadCourses() {
    const courseList = document.getElementById('course-list');
    courseList.innerHTML = '';
    courses.forEach(course => {
        const courseItem = document.createElement('div');
        courseItem.className = 'course-item';
        courseItem.innerHTML = `
            <h3>${course.title}</h3>
            <p>Instructor: ${course.instructor}</p>
            <button onclick="viewCourseDetails(${course.id})">Ver Detalles</button>
        `;
        courseList.appendChild(courseItem);
    });
}

function searchCourses() {
    const searchInput = document.getElementById('search-input').value.toLowerCase();
    const searchResults = document.getElementById('search-results');
    searchResults.innerHTML = '';

    const filteredCourses = courses.filter(course =>
        course.title.toLowerCase().includes(searchInput) ||
        course.description.toLowerCase().includes(searchInput)
    );

    filteredCourses.forEach(course => {
        const courseItem = document.createElement('div');
        courseItem.className = 'course-item';
        courseItem.innerHTML = `
            <h3>${course.title}</h3>
            <p>Instructor: ${course.instructor}</p>
            <button onclick="viewCourseDetails(${course.id})">Ver Detalles</button>
        `;
        searchResults.appendChild(courseItem);
    });
}

function viewCourseDetails(courseId) {
    const course = courses.find(course => course.id === courseId);
    const courseDetails = document.getElementById('course-details');
    const courseInfo = document.getElementById('course-info');

    courseInfo.innerHTML = `
        <h3>${course.title}</h3>
        <p>Descripción: ${course.description}</p>
        <p>Instructor: ${course.instructor}</p>
        <p>Duración: ${course.duration}</p>
        <p>Requisitos: ${course.requirements}</p>
    `;
    document.getElementById('enroll-button').setAttribute('onclick', `enrollCourse(${course.id})`);
    courseDetails.style.display = 'block';
}

function enrollCourse(courseId) {
    const course = courses.find(course => course.id === courseId);
    if (!enrolledCourses.includes(course)) {
        enrolledCourses.push(course);
        alert(`Te has inscrito en el curso: ${course.title}`);
        loadMyCourses();
    } else {
        alert('Ya estás inscrito en este curso.');
    }
}

function loadMyCourses() {
    const myCourseList = document.getElementById('my-course-list');
    myCourseList.innerHTML = '';
    enrolledCourses.forEach(course => {
        const courseItem = document.createElement('div');
        courseItem.className = 'course-item';
        courseItem.innerHTML = `
            <h3>${course.title}</h3>
            <button onclick="viewCourseDetails(${course.id})">Ver Detalles</button>
        `;
        myCourseList.appendChild(courseItem);
    });
}