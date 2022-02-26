import { Config, ConfigState } from "./configTypes";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export const initialState: ConfigState = {
  configs: [],
  formValues: null,
};

const configSlice = createSlice({
  name: "config",
  initialState,
  reducers: {
    initializeCreateConfigForm(state) {
      state.formValues = {
        _id: null,
        name: "",
        temp: 0,
        pressure: 0,
        grinding: 20,
      };
    },
    initializeUpdateConfigForm(state, action: PayloadAction<Config>) {
      state.formValues = action.payload;
    },
    setConfigs(state, action: PayloadAction<Config[]>) {
      state.configs = action.payload ?? [];
    },
  },
});

export const {
  setConfigs,
  initializeCreateConfigForm,
  initializeUpdateConfigForm,
} = configSlice.actions;

export default configSlice.reducer;
