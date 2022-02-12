import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { SensorState, Sensor } from "./sensorTypes";

const initialState: SensorState = {
  temperature: null,
  pressure: null,
};

const sensorSlice = createSlice({
  name: "sensor",
  initialState,
  reducers: {
    setTemperatureSensor(state, action: PayloadAction<Sensor>) {
      state.temperature = action.payload;
    },
    setPressureSensor(state, action: PayloadAction<Sensor>) {
      state.pressure = action.payload;
    },
  },
});

export const { setPressureSensor, setTemperatureSensor } = sensorSlice.actions;
export default sensorSlice.reducer;
