import React from "react";
import { useAppDispatch, useAppSelector } from "./app/hooks";
import { sensorWs } from "./sensors/sensorThunks";
import {
  selectPressureSensor,
  selectTempSensor,
} from "./sensors/sensorSelectors";
import { Sensor } from "./sensors/sensorTypes";
import SensorTile from "./sensors/components/SensorTile";
import "./App.scss";
import { statusWs, turnOff, turnOn } from "./status/statusThunks";
import { selectStatus } from "./status/statusSelector";
import PowerButton from "./status/components/PowerButton";

function App() {
  const dispatch = useAppDispatch();
  const currentTemp = useAppSelector(selectTempSensor);
  const currentPressure = useAppSelector(selectPressureSensor);
  const status = useAppSelector(selectStatus);

  const isPowerOn = status.status !== "Off";

  React.useEffect(() => {
    const closeSensorConnection = dispatch(sensorWs());
    const closeStatusConnection = dispatch(statusWs());

    return () => {
      closeSensorConnection();
      closeStatusConnection();
    };
  }, []);

  return (
    <div className="container">
      <div className="header">
        <h4>Coffee Maker Pro</h4>
        <PowerButton isPowerON={isPowerOn} />
      </div>
      <div className="sensor-wrapper">
        <SensorTile
          title="Temperature"
          sensorData={currentTemp}
          unit="Celsius"
        />
        <SensorTile title="Pressure" sensorData={currentPressure} unit="bar" />
        <SensorTile title="Grinding" sensorData={currentPressure} unit="Step" />
      </div>
      <p>{status.status}</p>
    </div>
  );
}

export default App;
