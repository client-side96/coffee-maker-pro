import { createSelector } from "@reduxjs/toolkit";
import { RootState } from "../app/store";
import { Config } from "./configTypes";

export const selectConfigs = createSelector(
  (state: RootState) => state.config.configs,
  (configs: Config[]) => configs
);

export const selectConfigFormValues = createSelector(
  (state: RootState) => state.config.formValues,
  (formValues: Config | null) => formValues
);
