````markdown
# MagicStream 🎥✨

A modern movie streaming platform with AI-powered recommendations, built using cutting-edge web technologies.

---

## Overview  

**MagicStream** is a full-stack web application that simulates a modern movie streaming platform. It combines a **React-based frontend** for a smooth, interactive user interface, a **Go backend using gin-gonic** for high-performance API services, and an **AI recommendation engine** powered by **LangChainGo** and **OpenAI** to provide personalized movie suggestions.  

The platform also leverages **MongoDB** as a scalable database to manage movies, metadata, and user preferences efficiently.

This project demonstrates how modern web technologies can work together to create a high-performance, intelligent, and user-friendly streaming experience.

---

## Key Features

- **Seamless Movie Streaming** – Simulated playback using **React and React-Player**.
- **High-Performance Backend** – Go-based APIs using **gin-gonic**.
- **AI Recommendations** – Personalized suggestions using **LangChainGo** and **OpenAI**.
- **Scalable Data Storage** – Efficiently stores movies and user data in **MongoDB**.
- **Clean & Modern UI** – Responsive and user-friendly interface with React.

---

## Tech Stack

| Layer             | Technology        |
|------------------ |----------------- |
| Frontend / Client | JavaScript / React |
| Backend / Server  | Go / gin-gonic    |
| Database / Storage| MongoDB           |
| AI / Recommendation | LangChainGo + OpenAI |

---

## Demo Video

Watch a step-by-step tutorial on building this app:  
[![MagicStream Video Tutorial](https://img.youtube.com/vi/jBf7of9JTV8/0.jpg)](https://youtu.be/jBf7of9JTV8)

---

## Installation & Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/Vijayakumar-Harish/magic-movie-stream.git
````

2. Install frontend dependencies:

   ```bash
   cd Client/magic-stream-client
   npm install
   npm run dev
   ```

3. Install backend dependencies and run Go server:

   ```bash
   cd Server/MagicStreamMoviesServer
   go mod tidy
   go run main.go
   ```

4. Ensure **MongoDB** is running locally or via Docker.

5. Open the app in your browser:

   ```
   http://localhost:5173
   ```

---

## Project Structure

```

├─ Client/               # React frontend
│  └─ magic-stream-client/
├─ Server/               # Go backend
│  └─ MagicStreamMoviesServer/
└─ README.md
```

---

## Contribution

Contributions are welcome! You can:

* Report issues or bugs
* Suggest improvements
* Submit pull requests

---

## License

This project is open source and available under the MIT License.

---

## Acknowledgements

* React
* Go / gin-gonic
* MongoDB
* OpenAI & LangChainGo
* IFRAME

```

