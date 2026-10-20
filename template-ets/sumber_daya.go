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
	tulisJSON(w, http.StatusCreated, s.tambahSumberDaya(sd))
}

func (s *Store) handleDaftarSumberDaya(w http.ResponseWriter, r *http.Request) {
	tulisJSON(w, http.StatusOK, s.sumberDaya)
}

func (s *Store) handleDetailSumberDaya(w http.ResponseWriter, r *http.Request) {
	id, ok := bacaID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	sd, ada := s.cariSumberDaya(id)
	if !ada {
		tulisError(w, http.StatusNotFound, "sumber daya tidak ditemukan")
		return
	}
	tulisJSON(w, http.StatusOK, sd)
}
