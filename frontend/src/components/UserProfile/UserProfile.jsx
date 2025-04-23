import React, { useState } from 'react';
import { FaEdit } from 'react-icons/fa';

const UserProfile = () => {
  const [user, setUser] = useState({
    email: 'user@example.com',
    password: '********',
    pseudo: 'monPseudo'
  });

  const [editingField, setEditingField] = useState(null);
  const [tempValue, setTempValue] = useState('');

  const handleEdit = (field) => {
    setEditingField(field);
    setTempValue(user[field]);
  };

  const handleSave = () => {
    setUser(prev => ({ ...prev, [editingField]: tempValue }));
    setEditingField(null);
  };

  return (
    <div className="user-profile">
      {['email', 'password', 'pseudo'].map((field) => (
        <div className="field-row" key={field} style={{ marginBottom: '1rem' }}>
          <label style={{ marginRight: '1rem' }}>{field.charAt(0).toUpperCase() + field.slice(1)}:</label>
          {editingField === field ? (
            <>
              <input
                type={field === 'password' ? 'password' : 'text'}
                value={tempValue}
                onChange={(e) => setTempValue(e.target.value)}
              />
              <button onClick={handleSave}>✅</button>
            </>
          ) : (
            <>
              <span style={{ marginRight: '0.5rem' }}>{user[field]}</span>
              <FaEdit style={{ cursor: 'pointer' }} onClick={() => handleEdit(field)} />
            </>
          )}
        </div>
      ))}
    </div>
  );
};

export default UserProfile;
