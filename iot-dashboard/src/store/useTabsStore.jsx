import { create } from "zustand";

const useTabsStore = create((set) => ({
  selectedTab: 1,
  changeTab: (tab) =>
    set((state) => ({
      selectedTab: tab,
    })),
}));

export default useTabsStore;
