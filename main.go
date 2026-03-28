package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var urlMap = make(map[string]string)

func kisaKodUret() string {
	harfler := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	kod := ""
	for i := 0; i < 6; i++ {
		kod += string(harfler[rand.Intn(len(harfler))])
	}
	return kod
}

func main() {

	fmt.Println("Sunucu başlatılıyor...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL Kısaltıcıya Hoşgeldiniz!")

	})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		uzunLink := r.URL.Query().Get("url")
		fmt.Fprintln(w, uzunLink)
		kisaKod := kisaKodUret()
		urlMap[kisaKod] = uzunLink
		fmt.Fprintln(w, "Kısa linkin: localhost:8080/r/"+kisaKod)
	})

	http.ListenAndServe(":8080", nil)
}
