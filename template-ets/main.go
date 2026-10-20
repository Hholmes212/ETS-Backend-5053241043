package main

import (
	"log"
	"net/http"
)

// NewHandler membuat penyimpanan baru yang kosong dan memasang semua route.
func NewHandler() http.Handler {
	s := NewStore()
	mux := http.NewServeMux()

	// B1: sumber daya
	mux.HandleFunc("POST /sumber-daya", s.handleBuatSumberDaya)
	mux.HandleFunc("GET /sumber-daya", s.handleDaftarSumberDaya)
	mux.HandleFunc("GET /sumber-daya/{id}", s.handleDetailSumberDaya)

	// B2 + B3: peminjaman
	mux.HandleFunc("POST /peminjaman", s.handleBuatPeminjaman)
	mux.HandleFunc("GET /peminjaman", s.handleDaftarPeminjaman)
	mux.HandleFunc("GET /peminjaman/{id}", s.handleDetailPeminjaman)

	return mux
}

func main() {
	log.Println("PinjamKampus berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", NewHandler()))
}
