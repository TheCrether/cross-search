import React from "react";
import "./App.scss";
import { INPUT_HEIGHT } from "../../lib/config";

const { remote } = window.require("electron");

const App: React.FC = () => {
  return (
    <div className="app">
      <div className="header">{/* TODO make header with dragging etc */}</div>
      <div className="input">
        <input type="text" style={{ height: INPUT_HEIGHT + "px" }} placeholder="Search" />
      </div>
    </div>
  );
};

export default App;
