import { MdDarkMode, MdLightMode } from "react-icons/md";
import useThemeStore from "../utils/themeStore";

function ThemeToggle() {
  const theme = useThemeStore((state) => state.theme);
  const toggleTheme = useThemeStore((state) => state.toggleTheme);

  return (
    <button
      onClick={toggleTheme}
      className="p-2 rounded-full hover:scale-110 transition-transform duration-200"
      style={{
        backgroundColor: theme === "hellokitty" ? "#ec4899" : "#ffd60a",
        color: theme === "hellokitty" ? "white" : "#1a1a2e",
      }}
      title={
        theme === "hellokitty" ? "Switch to Batman" : "Switch to Hello Kitty"
      }
    >
      {theme === "hellokitty" ? (
        <MdDarkMode size={20} />
      ) : (
        <MdLightMode size={20} />
      )}
    </button>
  );
}

export default ThemeToggle;
