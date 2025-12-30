const BASE_URL = import.meta.env.VITE_NODE_URL;

export const initSocket = async () => {
  return new Promise((resolve, reject) => {
    const wsUrl = BASE_URL.replace("http", "ws") + "/ws";

    const socket = new WebSocket(wsUrl);

    socket.onopen = () => {
      console.log("WebSocket connected");
      resolve(socket);
    };

    socket.onerror = (error) => {
      console.error("WebSocket error:", error);
      reject(error);
    };
  });
};
