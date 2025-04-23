import React, { useState } from "react";

export default function LoginForm() {
  const [identifier, setIdentifier] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    // 👉 Appelle ton API ici (ex: fetch ou axios)
    console.log("Tentative de connexion :", identifier, password);
  };

  return (
    <div className="login-container">
      <h2 className="login-title">Connexion</h2>
      <form onSubmit={handleSubmit} className="login-form">
        <div className="form-group">
          <label htmlFor="identifier">Email ou pseudo</label>
          <input
            type="text"
            id="identifier"
            value={identifier}
            onChange={(e) => setIdentifier(e.target.value)}
            className="form-input"
            placeholder="ex: moi@mail.com ou pseudonyme"
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
            placeholder="Mot de passe"
            required
          />
        </div>

        <button type="submit" className="login-button">
          Se connecter
        </button>
      </form>
    </div>
  );
}
