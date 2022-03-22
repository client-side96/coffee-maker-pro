import React from "react";
import { FastField, Form, Formik } from "formik";
import "./ConfigForm.scss";
import { Config } from "../configTypes";
import { useAppDispatch } from "../../app/hooks";
import { createConfig, updateConfig } from "../configThunks";

export type ConfigFormProps = { formValues: Config };

const ConfigForm: React.FC<ConfigFormProps> = ({ formValues }) => {
  const dispatch = useAppDispatch();
  const handleSubmit = (values: Config) => {
    if (values._id) {
      dispatch(updateConfig(values._id, values));
    } else {
      dispatch(createConfig(values));
    }
  };
  return (
    <div className="config-form">
      <h3>
        {formValues._id
          ? "Konfiguration bearbeiten"
          : "Neue Konfiguration erstellen"}
      </h3>
      <Formik
        initialValues={formValues}
        enableReinitialize
        onSubmit={(values) => {
          handleSubmit(values);
        }}
      >
        {({ handleChange }) => (
          <Form>
            <label>Name:</label>
            <FastField name="name" onChange={handleChange} />
            <label>Temperature:</label>
            <FastField name="temp" type="number" onChange={handleChange} />
            <label>Pressure:</label>
            <FastField name="pressure" type="number" onChange={handleChange} />
            <label>Grinding:</label>
            <FastField name="grinding" type="number" onChange={handleChange} />
            <label>Volume:</label>
            <FastField name="volume" type="number" onChange={handleChange} />
            <label>Time:</label>
            <FastField name="time" type="number" onChange={handleChange} />
            <button type="submit">Save</button>
          </Form>
        )}
      </Formik>
    </div>
  );
};

export default ConfigForm;
