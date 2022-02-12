import React from "react";
import { useAppDispatch, useAppSelector } from "./app/hooks";
import { sensorWs } from "./sensors/sensorThunks";
import {
  selectPressureSensor,
  selectTempSensor,
} from "./sensors/sensorSelectors";
import { Sensor } from "./sensors/sensorTypes";
import SensorTile from "./sensors/components/SensorTile";

function App() {
  const dispatch = useAppDispatch();
  const currentTemp = useAppSelector(selectTempSensor);
  const currentPressure = useAppSelector(selectPressureSensor);

  React.useEffect(() => {
    dispatch(sensorWs());
  }, []);

  return (
    <div style={{ display: "flex", flexDirection: "column" }}>
      <SensorTile title="Temperature" sensorData={currentTemp} unit="Celsius" />
      <SensorTile title="Pressure" sensorData={currentPressure} unit="HPa" />
    </div>
  );
}

export default App;
