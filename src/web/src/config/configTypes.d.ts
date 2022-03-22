export interface Config {
  _id: string | null;
  name: string;
  temp: number;
  pressure: number;
  grinding: number;
  volume: number;
  time: number;
  isApplied?: boolean;
}

export interface ConfigState {
  configs: Config[];
  formValues: Config | null;
}
