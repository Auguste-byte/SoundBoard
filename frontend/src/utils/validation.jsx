
export const validateLogin = (identifier, password) => {
    if (!identifier || identifier.length < 3) {
      return "Veuillez saisir un email ou pseudo valide.";
    }
  
    if (!password || password.length < 6) {
      return "Mot de passe trop court (min. 6 caractères).";
    }
  
    return null;
  };
  