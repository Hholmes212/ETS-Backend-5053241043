package main

import (
	"log"
	"net/http"
)

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

	// B6: ubah dan hapus
	mux.HandleFunc("PUT /peminjaman/{id}", s.handleUbahPeminjaman)
	mux.HandleFunc("DELETE /peminjaman/{id}", s.handleHapusPeminjaman)

	// Persona D: paket peminjaman
	mux.HandleFunc("POST /paket-peminjaman", s.handleBuatPaket)
	mux.HandleFunc("DELETE /paket-peminjaman/{id}", s.handleHapusPaket)

	// B4: route yang tidak dikenal tetap dijawab JSON, bukan teks biasa
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tulisError(w, http.StatusNotFound, "endpoint tidak ditemukan")
	})

	return mux
}

func main() {
	log.Println("PinjamKampus berjalan di :8080")
	log.Fatal(http.ListenAndServe(":8080", NewHandler()))
}
