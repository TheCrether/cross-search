import React, { Component } from "react";
import "./App.scss";
import { BASE_HEIGHT, HEADER_HEIGHT } from "../../lib/config";
import toggleWindow from "../../lib/window/toggleWindow";
import { Results } from "../Results/Results";
import apps from "../../lib/apps";
import { app } from "electron";

const { remote } = window.require("electron");

export class App extends Component<{}, {}> {
  state = {
    resultList: []
  }

  componentDidMount() {
    apps().then(results => {
      console.log(results);
      this.setState({ resultList: [this.state.resultList, ...results] })
    })
  }

  render() {
    console.log(this.state.resultList)
    return (
      <div className="app">
        <div id="header" style={{ height: HEADER_HEIGHT }}>
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
        <Results resultList={this.state.resultList}></Results>
      </div>
    );
  }
}

export default App;
