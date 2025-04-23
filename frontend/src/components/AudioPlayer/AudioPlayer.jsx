import { useRef, useState, useEffect } from "react";

export default function AudioPlayer({ title, src }) {
  const audioRef = useRef(null);
  const [isPlaying, setIsPlaying] = useState(false);
  const [currentTime, setCurrentTime] = useState(0);
  const [duration, setDuration] = useState(0);
  const [volume, setVolume] = useState(1); // 1 = volume max

  // Met à jour le temps écoulé et la durée totale de la musique
  useEffect(() => {
    const audio = audioRef.current;
    const updateTime = () => {
      setCurrentTime(audio.currentTime);
    };
    const setTotalDuration = () => {
      setDuration(audio.duration);
    };

    audio.addEventListener("timeupdate", updateTime);
    audio.addEventListener("loadedmetadata", setTotalDuration);

    return () => {
      audio.removeEventListener("timeupdate", updateTime);
      audio.removeEventListener("loadedmetadata", setTotalDuration);
    };
  }, []);

  // play/pause
  const togglePlay = () => {
    const audio = audioRef.current;
    if (isPlaying) {
      audio.pause();
    } else {
      audio.play();
    }
    setIsPlaying(!isPlaying);
  };

  // Gérer le changement de volume
  const handleVolumeChange = (e) => {
    const newVolume = e.target.value;
    setVolume(newVolume);
    audioRef.current.volume = newVolume;
  };

  // barre de progression du temps
  const handleTimeChange = (e) => {
    const newTime = e.target.value;
    setCurrentTime(newTime);
    audioRef.current.currentTime = newTime;
  };

  return (
    <div style={{ padding: "1rem", border: "1px solid #ccc", borderRadius: "10px", maxWidth: "400px" }}>
      <h3>{title}</h3>
      <button onClick={togglePlay}>
        {isPlaying ? "Pause" : "Play"}
      </button>
      <div style={{ margin: "1rem 0" }}>
        <label>Volume:</label>
        <input
          type="range"
          min="0"
          max="1"
          step="0.01"
          value={volume}
          onChange={handleVolumeChange}
          style={{ width: "100%" }}
        />
      </div>
      <div>
        <p>
          {Math.floor(currentTime)} sec / {Math.floor(duration)} sec
        </p>
        <input
          type="range"
          min="0"
          max={duration || 0}
          value={currentTime}
          onChange={handleTimeChange}
          style={{ width: "100%" }}
        />
      </div>
      <audio ref={audioRef} src={src} preload="auto" />
    </div>
  );
}
