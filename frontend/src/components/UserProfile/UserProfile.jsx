import React, { useState, useEffect } from 'react';
import { FaEdit } from 'react-icons/fa';

const UserProfile = () => {
  const [user, setUser] = useState({ email: '', pseudo: '' });
  const [editingField, setEditingField] = useState(null);
  const [tempValue, setTempValue] = useState('');
  const [message, setMessage] = useState('');

  const token = localStorage.getItem('token');

  // Charger les infos utilisateur depuis l'API
  useEffect(() => {
    const fetchProfile = async () => {
      try {
        const res = await fetch('/api/profile', {
          method: 'GET',
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        });

        const data = await res.json();
        if (res.ok) {
          setUser({ email: data.email, pseudo: data.pseudo });
        } else {
          console.error('Erreur :', data.message || 'Impossible de charger le profil');
        }
      } catch (err) {
        console.error('Erreur réseau :', err);
      }
    };

    if (token) fetchProfile();
  }, [token]);

  const handleEdit = (field) => {
    setEditingField(field);
    setTempValue(user[field]);
    setMessage('');
  };

  const handleSave = async () => {
    const updatedUser = { ...user, [editingField]: tempValue };

    try {
      const res = await fetch('/api/profile', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(updatedUser),
        
      });

      const data = await res.json();
      console.log("Token envoyé :", token);
      console.log("Réponse du backend :", data);

      if (res.ok) {
        setUser(updatedUser);
        setEditingField(null);
        setMessage('Profil mis à jour avec succès ✅');
      } else {
        setMessage(data.message || 'Erreur lors de la mise à jour.');
      }
    } catch (err) {
      setMessage('Erreur de connexion au serveur.');
    }
  };

  return (
    <div className="user-profile">
      <h2>Mon Profil</h2>
      {message && <p style={{ color: 'green' }}>{message}</p>}
  
      {['email', 'pseudo'].map((field) => (
        <div className="field-row" key={field} style={{ marginBottom: '1rem' }}>
          <label style={{ marginRight: '1rem' }}>
            {field.charAt(0).toUpperCase() + field.slice(1)} :
          </label>
          {editingField === field ? (
            <>
              <input
                type="text"
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
}  
export default UserProfile;
