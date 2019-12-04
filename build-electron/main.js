"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
exports.__esModule = true;
var config_1 = require("./lib/config");
var electron_1 = require("electron");
var electron_is_dev_1 = __importDefault(require("electron-is-dev"));
var path_1 = require("path");
var toggleWindow_1 = __importDefault(require("./lib/window/toggleWindow"));
var mainWindow;
function createWindow() {
    // Create the browser window.
    mainWindow = new electron_1.BrowserWindow({
        width: 600,
        height: config_1.BASE_HEIGHT,
        autoHideMenuBar: true,
        backgroundColor: "#00000000",
        alwaysOnTop: true,
        frame: false,
        minimizable: false,
        center: true,
        resizable: false,
        webPreferences: {
            nodeIntegration: true,
            preload: path_1.join(electron_1.app.getAppPath(), "src", 'preload.js')
        },
        title: "cross-search",
        transparent: true,
        titleBarStyle: "hidden",
        fullscreenable: false,
        hasShadow: true
    });
    mainWindow.setAlwaysOnTop(true, "modal-panel");
    if (electron_is_dev_1["default"]) {
        mainWindow.loadURL("http://localhost:3000");
        mainWindow.webContents.openDevTools();
        // TODO REMOVE WHEN EXTENSIONSFOLDER CAN BE FOUND LINUX
        Promise.resolve().then(function () { return __importStar(require("electron-devtools-installer")); }).then(function (installer) {
            installer(installer.REACT_DEVELOPER_TOOLS)
                .then(function (name) { return console.log("Added Extension:  " + name); })["catch"](function (err) { return console.log('An error occurred: ', err); });
        });
    }
    else {
        mainWindow.loadURL("file://" + path_1.join(__dirname, "../build/index.html"));
    }
    mainWindow.on("closed", function () { return (mainWindow = null); });
    mainWindow.focus();
    electron_1.globalShortcut.registerAll(config_1.HOTKEYS, function () { return toggleWindow_1["default"](mainWindow); });
}
electron_1.app.on("quit", function () { return electron_1.app.quit(); });
electron_1.app.on("ready", function () { return setTimeout(createWindow, 300); });
electron_1.app.on("window-all-closed", function () {
    if (process.platform !== "darwin") {
        electron_1.app.quit();
    }
});
electron_1.app.on("activate", function () {
    if (mainWindow === null) {
        createWindow();
    }
});
