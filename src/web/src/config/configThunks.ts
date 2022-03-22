import { AppDispatch } from "../app/store";
import {
  initializeCreateConfigForm,
  initializeUpdateConfigForm,
  setConfigs,
} from "./configReducer";
import { Config } from "./configTypes";

export const getConfigs = () => async (dispatch: AppDispatch) => {
  try {
    const response = await fetch("http://localhost:8080/api/config");
    const payload = await response.json();
    dispatch(setConfigs(payload));
  } catch (err) {
    console.error("Cannot retrieve configs", err);
    throw err;
  }
};

export const getConfigById = (id: string) => async (dispatch: AppDispatch) => {
  try {
    const response = await fetch(`http://localhost:8080/api/config/${id}`);
    const payload = await response.json();
    dispatch(initializeUpdateConfigForm(payload));
  } catch (err) {
    console.error("Cannot retrieve config by id", err);
    throw err;
  }
};

export const createConfig =
  (newConfig: Config) => async (dispatch: AppDispatch) => {
    try {
      await fetch(`http://localhost:8080/api/config`, {
        method: "POST",
        body: JSON.stringify(newConfig),
      });
      dispatch(initializeCreateConfigForm());
      await dispatch(getConfigs());
    } catch (err) {
      console.error("Cannot create config", err);
      throw err;
    }
  };

export const applyConfig = (id: string) => async (dispatch: AppDispatch) => {
  try {
    await fetch(`http://localhost:8080/api/config/apply/${id}`, {
      method: "PUT",
    });
    await dispatch(getConfigs());
  } catch (err) {
    console.error("Cannot apply config", err);
    throw err;
  }
};

export const updateConfig =
  (id: string, updatedConfig: Config) => async (dispatch: AppDispatch) => {
    try {
      await fetch(`http://localhost:8080/api/config/${id}`, {
        method: "PUT",
        body: JSON.stringify(updatedConfig),
      });
      dispatch(initializeCreateConfigForm());
      await dispatch(getConfigs());
    } catch (err) {
      console.error("Cannot update config", err);
      throw err;
    }
  };

export const deleteConfig = (id: string) => async (dispatch: AppDispatch) => {
  try {
    await fetch(`http://localhost:8080/api/config/${id}`, {
      method: "DELETE",
    });
    await dispatch(getConfigs());
  } catch (err) {
    console.error("Cannot delete config", err);
    throw err;
  }
};
