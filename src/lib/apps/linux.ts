import { app } from 'electron';
import { homedir } from 'os';
import { Result } from '../interfaces';
import { readFileSync, existsSync } from "fs";
import { join } from 'path';
import { spawnSync } from 'child_process';

const { XDG_DATA_DIRS } = window.process.env;
const additionalDirs = [`${homedir()}/.local/share`]
// console.log([...XDG_DATA_DIRS.split(":"), ...additionalDirs])
export const DIRS = [...XDG_DATA_DIRS.split(":"), ...additionalDirs].map((dir) => join(dir, "applications")).filter(existsSync)

export const EXTS = [".desktop"];

export function parseDesktopFile(path: string): Result {
  const file = readFileSync(path).toString();

  const keys = ["Name", "Exec", "Icon"];

  const disqualifying = ["NoDisplay", "Hidden"];

  let notAdd = false;

  const app: Result = {
    icon: "",
    name: "",
    exec: ""
  };
  disqualifying.forEach(dis => {
    const regex = new RegExp(`^${dis}=(.+)`, "m");
    const match = file.match(regex);
    if (match) {
      notAdd = true;
    }
  });
  keys.forEach(key => {
    const regex = new RegExp(`^${key}=(.+)`, "m");
    const match = file.match(regex);
    if (match) {
      app[key.toLowerCase()] = match[1];
    }
  });

  if (!app.name || notAdd) {
    return null;
  }

  getIcon(app.icon)

  return app;
}

function getIcon(icon: string): string {
  const regex = new RegExp(".*.(svg|png|xpm|gif|ico)$");
  if (regex.exec(icon)) {
    return icon;
  } else {
    try {
      // const python = spawnSync("python", [join(app.getAppPath(), "util", "get-icon.py"), `"${icon}"`]);
      // console.log(python.output.toString());
      // eslint-disable-next-line no-empty
    } catch (e) {
      console.log("error")
    }
  }
}

export function getResult(path: string): Result {
  return parseDesktopFile(path);
}
