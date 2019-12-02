import { Result } from "../interfaces";

export function search(): Promise<Result[]> {
  return new Promise(() => {
    getResults().then(results => {
      return results;
    })
  });
}

function getResults(): Promise<Result[]> {
  if (process.platform === "win32") {
    return require("./windows");
  } else if (process.platform === "darwin") {
    return require("./mac");
  } else if (process.platform === "linux") {
    return require("./linux");
  } else {
    return null;
  }
}

export default search;
