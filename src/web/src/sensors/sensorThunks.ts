import { AppDispatch } from "../app/store";
import { Sensor } from "./sensorTypes";
import { setSensor } from "./sensorReducer";

export const sensorWs = () => (dispatch: AppDispatch) => {
  const url = "ws://localhost:8080/api/sensors";
  const c = new WebSocket(url);

  const send = (data: string) => {
    c.send(data);
  };

  c.onmessage = function (msg) {
    const response: Sensor = JSON.parse(msg.data);
    dispatch(setSensor(response));
  };

  c.onopen = function () {};

  return c.close;
};
