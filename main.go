package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Sunucu başlatılıyor...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL Kısaltıcıya Hoşgeldiniz!")

	})

	http.ListenAndServe(":8080", nil)
}
