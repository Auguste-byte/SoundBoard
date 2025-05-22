import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { loginUser } from "../../api/auth";
import { validateLogin } from "../../utils/validation";

export default function LoginForm() {
  const navigate = useNavigate();

  const [identifier, setIdentifier] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [errorMsg, setErrorMsg] = useState("");

  // 🔁 Redirige vers /home si un token est déjà présent
  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      navigate("/home");
    }
  }, [navigate]);

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

      // ✅ Vérifie si un token est présent dans la réponse
      if (user && user.token) {
        localStorage.setItem("token", user.token);
        navigate("/home");
      } else {
        setErrorMsg("Token manquant dans la réponse du serveur.");
      }
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
