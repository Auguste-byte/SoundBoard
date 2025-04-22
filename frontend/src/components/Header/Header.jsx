import "./Header.css"
import AudioPlayer from "../AudioPlayer/AudioPlayer";

const Header = () => {
    
    return (
    <header className="header">
        <div className="user">
        bonjours
        </div>
        <div style={{ padding: '2rem' }}>
      <AudioPlayer
        title="Ma super musique"
        src="/assets/test.mp3"
      />
    </div>
    </header>
)};

export default Header;