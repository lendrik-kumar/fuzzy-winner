const API_BASE =
  import.meta.env.REACT_APP_API_URL ||
  import.meta.env.VITE_NODE_URL ||
  (typeof window !== "undefined" ? window.location.origin : "");

const normalize = (u) => (u.endsWith("/") ? u.slice(0, -1) : u);

export async function getMessages() {
  const url = `${normalize(API_BASE)}/messages`;
  const res = await fetch(url, { method: "GET" });
  if (!res.ok) {
    throw new Error(`Request failed with status ${res.status}`);
  }
  const data = await res.json();
  return Array.isArray(data) ? data : [];
}

export default { getMessages };
