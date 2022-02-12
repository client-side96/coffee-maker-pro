import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import sensorReducer from "../sensors/sensorReducer";

export const store = configureStore({
  reducer: {
    sensor: sensorReducer,
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
