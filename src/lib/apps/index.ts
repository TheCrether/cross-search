import { Result } from "../interfaces";
import fs, { exists, readdirSync } from "fs";
import os, { platform, networkInterfaces } from "os";


export function search(): Promise<Result[]> {
  return new Promise(() => {
    // let { DIRS, EXTS, getResult } = 
    getPlatform();
    // const directories: string[] = DIRS;
    // const extensions: string[] = EXTS;

    // directories.forEach(dir => {
    //   exists(dir, (exists => {
    //     if (!exists) {
    //       return;
    //     }
    //     const files = readdirSync(dir);
    //     // TODO change
    //     files.filter(file => file.includes(extensions[0]))
    //       .forEach(file => {
    //         console.log(getResult(file))
    //       })
    //   }))
    // })
    return [{ icon: "", name: "none", path: "" }] as Result[];
  });
}

function getPlatform(): any {
  console.log(process);
  if (process.platform === "win32") {
    const lel: any = require("./win");
    console.log(lel);
    return lel;
  } else if (process.platform === "darwin") {
    return require("./mac");
  } else if (process.platform === "linux") {
    return require("./linux");
  } else {
    return null;
  }
}

export default search;
