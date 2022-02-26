import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import sensorReducer from "../sensors/sensorReducer";
import statusReducer from "../status/statusReducer";
import configReducer from "../config/configReducer";

export const store = configureStore({
  reducer: {
    sensor: sensorReducer,
    status: statusReducer,
    config: configReducer,
  },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;
