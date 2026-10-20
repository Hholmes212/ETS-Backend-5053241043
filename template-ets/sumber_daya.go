package main

import "net/http"

func (s *Store) handleBuatSumberDaya(w http.ResponseWriter, r *http.Request) {
	var in sumberDayaInput
	if !bacaBody(r, &in) {
		tulisError(w, http.StatusBadRequest, "body harus berupa JSON yang valid")
		return
	}
	sd, err := validasiSumberDaya(in)
	if err != nil {
		tulisError(w, http.StatusBadRequest, err.Error())
		return
	}
	s.mu.Lock()
	baru := s.tambahSumberDaya(sd)
	s.mu.Unlock()
	tulisJSON(w, http.StatusCreated, baru)
}

func (s *Store) handleDaftarSumberDaya(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	daftar := append([]SumberDaya{}, s.sumberDaya...)
	s.mu.Unlock()
	tulisJSON(w, http.StatusOK, daftar)
}

func (s *Store) handleDetailSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, ok := bacaID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	s.mu.Lock()
	sd, ada := s.cariSumberDaya(id)
	s.mu.Unlock()
	if !ada {
		tulisError(w, http.StatusNotFound, "sumber daya tidak ditemukan")
		return
	}
	tulisJSON(w, http.StatusOK, sd)
}
