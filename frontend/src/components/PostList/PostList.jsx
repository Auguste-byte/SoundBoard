import React, { useEffect, useState } from "react";
import Post from "../Post/Post";

const PostList = () => {
  const [posts, setPosts] = useState([]);
  const [error, setError] = useState("");

  // Chargement initial des posts
  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const res = await fetch("/api/posts");
        const data = await res.json();

        if (res.ok) {
          setPosts(data);
        } else {
          setError(data.message || data.error || "Erreur lors du chargement des posts");
        }
      } catch (err) {
        setError("Erreur réseau");
      }
    };

    fetchPosts();
  }, []);

  // WebSocket pour les nouveaux posts en temps réel
  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws"); // utilise `wss` si HTTPS

    socket.onmessage = (event) => {
      try {
        const newPost = JSON.parse(event.data);
        setPosts((prev) => [newPost, ...prev]);
      } catch (err) {
        console.error("Erreur WebSocket JSON :", err);
      }
    };

    return () => socket.close();
  }, []);

  return (
    <div>
      <h2>🎵 Derniers posts</h2>
      {error && <p style={{ color: "red" }}>{error}</p>}
      {posts.length === 0 && !error && <p>Aucun post pour le moment.</p>}

      {posts.map((post) => (
        <Post
          key={post.id}
          author={post.author}
          date={post.created_at}
          title={post.title}
          description={post.content}
          audioSrc={`/uploads/${post.music_file}`}
        />
      ))}
    </div>
  );
};

export default PostList;
