import React from "react";
import { Sensor, SensorType } from "../sensorTypes";
import "./SensorTile.scss";
import TemperatureIcon from "../../icons/TemperatureIcon";
import PressureIcon from "../../icons/PressureIcon";

interface SensorTileProps {
  title: string;
  sensorData: Sensor | null;
  unit: string;
}

const SensorIcon: React.FC<{ type?: SensorType }> = ({ type }) => {
  switch (type) {
    case "temperature":
      return <TemperatureIcon className="sensor-icon" />;
    case "pressure":
      return <PressureIcon className="sensor-icon" />;
    default:
      return <></>;
  }
};

const SensorTile: React.FC<SensorTileProps> = ({ title, sensorData, unit }) => {
  return (
    <div className="sensor-tile">
      <div className="sensor-header">
        <h5>{title}</h5>
        <SensorIcon type={sensorData?.sensorType} />
      </div>
      <p>
        {sensorData ? sensorData.value : "-"} {unit}
      </p>
    </div>
  );
};

export default SensorTile;
