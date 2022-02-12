import { createSelector } from "@reduxjs/toolkit";
import { RootState } from "../app/store";
import { Sensor } from "./sensorTypes";

export const selectTempSensor = createSelector(
  (state: RootState) => state.sensor.temperature,
  (temp: Sensor | null) => temp
);

export const selectPressureSensor = createSelector(
  (state: RootState) => state.sensor.pressure,
  (pressure: Sensor | null) => pressure
);
