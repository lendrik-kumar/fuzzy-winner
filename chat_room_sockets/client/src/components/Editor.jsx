import React, { useState, useRef, useEffect } from "react";
import {
  PhoneIcon,
  VideoCameraIcon,
  FaceSmileIcon,
  PaperAirplaneIcon,
  MicrophoneIcon,
  XMarkIcon,
} from "@heroicons/react/24/solid";

const emojis = [
  "😀",
  "😂",
  "😍",
  "🥰",
  "😎",
  "😭",
  "🤩",
  "🥺",
  "😡",
  "😴",
  "🤯",
  "👍",
  "👏",
  "🙏",
  "💖",
  "🔥",
  "🎀",
  "🐱",
  "✨",
  "💬",
  "💐",
  "🌸",
  "🌈",
  "🍓",
  "🍩",
  "🍰",
];

const Editor = ({ socket, messages = [], onSend, username }) => {
  const [message, setMessage] = useState("");
  const [showEmojis, setShowEmojis] = useState(false);
  const messagesRef = useRef(null);

  useEffect(() => {
    if (messagesRef.current) {
      messagesRef.current.scrollTop = messagesRef.current.scrollHeight;
    }
  }, [messages]);

  const handleKeyDown = (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      const trimmed = message.trim();
      if (trimmed && onSend) {
        onSend(trimmed);
        setMessage("");
      }
    }
  };

  const themeIsBatman =
    document.documentElement.getAttribute("data-theme") === "batman";

  return (
    <div
      className="h-full flex flex-col"
      style={{
        backgroundColor: "var(--bg-primary)",
        transition: "background-color 0.2s ease",
      }}
    >
      <div
        className="flex items-center justify-between px-6 py-4"
        style={{
          borderBottomWidth: "1px",
          borderColor: "var(--border-color)",
          backgroundColor: "var(--card-bg)",
        }}
      >
        <div>
          <h2
            className="text-lg font-bold"
            style={{ color: "var(--text-secondary)" }}
          >
            Hiee
          </h2>
          <p
            className="text-sm"
            style={{ color: "var(--text-secondary)", opacity: 0.8 }}
          >
            Online • Chatting
          </p>
        </div>

        <div className="flex items-center gap-3">
          <button
            className="p-2 rounded-full transition"
            style={{ backgroundColor: "var(--accent-color)" }}
          >
            <PhoneIcon
              className="w-5 h-5"
              style={{ color: themeIsBatman ? "#1a1a2e" : "#fff" }}
            />
          </button>
          <button
            className="p-2 rounded-full transition"
            style={{ backgroundColor: "var(--accent-color)" }}
          >
            <VideoCameraIcon
              className="w-5 h-5"
              style={{ color: themeIsBatman ? "#1a1a2e" : "#fff" }}
            />
          </button>
        </div>
      </div>

      <div
        ref={messagesRef}
        className="flex-1 overflow-y-auto px-6 py-4 space-y-4"
      >
        {messages.map((m, i) => (
          <div
            key={i}
            className={`flex ${
              m.username === username ? "justify-end" : "justify-start"
            }`}
          >
            <div
              className="px-4 py-2 rounded-2xl max-w-xs shadow"
              style={{
                backgroundColor:
                  m.username === username
                    ? "var(--button-bg)"
                    : "var(--border-color)",
                color:
                  m.username === username
                    ? themeIsBatman
                      ? "#1a1a2e"
                      : "#fff"
                    : "var(--text-primary)",
              }}
            >
              <div className="text-xs opacity-80 mb-1">{m.username}</div>
              <div style={{ whiteSpace: "pre-wrap" }}>{m.body}</div>
            </div>
          </div>
        ))}
      </div>

      {showEmojis && (
        <div
          className="p-3 grid grid-cols-8 gap-2 text-xl"
          style={{
            backgroundColor: "var(--card-bg)",
            borderTopWidth: "1px",
            borderColor: "var(--border-color)",
          }}
        >
          {emojis.map((e, i) => (
            <button
              key={i}
              onClick={() => setMessage((prev) => prev + e)}
              className="hover:scale-110 transition"
            >
              {e}
            </button>
          ))}
        </div>
      )}

      <div
        className="p-4 flex items-center gap-3"
        style={{
          borderTopWidth: "1px",
          borderColor: "var(--border-color)",
          backgroundColor: "var(--card-bg)",
        }}
      >
        <button
          onClick={() => setShowEmojis((s) => !s)}
          className="p-2 rounded-full transition"
          style={{ backgroundColor: "var(--accent-color)" }}
        >
          {showEmojis ? (
            <XMarkIcon
              className="w-6 h-6"
              style={{ color: themeIsBatman ? "#1a1a2e" : "#fff" }}
            />
          ) : (
            <FaceSmileIcon
              className="w-6 h-6"
              style={{ color: themeIsBatman ? "#1a1a2e" : "#fff" }}
            />
          )}
        </button>

        <textarea
          rows={1}
          placeholder="Type a message..."
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={handleKeyDown}
          className="flex-1 px-4 py-2 rounded-full focus:outline-none resize-none"
          style={{
            borderWidth: "1px",
            borderColor: "var(--border-color)",
            backgroundColor: themeIsBatman ? "white" : "white" ,
            color: "var(--text-primary)",
            outlineColor: "var(--border-color)",
          }}
        />

        <button
          className="p-2 rounded-full transition"
          style={{ backgroundColor: "var(--accent-color)" }}
        >
          <MicrophoneIcon
            className="w-6 h-6"
            style={{ color: themeIsBatman ? "#1a1a2e" : "#fff" }}
          />
        </button>

        <button
          className="p-2 rounded-full transition"
          style={{ backgroundColor: "var(--button-bg)" }}
          onClick={() => {
            const trimmed = message.trim();
            if (trimmed && onSend) {
              onSend(trimmed);
              setMessage("");
            }
          }}
        >
          <PaperAirplaneIcon
            className="w-5 h-5 rotate-45"
            style={{ color: themeIsBatman ? "#1a1a2e" : "#fff" }}
          />
        </button>
      </div>
    </div>
  );
};

export default Editor;
