import React from "react";
import { useNavigate } from "react-router-dom";
import Header from "../../components/Header/Header";
import PostList from "../../components/PostList/PostList";

const HomePage = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    // Supprime le token
    localStorage.removeItem("token");

    // Redirige vers la page de login
    navigate("/");
  };

  return (
    <div>
      <Header />
      <h1>Bienvenue sur la page d'accueil</h1>
      <button onClick={handleLogout}>Se déconnecter</button>
      <PostList />
    </div>
  );
};

export default HomePage;
