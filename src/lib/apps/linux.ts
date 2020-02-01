import { Result } from '../interfaces';
import { readFileSync } from "fs";
export const DIRS = ["~/.local/share/applications/", "/usr/share/applications/", "/usr/local/share/applications/"];

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

export function getResult(path: string): Result {
  return parseDesktopFile(path);
}
