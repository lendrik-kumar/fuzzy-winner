import { create } from "zustand";
import { persist, devtools } from "zustand/middleware";

const data = (set) => ({
  username: "",
  setUsername: (name) => set({ username: name }),
});

const useDataStore = create()(
  devtools(
    persist(data, {
      name: "live-doc-data",
    })
  )
);

export default useDataStore;
