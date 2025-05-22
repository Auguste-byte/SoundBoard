import LoginForm from "../../components/LoginForm/LoginForm";
import { Link } from "react-router-dom";

import "./login.css";

const Login = () => {
  return (
    <div className="story-container">
      <LoginForm />
      <Link to="/register">
        <button>S'inscrire</button>
      </Link>
    </div>
  );
  
};

export default Login;