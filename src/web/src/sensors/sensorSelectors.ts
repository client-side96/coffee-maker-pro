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

export const selectVolumeSensor = createSelector(
  (state: RootState) => state.sensor.volume,
  (volume: Sensor | null) => volume
);
export const selectTimeSensor = createSelector(
  (state: RootState) => state.sensor.time,
  (time: Sensor | null) => time
);
export const selectGrindingSensor = createSelector(
  (state: RootState) => state.sensor.grinding,
  (grinding: Sensor | null) => grinding
);
