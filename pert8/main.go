package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	http.HandleFunc("/api/mahasiswa", user)
	fmt.Println("web Berjalan dengan port 1491")
	http.ListenAndServe(":1491", nil)
}

type lepkom struct {
	Nama   string `json:"nama_mahasiswa"`
	Npm    string `json:"npm_mahasiswa"`
	Kursus string `json:"kursus_mahasiswa"`
	Foto   string `json:"foto_mahasiswa"`
}

var data_mahasiswa = []lepkom{
	{
		Nama:   "Fitri Nur Afifah",
		Npm:    "10121491",
		Kursus: "Golang For Beginer",
		Foto:   "img/gambar1.jpg",
	},
	{
		Nama:   "Balya Najla",
		Npm:    "10121492",
		Kursus: "Cisco For Beginner",
		Foto:   "img/gambar1.jpg",
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Method == http.MethodGet {
		result, err := json.Marshal(data_mahasiswa)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(result)
		return
	}
}
