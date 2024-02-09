<div align="center">

<h1>go-react</h1>

![image](https://github.com/anurag-roy/go-react/assets/53750093/0b6a3c12-e3be-4f4c-99b2-4401f9946ef8)

</div>

---

This is a simple Go executable with an embedded React app.

This particular implementation uses:
- [embed](https://pkg.go.dev/embed) to embed static assets in the go binary
- [webview_go](https://github.com/webview/webview_go) to run HTML, CSS, Js without bundling a browser
- [vite](https://vitejs.dev/) to scaffold and bundle the UI

Essentially any JavaScript UI framework/library can be used like Angular, Vue, Svelte, Solid, etc as long as the generated code doesn't depend on Node.js and is runnable on a browser.

## Prerequisites

These need to be installed on your system to be able to develop/build the app:
- [webview](https://github.com/webview/webview)
- [go](https://go.dev)
- [node.js](https://nodejs.org)

## Build

Build the ui first to generate your static assets. Inside the `ui` folder, run:
```
npm run build
```

Then build the binary. Inside the root folder, run
```
go build
```

That's it! The size of the binary on linux-amd64 comes to ~7.3MB.
