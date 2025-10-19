import { useParams } from "react-router-dom";
import "./StreamMovie.css";

const StreamMovie = () => {
  const { yt_id: key } = useParams();

  return (
    <div className="react-player-container">
      {key && (
        <iframe
          src={`https://www.youtube.com/embed/${key}`}
          title="YouTube video player"
          width="100%"
          height="100%"
          frameBorder="0"
          allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
          allowFullScreen
        ></iframe>
      )}
    </div>
  );
};

export default StreamMovie;
