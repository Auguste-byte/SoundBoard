import Header from "../components/Header/Header";


const MainLayout = ({children}) => (
  <div>
      <Header />
      <main>{children}</main>
  </div>
);

export default MainLayout;