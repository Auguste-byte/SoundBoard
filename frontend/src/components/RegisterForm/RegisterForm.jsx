import React, { useState } from "react";

export default function RegisterForm() {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    // 👉 Ici tu appelles ton API d'inscription
    console.log("Tentative d'inscription :", { email, username, password });
  };

  return (
    <div className="register-container">
      <h2 className="register-title">Créer un compte</h2>
      <form onSubmit={handleSubmit} className="register-form">
        <div className="form-group">
          <label htmlFor="email">Adresse email</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="form-input"
            placeholder="ex: utilisateur@mail.com"
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="username">Nom d'utilisateur</label>
          <input
            type="text"
            id="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="form-input"
            placeholder="ex: musicfan93"
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="password">Mot de passe</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="form-input"
            placeholder="••••••••"
            required
          />
        </div>

        <button type="submit" className="register-button">
          S'inscrire
        </button>
      </form>
    </div>
  );
}
