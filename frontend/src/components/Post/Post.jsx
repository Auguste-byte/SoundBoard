import React, { useState } from "react";
import AudioPlayer from "../AudioPlayer/AudioPlayer"; // ajuste le chemin si besoin

const Post = ({ postId, author, date, title, description, audioSrc }) => {
  const [likes, setLikes] = useState(0); // compteur local (tu peux le remplacer par une prop si tu veux charger le nombre initial)

  const handleLike = async () => {
    try {
      const token = localStorage.getItem("token");
      const res = await fetch("/api/posts/like", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ post_id: postId }),
      });

      const data = await res.json();

      if (res.ok) {
        setLikes((prev) => prev + 1);
      } else {
        console.error("Erreur Like:", data.error || data.message);
      }
    } catch (err) {
      console.error("Erreur réseau Like", err);
    }
  };

  const handleComment = () => {
    // futur : ouvrir un champ de commentaire
    console.log("💬 Commenter ce post", postId);
  };

  return (
    <div className="post">
      <div className="post-header">
        <p>
          <strong>{author}</strong> – {new Date(date).toLocaleDateString()}
        </p>
      </div>

      <div className="post-body">
        <h3>{title}</h3>
        {description && <p>{description}</p>}
        <AudioPlayer src={audioSrc} />
      </div>

      <div className="post-actions">
        <button onClick={handleLike}>👍 {likes}</button>
        <button onClick={handleComment}>💬 Commenter</button>
      </div>
    </div>
  );
};

export default Post;
