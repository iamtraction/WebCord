package main

import "github.com/webview/webview"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("WebCord for Discord")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("https://canary.discord.com/app")
	w.Run()
}
