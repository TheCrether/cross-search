import { Result } from './../interfaces';
import { join, sep } from "path";
import { shell } from "electron";

const { APPDATA, ProgramData, USERPROFILE } = process.env;

export const DIRS = [
  USERPROFILE && join(USERPROFILE, "Desktop"),
  APPDATA && join(APPDATA, "Microsoft", "Windows", "Start Menu", "Programs"),
  ProgramData && join(ProgramData, "Microsoft", "Windows", "Start Menu", "Programs")
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
