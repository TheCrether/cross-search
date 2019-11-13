const { app, BrowserWindow } = require("electron");
const path = require("path");
let win;
function createWindow() {
  // Create the browser window.
  win = new BrowserWindow({ width: 800, height: 600 });

  // and load the index.html of the app.
  win.loadURL(`file://${path.join(__dirname, "../build/index.html")}`);
  win.webContents.openDevTools();
}
app.on("ready", createWindow);
