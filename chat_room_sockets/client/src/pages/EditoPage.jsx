import { useState, useEffect, useRef } from "react";
import Client from "../components/Client";
import Editor from "../components/Editor";
import { Button } from "@mui/material";
import { DocumentTextIcon } from "@heroicons/react/24/solid";
import { useNavigate } from "react-router-dom";
import toast from "react-hot-toast";
import useDataStore from "../utils/store";
import ThemeToggle from "../components/ThemeToggle";
import { initSocket } from "../lib/socket";
import { getMessages } from "../lib/ApiService";

const EditoPage = () => {
  const { username } = useDataStore();
  const socketRef = useRef(null);
  const [connectionStatus, setConnectionStatus] = useState("Connecting...");
  const [messages, setMessages] = useState([]);
  const [loadingMessages, setLoadingMessages] = useState(false);
  const navigate = useNavigate();
  const [clients, setClients] = useState([]);

  useEffect(() => {
    const init = async () => {
      if (socketRef.current) return;

      try {
        socketRef.current = await initSocket();

        setConnectionStatus("Connected 💖");
        toast.success("Connected to the room 🎀");
        
        // Handle incoming messages
        socketRef.current.onmessage = (event) => {
          try {
            const serverMsg = JSON.parse(event.data);
            // server uses lowercase keys: { type: <int>, body: "..." }
            const rawBody = serverMsg.body;
            const tryParse = (v) => {
              try {
                return JSON.parse(v);
              } catch (e) {
                return null;
              }
            };

            const normalize = (m) => {
              if (!m) return { username: "", body: "" };
              if (typeof m === "string") {
                const p = tryParse(m);
                if (p && p.type === "message")
                  return { username: p.username || "", body: p.body || "" };
                return { username: "", body: m };
              }
              if (typeof m === "object") {
                if (m.body && typeof m.body === "string") {
                  const p = tryParse(m.body);
                  if (p && p.type === "message")
                    return {
                      username: p.username || m.username || "",
                      body: p.body || "",
                    };
                }
                if (m.type === "message" && (m.body || m.username))
                  return { username: m.username || "", body: m.body || "" };
                return {
                  username: m.username || "",
                  body: String(m.body || ""),
                };
              }
              return { username: "", body: String(m) };
            };

            const inner = tryParse(rawBody) || serverMsg;
            if (inner && inner.type === "message") {
              const nm = normalize(inner);
              setMessages((prev) => [...prev, nm]);
            } else if (inner && inner.type === "join") {
              toast.success(`${inner.username || "A user"} joined 🐱`);
            } else {
              const nm = normalize(rawBody);
              if (nm && nm.body) setMessages((prev) => [...prev, nm]);
            }
          } catch (e) {
            console.log("Non-JSON message from server:", event.data);
          }
        };

        socketRef.current.onerror = (error) => {
          console.error("WebSocket error:", error);
          toast.error("Connection error 💔");
          setConnectionStatus("Disconnected 💔");
        };

        socketRef.current.onclose = () => {
          setConnectionStatus("Disconnected 💔");
          console.log("WebSocket closed");
        };
      } catch (error) {
        console.error("Connection failed:", error);
        toast.error("Could not connect to the room 💔");
        setConnectionStatus("Failed to connect 💔");
      }
    };

    init();

    return () => {
      if (socketRef.current) {
        socketRef.current.close();
      }
    };
  }, [username]);

  useEffect(() => {
    let mounted = true;
    const fetchMessages = async () => {
      setLoadingMessages(true);
      try {
        const data = await getMessages();
        if (!mounted) return;
        const tryParse = (v) => {
          try {
            return JSON.parse(v);
          } catch (e) {
            return null;
          }
        };

        const normalize = (m) => {
          if (!m) return { username: "", body: "" };
          if (typeof m === "string") {
            const p = tryParse(m);
            if (p && p.type === "message")
              return { username: p.username || "", body: p.body || "" };
            return { username: "", body: m };
          }
          if (typeof m === "object") {
            if (m.body && typeof m.body === "string") {
              const p = tryParse(m.body);
              if (p && p.type === "message")
                return {
                  username: p.username || m.username || "",
                  body: p.body || "",
                };
            }
            if (m.type === "message" && (m.body || m.username))
              return { username: m.username || "", body: m.body || "" };
            return { username: m.username || "", body: String(m.body || "") };
          }
          return { username: "", body: String(m) };
        };

        const normalized = Array.isArray(data) ? data.map(normalize) : [];
        setMessages((prev) => {
          const existing = Array.isArray(prev) ? prev : [];
          return [...existing, ...normalized.filter((d) => d && d.body)];
        });
      } catch (err) {
        console.error("Failed to fetch messages:", err);
        toast.error("Failed to load messages");
      } finally {
        if (mounted) setLoadingMessages(false);
      }
    };

    fetchMessages();
    return () => {
      mounted = false;
    };
  }, []);

  const handleLeaveRoom = () => {
    setClients([]);
    navigate("/");
  };

  return (
    <div
      className="flex h-screen"
      style={{
        backgroundColor: "var(--bg-secondary)",
        transition: "background-color 0.3s ease",
      }}
    >
      {/* Sidebar */}
      <div
        className="w-72 backdrop-blur-xl shadow-2xl rounded-r-3xl border-r p-6 flex flex-col"
        style={{
          backgroundColor: "var(--card-bg)",
          borderColor: "var(--border-color)",
        }}
      >
        <div className="flex items-center gap-3 mb-8">
          <DocumentTextIcon
            className="w-8 h-8"
            style={{ color: "var(--text-secondary)" }}
          />
          <h1
            className="text-2xl font-extrabold tracking-tight"
            style={{ color: "var(--text-secondary)" }}
          >
            Hiee
          </h1>
          <div className="ml-auto">
            <ThemeToggle />
          </div>
        </div>

        <div
          className="mb-6 text-lg font-semibold"
          style={{ color: "var(--text-secondary)" }}
        >
          🐱 {username}
        </div>

        <div
          className="mb-6 text-sm"
          style={{ color: "var(--text-secondary)" }}
        >
          Status: <span className="text-green-500">{connectionStatus}</span>
        </div>

        <h3
          className="text-sm font-medium mb-3"
          style={{ color: "var(--text-secondary)" }}
        >
          Connected Users
        </h3>

        <div className="flex-1 space-y-2 overflow-auto pr-1">
          {clients.map((client) => (
            <Client key={client.socketId} username={client.username} />
          ))}
        </div>

        <div className="space-y-4 mt-6">
          <button
            onClick={handleLeaveRoom}
            className="w-full py-3 rounded-xl font-semibold shadow-lg hover:scale-105 transition"
            style={{
              backgroundColor: "var(--accent-color)",
              color:
                document.documentElement.getAttribute("data-theme") === "batman"
                  ? "#1a1a2e"
                  : "#fff",
            }}
          >
            💔 Leave Room
          </button>
        </div>
      </div>

      {/* Editor Area */}
      <div
        className="flex-1 backdrop-blur-xl rounded-l-3xl shadow-inner overflow-hidden"
        style={{
          backgroundColor: "var(--bg-primary)",
          transition: "background-color 0.3s ease",
        }}
      >
        <div className="relative h-full">
          {loadingMessages && (
            <div className="absolute inset-0 z-20 flex items-center justify-center">
              <div
                className="px-4 py-2 rounded"
                style={{ backgroundColor: "rgba(0,0,0,0.6)", color: "#fff" }}
              >
                Loading messages...
              </div>
            </div>
          )}

          <Editor
            socket={socketRef}
            username={username}
            messages={messages}
            onSend={(body) => {
              if (!socketRef.current) return;
              const payload = JSON.stringify({
                type: "message",
                username,
                body,
              });
              try {
                socketRef.current.send(payload);
              } catch (e) {
                console.error("Send failed", e);
              }
            }}
          />
        </div>
      </div>
    </div>
  );
};

export default EditoPage;
