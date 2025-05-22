import React, { useState } from 'react';
import CreatePost from '../CreatePost/CreatePost';

const PostModalWrapper = () => {
  const [isOpen, setIsOpen] = useState(false);

  const handlePostCreated = (data) => {
    console.log("Nouveau post créé :", data);
    setIsOpen(false);
  };

  return (
    <div>
      <button onClick={() => setIsOpen(true)}>Créer un post</button>

      {isOpen && (
        <div className="modal-overlay">
          <div className="modal-content">
            <button onClick={() => setIsOpen(false)}>✖️ Fermer</button>
            <CreatePost onPostCreated={handlePostCreated} />
          </div>
        </div>
      )}
    </div>
  );
};

export default PostModalWrapper;
