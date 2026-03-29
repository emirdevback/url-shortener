package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/skip2/go-qrcode"
)

var urlMap = make(map[string]string)
var db *sql.DB

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

	var err error
	db, err = sql.Open("sqlite3", "linkler.db")
	if err != nil {
		fmt.Println("Veri tabanı acılamadı: ", err)
		return
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS linkler(
    kisaKod TEXT PRIMARY KEY,
    uzunLink TEXT
)`)
	if err != nil {
		fmt.Println("Tablo oluşturulamadı:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		uzunLink := r.URL.Query().Get("url")

		for kod, link := range urlMap {
			if link == uzunLink {
				fmt.Fprintln(w, kod)
				return
			}
		}

		kisaKod := kisaKodUret()
		urlMap[kisaKod] = uzunLink
		fmt.Fprintln(w, kisaKod)
	})

	http.HandleFunc("/r/", func(w http.ResponseWriter, r *http.Request) {
		kisaKod := r.URL.Path[3:]
		if kisaKod == "" {
			http.NotFound(w, r)
			return
		}
		uzunLink := urlMap[kisaKod]
		if len(uzunLink) > 0 && uzunLink[:4] != "http" {
			uzunLink = "https://" + uzunLink
		}
		http.Redirect(w, r, uzunLink, http.StatusFound)

	})

	http.HandleFunc("/qr/", func(w http.ResponseWriter, r *http.Request) {
		kisaKod := r.URL.Path[4:]
		if kisaKod == "" {
			http.NotFound(w, r)
			return
		}
		link := "http://localhost:8080/r/" + kisaKod
		png, err := qrcode.Encode(link, qrcode.Medium, 512)
		if err != nil {
			http.Error(w, "Qr kod oluşturulamadı ", 500)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(png)
	})

	http.ListenAndServe(":8080", nil)
}
