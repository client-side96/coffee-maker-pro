import React from "react";
import { turnOff, turnOn } from "../statusThunks";
import "./PowerButton.scss";

type PowerButtonProps = {
  isPowerON: boolean;
};

const PowerButton: React.FC<PowerButtonProps> = ({ isPowerON }) => {
  const handlePower = () => {
    if (isPowerON) {
      turnOff();
    } else {
      turnOn();
    }
  };
  return (
    <a className="power-btn" onClick={handlePower}>
      {isPowerON ? "Ausschalten" : "Anschalten"}
    </a>
  );
};

export default PowerButton;
