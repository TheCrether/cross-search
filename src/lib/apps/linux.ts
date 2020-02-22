import { homedir } from 'os';
import { Result } from '../interfaces';
import { readFileSync, existsSync } from "fs";
import { join } from 'path';

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



  return app;
}

function getIcon(icon: string) {
  if (icon.includes("")) {
    console.log("ree");
  }
}

export function getResult(path: string): Result {
  return parseDesktopFile(path);
}
