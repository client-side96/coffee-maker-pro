import React from "react";
import clsx from "clsx";
import { Sensor, SensorType } from "../sensorTypes";
import "./SensorTile.scss";
import TemperatureIcon from "../../icons/TemperatureIcon";
import PressureIcon from "../../icons/PressureIcon";
import { Status } from "../../status/statusTypes";

interface SensorTileProps {
  title: string;
  sensorData: Sensor | null;
  unit: string;
  currentStatus: Status;
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

const SensorTile: React.FC<SensorTileProps> = ({
  title,
  sensorData,
  unit,
  currentStatus,
}) => {
  const colorClassName = currentStatus.toLowerCase();
  return (
    <div className={clsx("sensor-tile", colorClassName)}>
      <div className="sensor-header">
        <h5>{title}</h5>
        <SensorIcon type={sensorData?.sensorType} />
      </div>
      <p>
        {sensorData !== undefined && sensorData !== null
          ? sensorData.value
          : "-"}{" "}
        {unit}
      </p>
    </div>
  );
};

export default SensorTile;
