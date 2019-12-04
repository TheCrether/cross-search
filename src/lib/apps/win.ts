import { Result } from './../interfaces';
import { join } from "path";
import { shell } from "electron";

const { APPDATA, ProgramData, USERPROFILE } = process.env;

export const DIRS = [
  USERPROFILE && join(USERPROFILE, "Desktop"),
  APPDATA && join(APPDATA, "Microsoft", "Windows", "Start Menu", "Programs"),
  ProgramData && join(ProgramData, "Microsoft", "Windows", "Start Menu", "Programs")
];

export const exts = [
  "lnk", "exe"
];

export function getResult(path: string): Result {
  const details = shell.readShortcutLink(path);

  return {
    icon: details.icon || "",
    name: "test",
    exec: details.target
  } as Result;
}
