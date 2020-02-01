import { BASE_HEIGHT, HOTKEYS } from './lib/config';
import { app, BrowserWindow, globalShortcut, } from "electron";
import isDev from "electron-is-dev";
import { join } from "path";
import toggleWindow from "./lib/window/toggleWindow";


let mainWindow: BrowserWindow;
function createWindow() {
  // Create the browser window.
  mainWindow = new BrowserWindow({
    width: 600,
    height: BASE_HEIGHT,
    autoHideMenuBar: true,
    backgroundColor: "#00000000",
    alwaysOnTop: true,
    frame: false,
    minimizable: false,
    center: true,
    resizable: false,
    webPreferences: {
      nodeIntegration: true,
      preload: join(app.getAppPath(), "src", 'preload.js')
    },
    title: "cross-search",
    transparent: true,
    titleBarStyle: "hidden",
    fullscreenable: false,
    hasShadow: true
  });

  mainWindow.setBounds({ y: 100 })

  mainWindow.setAlwaysOnTop(true, "modal-panel");

  if (isDev) {
    mainWindow.loadURL("http://localhost:3000");
    mainWindow.webContents.openDevTools();
  } else {
    mainWindow.loadURL(`file://${join(__dirname, "../build/index.html")}`);
  }
  mainWindow.on("closed", () => (mainWindow = null));
  mainWindow.focus();

  globalShortcut.registerAll(HOTKEYS, () => toggleWindow(mainWindow));
}

app.on("quit", () => app.quit())

app.on("ready", () => setTimeout(createWindow, 300));

app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    app.quit();
  }
});

app.on("activate", () => {
  if (mainWindow === null) {
    createWindow();
  }
});
