import { Result } from "../interfaces";

export function search(): Promise<Result[]> {
  return new Promise(() => {
    let { DIRS, EXTS, getResult } = getPlatform();
    const directories: string[] = DIRS;
    const extensions: string[] = EXTS;
  });
}

function getPlatform(): any {
  if (process.platform === "win32") {
    return require("./win");
  } else if (process.platform === "darwin") {
    return require("./mac");
  } else if (process.platform === "linux") {
    return require("./linux");
  } else {
    return null;
  }
}

export default search;
