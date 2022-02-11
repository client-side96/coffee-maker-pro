import React from "react";

function App() {
  const url = "ws://localhost:8080/api/sensors";
  const c = new WebSocket(url);

  const send = (data: string) => {
    c.send(data);
  };

  c.onmessage = function (msg) {
    console.log(msg);
  };

  c.onopen = function () {};

  return <>Test</>;
}

export default App;
