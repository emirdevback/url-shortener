package main

import (
	"fmt"
	"net/http"
)

var urlMap = make(map[string]string)

func main() {

	fmt.Println("Sunucu başlatılıyor...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL Kısaltıcıya Hoşgeldiniz!")

	})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Link kısaltma endpoint'i")
	})

	http.ListenAndServe(":8080", nil)
}
