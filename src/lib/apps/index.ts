import { Result } from "../interfaces";
const { exists, readdirSync } = window.require("fs");
const { join } = window.require("path");
const { DIRS, EXTS, getResult } = getPlatform();
const directories: string[] = DIRS;
const extensions: string[] = EXTS;

export function search(): Promise<Result[]> {
  return new Promise((resolve, reject) => {

    directories.forEach(dir => {
      exists(dir, (ex => {
        if (!ex) {
          return;
        }
        const files: string[] = readdirSync(dir);
        const entries: Result[] = [];
        getDesktopEntries(files, entries, dir);
        return resolve(entries);
      }))

    })
    return [] as Result[];
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
    let yes = true;
    for (let ext of extensions) {
      if (file.endsWith(ext)) {
        realFiles.push(file);
        break;
      }
    }
    if (!yes) {
      dirs.push(file);
    }
  })
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