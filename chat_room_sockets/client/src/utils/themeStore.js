import { create } from "zustand";
import { persist } from "zustand/middleware";

const useThemeStore = create(
  persist(
    (set) => ({
      theme: "hellokitty", // 'hellokitty' or 'batman'
      toggleTheme: () =>
        set((state) => ({
          theme: state.theme === "hellokitty" ? "batman" : "hellokitty",
        })),
      setTheme: (theme) => set({ theme }),
    }),
    {
      name: "app-theme-store",
    }
  )
);

export default useThemeStore;
