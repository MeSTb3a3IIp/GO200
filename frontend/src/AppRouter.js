import {Routes, Route} from "react-router-dom";
export default function AppRouter() {
    <Routes>
        <Route path={"/auth"} element={<Auth />}></Route>
        <Route path={"/"} element={<Main />}></Route>
        <Route path={"/tasks"} element={<Tasks />}></Route>
        <Route path={"/task/[id]"} element={<Task />}></Route>
        <Route path={"/"} element={<x />}></Route>
        <Route path={"/"} element={<x />}></Route>
        <Route path={"/"} element={<x />}></Route>
        <Route path={"/"} element={<x />}></Route>
        <Route path={"/"} element={<x />}></Route>
        <Route path={"/"} element={<x />}></Route>
        <Route path={"/profile/[id]"} element={<profile />}></Route>
        <Route path={"/info"} element={<Info />}></Route>
    </Routes>
}