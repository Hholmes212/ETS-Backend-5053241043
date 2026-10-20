package main

import (
	"errors"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

var errBentrok = errors.New("kapasitas sumber daya pada slot tersebut sudah penuh")

func (s *Store) buatPeminjaman(in peminjamanInput) (Peminjaman, int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, err := s.validasiPeminjaman(in)
	if err != nil {
		return Peminjaman{}, http.StatusBadRequest, err
	}
	if s.adaBentrok(p, 0) {
		return Peminjaman{}, http.StatusConflict, errBentrok
	}
	return s.tambahPeminjaman(p), http.StatusCreated, nil
}

func (s *Store) ubahPeminjaman(id int, in peminjamanInput) (Peminjaman, int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	idx := s.indeksPeminjaman(id)
	if idx < 0 {
		return Peminjaman{}, http.StatusNotFound, errors.New("peminjaman tidak ditemukan")
	}
	lama := s.peminjaman[idx]

	p, err := s.validasiPeminjaman(in)
	if err != nil {
		return Peminjaman{}, http.StatusBadRequest, err
	}
	p.ID = lama.ID
	p.Status = lama.Status
	p.PaketID = lama.PaketID

	if s.adaBentrok(p, id) {
		return Peminjaman{}, http.StatusConflict, errBentrok
	}
	s.peminjaman[idx] = p
	return p, http.StatusOK, nil
}

func (s *Store) hapusPeminjaman(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	idx := s.indeksPeminjaman(id)
	if idx < 0 {
		return false
	}
	s.peminjaman = slices.Delete(s.peminjaman, idx, idx+1)
	return true
}

func (s *Store) ambilPeminjaman(id int) (Peminjaman, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.cariPeminjaman(id)
}

func (s *Store) handleBuatPeminjaman(w http.ResponseWriter, r *http.Request) {
	var in peminjamanInput
	if !bacaBody(r, &in) {
		tulisError(w, http.StatusBadRequest, "body harus berupa JSON yang valid")
		return
	}
	p, kode, err := s.buatPeminjaman(in)
	if err != nil {
		tulisError(w, kode, err.Error())
		return
	}
	tulisJSON(w, kode, p)
}

func (s *Store) handleDaftarPeminjaman(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	var sdID *int
	if v := q.Get("sumber_daya_id"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			tulisError(w, http.StatusBadRequest, "sumber_daya_id harus berupa angka")
			return
		}
		sdID = &n
	}
	tanggal := strings.TrimSpace(q.Get("tanggal"))
	status := strings.TrimSpace(q.Get("status"))

	s.mu.Lock()
	hasil := s.filterPeminjaman(sdID, tanggal, status)
	s.mu.Unlock()
	tulisJSON(w, http.StatusOK, hasil)
}

func (s *Store) handleDetailPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, ok := bacaID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	p, ada := s.ambilPeminjaman(id)
	if !ada {
		tulisError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
		return
	}
	tulisJSON(w, http.StatusOK, p)
}

func (s *Store) handleUbahPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, ok := bacaID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	if _, ada := s.ambilPeminjaman(id); !ada {
		tulisError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
		return
	}
	var in peminjamanInput
	if !bacaBody(r, &in) {
		tulisError(w, http.StatusBadRequest, "body harus berupa JSON yang valid")
		return
	}
	p, kode, err := s.ubahPeminjaman(id, in)
	if err != nil {
		tulisError(w, kode, err.Error())
		return
	}
	tulisJSON(w, kode, p)
}

func (s *Store) handleHapusPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, ok := bacaID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	if !s.hapusPeminjaman(id) {
		tulisError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
