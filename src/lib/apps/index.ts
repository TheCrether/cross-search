import { Result } from "../interfaces";
import { existsSync, readdirSync, statSync } from "fs"
import { join } from "path";
import { homedir } from "os";

const { DIRS, EXTS, getResult } = getPlatform();
const directories: string[] = DIRS;
const extensions: string[] = EXTS;

export function search(): Promise<Result[]> {
  return new Promise((resolve) => {
    let result: Result[] = [];

    directories.forEach(dir => {
      dir = dir.replace("~", homedir())
      try {
        if (existsSync(dir)) {
          const files: string[] = readdirSync(dir);
          const entries: Result[] = [];
          getDesktopEntries(files, entries, dir);
          result = [...result, ...entries];
        }
      } catch (error) {
        console.error((error as Error).message);
      }
    })
    resolve(result);
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

function getDesktopEntries(files: string[], entries: Result[], baseDir: string): void {
  const realFiles: string[] = [];
  const dirs: string[] = [];
  files.forEach(file => {
    if (existsSync(join(baseDir, file))) {
      if (statSync(join(baseDir, file)).isDirectory()) {
        dirs.push(file);
      } else {
        for (const ext of extensions) {
          if (file.endsWith(ext)) {
            realFiles.push(file);
            break;
          }
        }
      }
    }
  });
  realFiles.forEach(file => {
    try {
      const result = getResult(join(baseDir, file))
      if (result) entries.push(result)
    } catch (e) {
      console.error(join(baseDir, file));
    }
  });
  dirs.forEach(dir => {
    try {
      const results: Result[] = [];
      getDesktopEntries(readdirSync(join(baseDir, dir)), results, join(baseDir, dir));
      entries = [...entries, ...results];
    } catch (error) {
      console.error(error);
    }
  })
}