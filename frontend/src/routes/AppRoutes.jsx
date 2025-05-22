import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import MainLayout from "../layout/MainLayout";
import Login from "../pages/Login/login";
import Register from "../pages/Register/register";
import HomePage from "../pages/home/homePage";
import ProtectedRoute from "../utils/ProtectedRoutes";
import UserProfile from "../components/UserProfile/UserProfile";

const AppRoutes = () => (
  <Router>
    <MainLayout>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/register" element={<Register />} />

        
        {/* ✅ Route protégée */}
        <Route
          path="/home"
          element={
            <ProtectedRoute>
              <HomePage />
            </ProtectedRoute>
          }
        />

        <Route
          path="/profile"
          element={
            <ProtectedRoute>
              <UserProfile />
            </ProtectedRoute>
          }
        />
        
        {/* Rediriger les chemins inconnus vers la page de login */}
        <Route path="*" element={<Login />} />
      </Routes>
    </MainLayout>
  </Router>
);

export default AppRoutes;
