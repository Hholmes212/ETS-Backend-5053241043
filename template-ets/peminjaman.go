package main

import (
	"net/http"
	"strconv"
	"strings"
)

func (s *Store) handleBuatPeminjaman(w http.ResponseWriter, r *http.Request) {
	var in peminjamanInput
	if !bacaBody(r, &in) {
		tulisError(w, http.StatusBadRequest, "body harus berupa JSON yang valid")
		return
	}
	p, err := s.validasiPeminjaman(in)
	if err != nil {
		tulisError(w, http.StatusBadRequest, err.Error())
		return
	}
	tulisJSON(w, http.StatusCreated, s.tambahPeminjaman(p))
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

	tulisJSON(w, http.StatusOK, s.filterPeminjaman(sdID, tanggal, status))
}

func (s *Store) handleDetailPeminjaman(w http.ResponseWriter, r *http.Request) {
	id, ok := bacaID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	p, ada := s.cariPeminjaman(id)
	if !ada {
		tulisError(w, http.StatusNotFound, "peminjaman tidak ditemukan")
		return
	}
	tulisJSON(w, http.StatusOK, p)
}
