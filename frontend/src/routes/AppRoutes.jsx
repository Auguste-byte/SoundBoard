import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MainLayout from "../layout/MainLayout";
import Login from "../pages/Login/login";
import RegisterForm from "../components/RegisterForm/RegisterForm";

const AppRoutes = () => (
  <Router>
    <MainLayout>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/register" element={<RegisterForm/>} />
      </Routes>
    </MainLayout>
  </Router>
);

export default AppRoutes;
