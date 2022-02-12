import { AppDispatch } from "../app/store";
import { Sensor } from "./sensorTypes";
import { setPressureSensor, setTemperatureSensor } from "./sensorReducer";

export const sensorWs = () => (dispatch: AppDispatch) => {
  const url = "ws://localhost:8080/api/sensors";
  const c = new WebSocket(url);

  const send = (data: string) => {
    c.send(data);
  };

  c.onmessage = function (msg) {
    const response: Sensor = JSON.parse(msg.data);
    if (response.sensorType === "temperature") {
      dispatch(setTemperatureSensor(response));
    } else {
      dispatch(setPressureSensor(response));
    }
  };

  c.onopen = function () {};
};
