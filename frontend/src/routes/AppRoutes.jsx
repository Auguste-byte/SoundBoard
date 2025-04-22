import { BrowserRouter as Router} from "react-router-dom";
import MainLayout from "../layout/MainLayout";
import Home from "../pages/login";







const AppRoutes = () => (
  <Router>
    <MainLayout>
      <Home />
    </MainLayout>
  </Router>
);

export default AppRoutes;