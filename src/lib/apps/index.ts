import { Result } from "../interfaces";
const { existsSync, readdirSync, statSync } = window.require("fs");
const { join } = window.require("path");
const { homedir } = window.require("os");
const { DIRS, EXTS, getResult } = getPlatform();
const directories: string[] = DIRS;
const extensions: string[] = EXTS;

export function search(): Promise<Result[]> {
  return new Promise((resolve, reject) => {
    let result: Result[] = [];

    directories.forEach(dir => {
      dir = dir.replace("~", homedir())
      if (!existsSync(dir)) return;
      const files: string[] = readdirSync(dir);
      const entries: Result[] = [];
      getDesktopEntries(files, entries, dir);
      result = [...result, ...entries];
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
        for (let ext of extensions) {
          if (file.endsWith(ext)) {
            realFiles.push(file);
            break;
          }
        }
      }
    }
  });
  realFiles.forEach(file => {
    const result = getResult(join(baseDir, file))
    if (result) entries.push(result)
  });
  dirs.forEach(dir => {
    const results: Result[] = [];
    getDesktopEntries(readdirSync(dir), results, join(baseDir, dir));
    entries = [...entries, ...results];
  })
}