import React, { useState } from "react";
import { loginUser } from "../../api/auth";
import { validateLogin } from "../../utils/validation";

export default function LoginForm() {
  const [identifier, setIdentifier] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [errorMsg, setErrorMsg] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    setErrorMsg(""); // Reset des erreurs

    const error = validateLogin(identifier, password);
    if (error) {
      setErrorMsg(error);
      return;
    }

    setLoading(true);
    try {
      const user = await loginUser(identifier, password);
      console.log("Connexion réussie :", user);
      // Tu peux rediriger ici avec useNavigate ou autre
    } catch (err) {
      setErrorMsg(
        err.response?.data?.error || "Erreur serveur ou identifiants invalides."
      );
    } finally {
      setLoading(false);
    }
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

        {/* Affichage du message d'erreur */}
        {errorMsg && (
          <p className="error-message" style={{ color: "red", marginTop: "10px" }}>
            {errorMsg}
          </p>
        )}

        <button type="submit" className="login-button" disabled={loading}>
          {loading ? "Connexion..." : "Se connecter"}
        </button>
      </form>
    </div>
  );
}
