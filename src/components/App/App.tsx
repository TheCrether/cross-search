import React from "react";
import "./App.scss";
import { BASE_HEIGHT, HEADER_HEIGHT } from "../../lib/config";
import toggleWindow from "../../lib/window/toggleWindow";


const { remote } = window.require("electron");

const App: React.FC = () => {
  return (
    <div className="app">
      <div id="header" style={{ height: HEADER_HEIGHT }}>
        {/* TODO make header with dragging etc */}
        <div className="title">cross-search</div>
        <button
          id="close"
          onClick={() => toggleWindow(remote.getCurrentWindow())}
        >
          <i className="material-icons">clear</i>
        </button>
      </div>
      <div className="input">
        <input
          type="text"
          style={{ height: BASE_HEIGHT - HEADER_HEIGHT + "px" }}
          placeholder="Search"
        />
      </div>
    </div>
  );
};

export default App;
