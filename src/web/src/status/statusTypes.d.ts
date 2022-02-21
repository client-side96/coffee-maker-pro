export type Status = "Off" | "Idle" | "Applying" | "Ready" | "Brewing";

export interface StatusState {
  id?: string;
  status: Status;
}

export interface StatusResponse {
  _id: string;
  value: Status;
}
