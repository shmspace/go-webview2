package main

import (
	webview "github.com/jchv/go-webview2"
	"log"
	"time"
)

func main() {
	var w webview.WebView
	go func() {
		w = webview.NewWithOptions(webview.WebViewOptions{
			Debug:     true,
			AutoFocus: true,
			WindowOptions: webview.WindowOptions{
				Title:  "Minimal webview example",
				Width:  800,
				Height: 600,
				IconId: 2, // icon resource id
				Center: true,
			},
		})
		if w == nil {
			log.Fatalln("Failed to load webview.")
		}
		defer w.Destroy()
		w.SetSize(800, 600, webview.HintFixed)
		w.Navigate("https://www.taobao.com")
		w.Run()
	}()
	go func() {
		for {
			time.Sleep(3 * time.Second)
			w.Show()
			//w.Show()
			time.Sleep(3 * time.Second)
			w.Hide()
			w.Navigate("https://www.baidu.com")
			//w.Hide()
		}
	}()
	time.Sleep(5 * time.Minute)
}
