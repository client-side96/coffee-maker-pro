import { AppDispatch } from "../app/store";
import { Sensor } from "../sensors/sensorTypes";
import {
  setPressureSensor,
  setTemperatureSensor,
} from "../sensors/sensorReducer";
import { setStatus } from "./statusReducer";
import { Status, StatusResponse, StatusState } from "./statusTypes";

export const statusWs = () => (dispatch: AppDispatch) => {
  const url = "ws://localhost:8080/api/status";
  const c = new WebSocket(url);

  const send = (data: string) => {
    c.send(data);
  };

  c.onmessage = function (msg) {
    const response: StatusResponse = JSON.parse(msg.data);
    console.log(response);
    dispatch(setStatus(response));
  };

  c.onopen = function () {};

  return c.close;
};

export const turnOn = async () => {
  try {
    await fetch("http://localhost:8080/api/status/power/on", {
      method: "POST",
    });
  } catch (err) {
    console.error("Cannot turn on coffee maker", err);
    throw err;
  }
};

export const turnOff = async () => {
  try {
    await fetch("http://localhost:8080/api/status/power/off", {
      method: "POST",
    });
  } catch (err) {
    console.error("Cannot turn on coffee maker", err);
    throw err;
  }
};
