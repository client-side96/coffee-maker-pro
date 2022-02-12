import React from "react";
import { Sensor } from "../sensorTypes";

interface SensorTileProps {
  title: string;
  sensorData: Sensor | null;
  unit: string;
}

const SensorTile: React.FC<SensorTileProps> = ({ title, sensorData, unit }) => {
  return (
    <p>
      {title}: {sensorData ? sensorData.value : "-"} {unit}
    </p>
  );
};

export default SensorTile;
