package main

import (
	"embed"
	"io/fs"
	"net/http"

	webview "github.com/webview/webview_go"
)

//go:embed ui/dist/*
var staticFiles embed.FS

func main() {
	staticSubFS, _ := fs.Sub(staticFiles, "ui/dist")
	http.Handle("/", http.FileServer(http.FS(staticSubFS)))
	go http.ListenAndServe(":8080", nil)

	w := webview.New(true)
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://localhost:8080")
	w.Run()
}
