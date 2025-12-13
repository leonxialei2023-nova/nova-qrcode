package main

import (
	"encoding/base64"
	"net/http"

	"nova/qrcode/pkg/qrcode"
)

func handler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "missing text", 400)
		return
	}

	png, err := qrcode.EncodePNG(text, qrcode.Options{
		Size:  256,
		Level: qrcode.Medium,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}

func base64Handler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")

	png, _ := qrcode.EncodePNG(text, qrcode.Options{})
	b64 := base64.StdEncoding.EncodeToString(png)

	w.Write([]byte(b64))
}

func main() {
	http.HandleFunc("/png", handler)
	http.HandleFunc("/base64", base64Handler)

	http.ListenAndServe(":8080", nil)
}
