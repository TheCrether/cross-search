"use strict";
exports.__esModule = true;
/**
 * Show or hide main window
 * @return {BrowserWindow} appWindow
 */
exports["default"] = (function (appWindow) {
    if (appWindow.isVisible()) {
        appWindow.blur(); // once for blurring the content of the window(?)
        // appWindow.blur(); // twice somehow restores focus to prev foreground window
        appWindow.hide();
    }
    else {
        appWindow.show();
        appWindow.focus();
    }
});
