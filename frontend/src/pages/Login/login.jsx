import LoginForm from "../../components/LoginForm/LoginForm";
import PostItem from "../../components/Post/Post";
import RegisterForm from "../../components/RegisterForm/RegisterForm";
import UserProfile from "../../components/UserProfile/UserProfile";
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