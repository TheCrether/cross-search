import { Result, Platform } from "../interfaces";
const { DIRS, EXTS } = getPlatform();

export function search(): Promise<Result[]> {
  return new Promise(() => {
    
    return [{ id: 1, name: "lel", icon: "" }];
  });
}

function getPlatform(): Partial<Platform> {
  let platform;

  if (process.platform === "win32") {
    platform = require("./windows");
  } else if (process.platform === "darwin") {
    platform = require("./mac");
  } else if (process.platform === "linux") {
    platform = require("./linux");
  }

  return platform;
}

export default search;
