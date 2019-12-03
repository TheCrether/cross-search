import { Result } from './../interfaces';
import { readFileSync } from 'fs';
export const DIRS = ["~/.local/share/applications/", "/usr/share/applications/", "/usr/local/share/applications/"];

export const EXTS = [".desktop"];

export function parseDesktopFile(path: string): Result {
  const file = readFileSync(path).toString();

  const keys = ["Name"];

  let app: Result;
  keys.forEach(key => {
    const regex = new RegExp("^RegExr=(.+)", "m");
    const match = file.match(regex);
    console.log(match);
  })

  return app;
}

export function getResult(path: string): Result {
  return parseDesktopFile(path);
}

