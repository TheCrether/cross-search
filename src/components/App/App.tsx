import React from "react";
import "./App.scss";

const { remote } = window.require("electron");

const App: React.FC = () => {
  return (
    <div className="app">{remote.getCurrentWindow().getBounds().height}</div>
  );
};

export default App;
