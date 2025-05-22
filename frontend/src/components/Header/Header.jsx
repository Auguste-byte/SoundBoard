import React from "react";
import { useNavigate } from "react-router-dom";
import "./Header.css";
import PostModalWrapper from "../PostModal/PostModal";

const Header = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/");
  };

  const goToProfile = () => {
    navigate("/profile");
  };

  return (
    <header className="header">
      <div className="header-left">
        <img src="/assets/logo.avif" alt="logo" className="logo" />
        <h1 className="appliName">SOUNDBOARD</h1>
      </div>
      <div className="header-right">
        <button className="header-btn" onClick={goToProfile}>Profil</button>
        <button className="header-btn" onClick={handleLogout}>Déconnexion</button>
        <PostModalWrapper />
      </div>
    </header>
  );
};

export default Header;
