import React from "react";
import toast from "react-hot-toast";
import { useNavigate } from "react-router";
import useDataStore from "../utils/store";
import ThemeToggle from "../components/ThemeToggle";

const Home = () => {
  const { username, setUsername } = useDataStore();
  const navigate = useNavigate();
  const handleJoin = () => {
    if (!username) {
      toast.error("Please enter a username 💗");
      return;
    }
    navigate(`/editor`);
  };

  const handleKeyDown = (e) => {
    if (e.key === "Enter") handleJoin();
  };

  return (
    <div
      className="min-h-screen flex flex-col relative overflow-hidden"
      style={{
        backgroundColor: "var(--bg-secondary)",
        transition: "background-color 0.3s ease",
      }}
    >
      {/* Theme toggle button */}
      <div className="absolute top-4 right-4 z-10">
        <ThemeToggle />
      </div>

      {/* Floating background bubbles */}
      <div
        className="absolute -top-40 -right-40 w-96 h-96 rounded-full animate-float"
        style={{
          backgroundColor: "var(--border-color)",
          opacity: "0.4",
        }}
      ></div>
      <div
        className="absolute -bottom-40 -left-40 w-96 h-96 rounded-full animate-float"
        style={{
          backgroundColor: "var(--accent-color)",
          opacity: "0.4",
        }}
      ></div>

      <div className="flex-1 flex items-center justify-center px-4">
        <div
          className="backdrop-blur-xl p-8 rounded-3xl shadow-2xl transition w-full max-w-md"
          style={{
            backgroundColor: "var(--card-bg)",
            borderColor: "var(--border-color)",
            borderWidth: "2px",
          }}
        >
          {/* Sparkle header */}
          <div className="flex justify-center gap-3 text-3xl mb-2">
            <span className="animate-sparkle">✨</span>
            <span className="animate-sparkle delay-150">💖</span>
            <span className="animate-sparkle delay-300">✨</span>
          </div>

          <h4
            className="text-3xl font-bold text-center mb-6"
            style={{ color: "var(--text-secondary)" }}
          >
            💗 Join a Room
          </h4>

          <div className="space-y-4">
            <input
              type="text"
              placeholder="🐱 Username"
              className="w-full px-4 py-3 rounded-xl focus:outline-none focus:ring-2 transition"
              style={{
                borderWidth: "2px",
                borderColor: "var(--border-color)",
                backgroundColor:
                  document.documentElement.getAttribute("data-theme") ===
                  "white",
                color: "var(--text-primary)",
              }}
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              onKeyDown={handleKeyDown}
            />
          </div>

          <button
            onClick={handleJoin}
            className="mt-6 w-full py-3 rounded-xl font-semibold shadow-lg hover:scale-105 transition active:scale-95"
            style={{
              backgroundColor: "var(--button-bg)",
              color:
                document.documentElement.getAttribute("data-theme") === "white"
            }}
          >
            ✨ Join Room
          </button>
        </div>
      </div>
    </div>
  );
};

export default Home;
