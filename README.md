# cross-search

An cross-platform tool for searching desktop apps etc.

- [cross-search](#cross-search)
  - [About](#about)
  - [Available Scripts](#available-scripts)
    - [`npm run dev`](#npm-run-dev)
    - [`npm run dev:win`](#npm-run-devwin)
    - [`npm run test`](#npm-run-test)
    - [`npm run build`](#npm-run-build)
    - [`npm run eject`](#npm-run-eject)

## About

I wanted to have something like [cerebro](https://github.com/KELiON/cerebro/) and spotlight for Linux and Windows.<br>
Inspirations are from these two applications and also some code parts may be reused from cerebro for easiness.<br>
This project is made with Electron

## Available Scripts

### `npm run dev`

Launches the react app (`npm run start`) without a browser and then the electron app

### `npm run dev:win`

Does the same as above but for windows because you cant set environment variables like on Linux;

### `npm run test`

Launches the test runner in the interactive watch mode.<br />
See the section about [running tests](https://facebook.github.io/create-react-app/docs/running-tests) for more information.

### `npm run build`

Builds the app for production to the `build` folder.<br />
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.<br />
Your app is ready to be deployed!

See the section about [deployment](https://facebook.github.io/create-react-app/docs/deployment) for more information.

### `npm run eject`

**Note: this is a one-way operation. Once you `eject`, you can’t go back!**

If you aren’t satisfied with the build tool and configuration choices, you can `eject` at any time. This command will remove the single build dependency from your project.

Instead, it will copy all the configuration files and the transitive dependencies (Webpack, Babel, ESLint, etc) right into your project so you have full control over them. All of the commands except `eject` will still work, but they will point to the copied scripts so you can tweak them. At this point you’re on your own.

You don’t have to ever use `eject`. The curated feature set is suitable for small and middle deployments, and you shouldn’t feel obligated to use this feature. However we understand that this tool wouldn’t be useful if you couldn’t customize it when you are ready for it.
