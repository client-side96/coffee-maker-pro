import React from "react";
import "./ConfigListItem.scss";
import { Config } from "../configTypes";

export type ConfigListItemProps = {
  config: Config;
  onEdit: (id: string | null) => void;
  onDelete: (id: string | null) => void;
};

const ConfigListItem: React.FC<ConfigListItemProps> = ({
  config,
  onEdit,
  onDelete,
}) => {
  return (
    <div className="config-item">
      <h4 className="config-name">{config.name}</h4>
      <p>Temperature: {config.temp} Celsius</p>
      <p>Pressure: {config.pressure} bar</p>
      <p>Grinding: {config.grinding}</p>
      <div className="config-actions">
        <a onClick={() => onEdit(config._id)}>Edit</a>
        <a onClick={() => onDelete(config._id)}>Delete</a>
        <a>{config.isApplied ? "Applied" : "Apply"}</a>
      </div>
    </div>
  );
};

export default ConfigListItem;
