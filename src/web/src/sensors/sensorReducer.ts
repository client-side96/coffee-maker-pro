import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { SensorState, Sensor } from "./sensorTypes";

const initialState: SensorState = {
  temperature: null,
  pressure: null,
  volume: null,
  time: null,
  grinding: null,
};

const sensorSlice = createSlice({
  name: "sensor",
  initialState,
  reducers: {
    setSensor(state, action: PayloadAction<Sensor>) {
      state[action.payload.sensorType] = action.payload;
    },
  },
});

export const { setSensor } = sensorSlice.actions;
export default sensorSlice.reducer;
