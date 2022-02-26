import React from "react";
import "./ConfigList.scss";
import { Config } from "../configTypes";
import ConfigListItem from "./ConfigListItem";
import { useAppDispatch } from "../../app/hooks";
import { initializeUpdateConfigForm } from "../configReducer";
import { deleteConfig, getConfigById } from "../configThunks";

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
  return (
    <div className="config-list">
      <h3>Saved configurations:</h3>
      {configs && configs.length ? (
        configs.map((config, index) => (
          <ConfigListItem
            key={`${config._id}-${index}`}
            config={config}
            onEdit={handleEdit}
            onDelete={handleDelete}
          />
        ))
      ) : (
        <p>No configs available</p>
      )}
    </div>
  );
};

export default ConfigList;
