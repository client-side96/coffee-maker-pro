import React from "react";
import "./ConfigList.scss";
import { Config } from "../configTypes";
import ConfigListItem from "./ConfigListItem";
import { useAppDispatch } from "../../app/hooks";
import { initializeUpdateConfigForm } from "../configReducer";
import { applyConfig, deleteConfig, getConfigById } from "../configThunks";

export type ConfigListProps = {
  configs: Config[];
};

const ConfigList: React.FC<ConfigListProps> = ({ configs }) => {
  const dispatch = useAppDispatch();
  const handleEdit = (id: string | null) => {
    if (id) {
      dispatch(getConfigById(id));
    }
  };
  const handleDelete = (id: string | null) => {
    if (id) {
      dispatch(deleteConfig(id));
    }
  };
  const handleApply = (id: string | null) => {
    if (id) {
      dispatch(applyConfig(id));
    }
  };
  return (
    <div className="config-list">
      <h3>Gespeicherte Konfigurationen:</h3>
      {configs && configs.length ? (
        configs.map((config, index) => (
          <ConfigListItem
            key={`${config._id}-${index}`}
            config={config}
            onEdit={handleEdit}
            onDelete={handleDelete}
            onApply={handleApply}
          />
        ))
      ) : (
        <p>Bisher wurden keine Konfigurationen abgespeichert.</p>
      )}
    </div>
  );
};

export default ConfigList;
