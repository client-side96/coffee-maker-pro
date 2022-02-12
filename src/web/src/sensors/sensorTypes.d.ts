export type SensorType = "temperature" | "pressure";

export interface Sensor {
  value: number;
  timestamp: Date;
  sensorType: SensorType;
}

export interface SensorState {
  temperature: Sensor | null;
  pressure: Sensor | null;
}
