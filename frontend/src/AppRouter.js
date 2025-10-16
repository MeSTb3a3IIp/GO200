import {Routes, Route} from "react-router-dom";
import Main from './pages/Main.jsx';
import Auth from './pages/Auth.jsx';
import Info from './pages/Info.jsx';
import Tasks from './pages/Tasks.jsx';
import Task from './pages/Task.jsx';
import Profile from './pages/Profile.jsx';
import Course from './pages/Course.jsx';
import Creater from './pages/Creater.jsx';
import Comment from './pages/Comment.jsx';

export default function AppRouter() {
    return (
        <Routes>
            <Route path={"/auth"} element={<Auth />}></Route>
            <Route path={"/"} element={<Main />}></Route>
            <Route path={"/tasks"} element={<Tasks />}></Route>
            <Route path={"/task/:id"} element={<Task />}></Route>
            {/*<Route path={"/news"} element={<News/>}></Route>*/}
            <Route path={"/course"} element={<Course />}></Route>
            <Route path={"/creater"} element={<Creater />}></Route>
            <Route path={"/comment"} element={<Comment />}></Route>
            <Route path={"/profile/:id"} element={<Profile />}></Route>
            <Route path={"/info"} element={<Info />}></Route>
        </Routes>
    );
}