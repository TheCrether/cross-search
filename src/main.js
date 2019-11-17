const { app, BrowserWindow } = require("electron");
const isDev = require("electron-is-dev");
const path = require("path");
const conf = require("./lib/config");
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
    fullscreenable: false
  });

  mainWindow.setAlwaysOnTop(true, "modal-panel");

  // and load the index.html of the app.
  if (isDev) {
    mainWindow.loadURL("http://localhost:3000");
  } else {
    mainWindow.loadURL(`file://${path.join(__dirname, "../build/index.html")}`);
  }
  mainWindow.webContents.openDevTools();
  mainWindow.on("closed", () => (mainWindow = null));
}
app.on("ready", createWindow);

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
