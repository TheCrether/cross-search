import React from "react";
import { Result } from "../../lib/interfaces";
import "./results.scss";

export class Results extends React.Component<Props, {}> {
  render() {
    return <div id="results"></div>;
  }
}

export interface Props {
  resultList: Result[];
}
