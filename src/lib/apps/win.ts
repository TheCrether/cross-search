import path from "path";
import { shell } from "electron";

const { APPDATA, ProgramData, USERPROFILE } = process.env;

export const DIRS = [
  USERPROFILE && `${USERPROFILE}\\Desktop\\`,
  APPDATA && `${APPDATA}\\Microsoft\\Windows\\Start Menu\\Programs\\`,
  ProgramData && `${ProgramData}\\Microsoft\\Windows\\Start Menu\\Programs\\`
];

export const exts = [
  "lnk", "exe"
];
