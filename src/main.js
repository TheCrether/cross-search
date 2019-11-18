// jshint ignore: start
const { app, BrowserWindow, globalShortcut } = require("electron");
const isDev = require("electron-is-dev");
const path = require("path");
const conf = require("./lib/config");
const toggleWindow = require("./lib/window/toggleWindow");
let mainWindow;
function createWindow() {
  // Create the browser window.
  mainWindow = new BrowserWindow({
    width: 600,
    height: conf.INPUT_HEIGHT + 20,
    autoHideMenuBar: true,
    backgroundColor: "#00000000",
    alwaysOnTop: true,
    frame: false,
    minimizable: false,
    center: true,
    resizable: false,
    webPreferences: {
      nodeIntegration: true
    },
    title: "cross-search",
    transparent: true,
    titleBarStyle: "hidden",
    fullscreenable: false,
    hasShadow: true
  });

  mainWindow.setAlwaysOnTop(true, "modal-panel");

  // and load the index.html of the app.
  if (isDev) {
    mainWindow.loadURL("http://localhost:3000");
    mainWindow.webContents.openDevTools();
  } else {
    mainWindow.loadURL(`file://${path.join(__dirname, "../build/index.html")}`);
  }
  mainWindow.on("closed", () => (mainWindow = null));
  mainWindow.focus();

  globalShortcut.register(conf.HOTKEY, () => toggleWindow(mainWindow));
}

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
