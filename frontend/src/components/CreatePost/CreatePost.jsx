import React, { useState, useEffect } from "react";

function CreatePost({ onPostCreated }) {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');
  const [styleId, setStyleId] = useState('');
  const [styles, setStyles] = useState([]); // Liste dynamique
  const [audioFile, setAudioFile] = useState(null);
  const [error, setError] = useState('');
  const [message, setMessage] = useState('');

  useEffect(() => {
    const fetchStyles = async () => {
      try {
        const res = await fetch("/api/style");
        const data = await res.json();
        setStyles(data);
      } catch (err) {
        console.error("Erreur chargement styles :", err);
        setError("Impossible de charger les styles");
      }
    };

    fetchStyles();
  }, []);

  const handleAudioChange = (e) => {
    const file = e.target.files[0];
    setAudioFile(file);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setMessage('');

    if (!title || !styleId || !audioFile) {
      setError("Titre, style et fichier audio sont obligatoires");
      return;
    }

    const formData = new FormData();
    formData.append("title", title);
    formData.append("content", content);
    formData.append("style_id", styleId);
    formData.append("audio", audioFile);

    try {
      const token = localStorage.getItem("token");
      const res = await fetch("/api/posts", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token}`,
        },
        body: formData,
      });

      const data = await res.json();

      if (res.ok) {
        setMessage("Post publié ✅");
        setTitle('');
        setContent('');
        setStyleId('');
        setAudioFile(null);
        if (onPostCreated) onPostCreated(data);
      } else {
        setError(data.message || "Erreur lors de la publication");
      }
    } catch (err) {
      setError("Erreur réseau");
    }
  };

  return (
    <form onSubmit={handleSubmit} encType="multipart/form-data">
      <div>
        <label>Titre *</label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
      </div>

      <div>
        <label>Contenu du post</label>
        <textarea
          value={content}
          onChange={(e) => setContent(e.target.value)}
          placeholder="Ajoute ton message ici (optionnel)"
        />
      </div>

      <div>
        <label>Style musical *</label>
        <select
          value={styleId}
          onChange={(e) => setStyleId(e.target.value)}
          required
        >
          <option value="">-- Choisir un style --</option>
          {styles.map((style) => (
            <option key={style.id} value={style.id}>
              {style.name}
            </option>
          ))}
        </select>
      </div>

      <div>
        <label>Fichier audio *</label>
        <input
          type="file"
          accept="audio/*"
          onChange={handleAudioChange}
          required
        />
      </div>

      {error && <p style={{ color: "red" }}>{error}</p>}
      {message && <p style={{ color: "green" }}>{message}</p>}

      <button type="submit">Publier</button>
    </form>
  );
}

export default CreatePost;
