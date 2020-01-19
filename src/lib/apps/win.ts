import { Result } from './../interfaces';
import { join, sep } from "path";
import { shell } from "electron";
import { homedir } from "os";
import { lstatSync } from 'fs';

const { APPDATA, ProgramData } = window.process.env;
let DESKTOP = join(homedir(), "Desktop");

try {
  lstatSync(join(homedir(), "Desktop"));
} catch (e) {
  try {
    DESKTOP = shell.readShortcutLink(join(homedir(), "Desktop.lnk")).target;
  } catch (e) {
    DESKTOP = join(homedir(), "Desktop");
  }
}


export const DIRS = [
  DESKTOP,
  join(APPDATA, "Microsoft", "Windows", "Start Menu", "Programs"),
  join(ProgramData, "Microsoft", "Windows", "Start Menu", "Programs")
];

export const EXTS = [
  "lnk", "exe"
];

export function getResult(path: string): Result {
  const details = shell.readShortcutLink(path);

  const split = details.target.split(sep);

  return {
    icon: details.icon || "",
    name: split[split.length - 1].replace(/\.*/, ""),
    exec: details.target
  } as Result;
}
