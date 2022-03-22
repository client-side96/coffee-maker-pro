import React from "react";
import { useAppDispatch, useAppSelector } from "./app/hooks";
import { sensorWs } from "./sensors/sensorThunks";
import {
  selectGrindingSensor,
  selectPressureSensor,
  selectTempSensor,
  selectTimeSensor,
  selectVolumeSensor,
} from "./sensors/sensorSelectors";
import SensorTile from "./sensors/components/SensorTile";
import "./App.scss";
import { statusWs } from "./status/statusThunks";
import { selectStatus } from "./status/statusSelector";
import PowerButton from "./status/components/PowerButton";
import ConfigList from "./config/components/ConfigList";
import ConfigForm from "./config/components/ConfigForm";
import {
  selectConfigFormValues,
  selectConfigs,
} from "./config/configSelectors";
import { getConfigs } from "./config/configThunks";
import { initializeCreateConfigForm } from "./config/configReducer";

function App() {
  const dispatch = useAppDispatch();
  const currentTemp = useAppSelector(selectTempSensor);
  const currentPressure = useAppSelector(selectPressureSensor);
  const currentVolume = useAppSelector(selectVolumeSensor);
  const currentTime = useAppSelector(selectTimeSensor);
  const currentGrinding = useAppSelector(selectGrindingSensor);
  const status = useAppSelector(selectStatus);
  const configs = useAppSelector(selectConfigs);
  const configFormValues = useAppSelector(selectConfigFormValues);

  const isPowerOn = status.status !== "Off";

  React.useEffect(() => {
    const closeSensorConnection = dispatch(sensorWs());
    const closeStatusConnection = dispatch(statusWs());
    dispatch(getConfigs());
    dispatch(initializeCreateConfigForm());

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
          title="Wassertemperatur"
          sensorData={currentTemp}
          unit="Celsius"
          currentStatus={status.status}
        />
        <SensorTile
          title="Wasserdruck"
          sensorData={currentPressure}
          unit="bar"
          currentStatus={status.status}
        />
        <SensorTile
          title="Mahlgrad"
          sensorData={currentGrinding}
          unit="Step"
          currentStatus={status.status}
        />
        <SensorTile
          title="Wasservolumen"
          sensorData={currentVolume}
          unit="ml"
          currentStatus={status.status}
        />
        <SensorTile
          title="Durchlaufzeit"
          sensorData={currentTime}
          unit="Sekunden"
          currentStatus={status.status}
        />
      </div>
      <div className="config-wrapper">
        <ConfigList configs={configs} />
        {configFormValues && <ConfigForm formValues={configFormValues} />}
      </div>
    </div>
  );
}

export default App;
