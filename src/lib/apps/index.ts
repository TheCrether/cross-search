import { Result } from "../interfaces";
import fs, { exists, readdirSync } from "fs";


export function search(): Promise<Result[]> {
  return new Promise(() => {
    let { DIRS, EXTS, getResult } =
      getPlatform();
    const directories: string[] = DIRS;
    const extensions: string[] = EXTS;

    directories.forEach(dir => {
      fs.exists(dir, (ex => {
        if (!ex) {
          return;
        }
        const files = readdirSync(dir);
        // TODO change
        files.filter(file => file.includes(extensions[0]))
          .forEach(file => {
            console.log(getResult(file))
          })
      }))
    })
    return [{ icon: "", name: "none", path: "" }] as Result[];
  });
}

function getPlatform(): any {
  if (global.process.platform === "win32") {
    return require("./win");
  } else if (global.process.platform === "darwin") {
    return require("./mac");
  } else if (global.process.platform === "linux") {
    return require("./linux");
  } else {
    return null;
  }
}

export default search;
