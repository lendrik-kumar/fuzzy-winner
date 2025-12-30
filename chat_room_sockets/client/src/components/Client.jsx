import React from "react";
import Avatar from "react-avatar";

const Client = ({ username }) => {
  return (
    <div
      className="flex items-center gap-3 p-2 rounded-lg transition-colors"
      style={{
        backgroundColor: "var(--border-color)",
        opacity: "0.1",
      }}
    >
      <Avatar name={username} size={32} round="8px" textSizeRatio={2} />
      <span
        className="text-sm font-medium"
        style={{ color: "var(--text-primary)" }}
      >
        {username}
      </span>
    </div>
  );
};

export default Client;
