import { Result } from './../interfaces';
const { readFileSync } = window.require('fs');
export const DIRS = ["~/.local/share/applications/", "/usr/share/applications/", "/usr/local/share/applications/"];

export const EXTS = [".desktop"];

export function parseDesktopFile(path: string): Result {
  const file = readFileSync(path).toString();

  const keys = ["Name", "Exec"];

  let app: Result = {
    icon: "",
    name: "",
    exec: ""
  };
  keys.forEach(key => {
    const regex = new RegExp(`^${key}=(.+)`, "m");
    const match = file.match(regex);
    if (match) {
      app[key.toLowerCase()] = match[1];
    }
  })

  if (!app.name) {
    return null;
  }

  return app;
}

export function getResult(path: string): Result {
  return parseDesktopFile(path);
}

export function getIcon(result: Result): Result {
  
  return null;
}

