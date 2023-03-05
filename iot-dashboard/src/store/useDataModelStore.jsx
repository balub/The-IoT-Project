import { create } from "zustand";

const useDataModelStore = create((set) => ({
  models: [
    {
      id: 0,
      title: "",
      type: "",
      required: true,
      component: "",
    },
  ],
  addElementToArray: (data) =>
    set((state) => ({
      models: [...state.models, { ...data, id: state.models.length + 1 }],
    })),
  updateModelState: (data) => set({ models: data }),
}));

export default useDataModelStore;
