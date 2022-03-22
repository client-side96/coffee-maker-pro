import React from "react";
import "./ConfigListItem.scss";
import { Config } from "../configTypes";

export type ConfigListItemProps = {
  config: Config;
  onEdit: (id: string | null) => void;
  onDelete: (id: string | null) => void;
  onApply: (id: string | null) => void;
};

const ConfigListItem: React.FC<ConfigListItemProps> = ({
  config,
  onEdit,
  onDelete,
  onApply,
}) => {
  return (
    <div className="config-item">
      <h4 className="config-name">{config.name}</h4>
      <p>Wassertemperatur: {config.temp} Celsius</p>
      <p>Wasserdruck: {config.pressure} bar</p>
      <p>Mahlgrad: {config.grinding} Step</p>
      <p>Wasservolumen: {config.volume} ml</p>
      <p>Durchlaufzeit: {config.time} Sekunden</p>
      <div className="config-actions">
        <a onClick={() => onEdit(config._id)}>Bearbeiten</a>
        <a onClick={() => onDelete(config._id)}>Löschen</a>
        {config.isApplied ? (
          <p style={{ color: "green", display: "inline-block" }}>Aktiv</p>
        ) : (
          <a onClick={() => onApply(config._id)}>Aktivieren</a>
        )}
      </div>
    </div>
  );
};

export default ConfigListItem;
