import { Status, StatusResponse, StatusState } from "./statusTypes";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const initialState: StatusState = {
  status: "Off",
};

const statusSlice = createSlice({
  name: "status",
  initialState,
  reducers: {
    setStatus(state, action: PayloadAction<StatusResponse>) {
      state.status = action.payload.value;
      state.id = action.payload._id;
    },
  },
});

export const { setStatus } = statusSlice.actions;
export default statusSlice.reducer;
