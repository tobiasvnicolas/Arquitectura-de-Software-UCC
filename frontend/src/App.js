import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Home from './pages/Home';
import CoursePage from './pages/CoursePage';
import MyCoursesPage from './pages/MyCoursesPage';

const App = () => {
    return (
        <Router>
            <Switch>
                <Route path="/" exact component={Home} />
                <Route path="/course/:courseId" component={CoursePage} />
                <Route path="/my-courses" component={MyCoursesPage} />
            </Switch>
        </Router>
    );
};

export default App;
