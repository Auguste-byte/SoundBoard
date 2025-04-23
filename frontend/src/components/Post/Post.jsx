import React from 'react';
import AudioPlayer from '../AudioPlayer/AudioPlayer';

const PostItem = ({ title, description, date, children }) => {
  return (
    <div className="post-item" style={{ border: '1px solid #ccc', padding: '1rem', marginBottom: '1rem' }}>
      <h2 style={{ margin: '0 0 0.5rem' }}>{title}</h2>
      <p style={{ margin: '0 0 0.5rem' }}>{description}</p>
      <small style={{ color: '#777' }}>{new Date(date).toLocaleDateString()}</small>

      <div className="post-extra" style={{ marginTop: '1rem' }}>
        {children}
      </div>
      <AudioPlayer />
    </div>
  );
};

export default PostItem;
