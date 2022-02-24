import { createSelector } from "@reduxjs/toolkit";
import { RootState } from "../app/store";
import { Status, StatusState } from "./statusTypes";

export const selectStatus = createSelector(
  (state: RootState) => state.status,
  (status: StatusState) => status
);
