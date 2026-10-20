package main

import (
	"errors"
	"net/http"
	"slices"
)

func (s *Store) buatPaket(in paketInput) (paketHasil, int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if n := len(in.Item); n < 1 || n > 10 {
		return paketHasil{}, http.StatusBadRequest, errors.New("item harus berisi 1 sampai 10 elemen")
	}

	calon := make([]Peminjaman, 0, len(in.Item))
	dipakai := map[int]bool{}
	for _, it := range in.Item {
		p, err := s.validasiPeminjaman(peminjamanInput{
			SumberDayaID: it.SumberDayaID,
			Peminjam:     in.Peminjam,
			Tanggal:      in.Tanggal,
			SlotMulai:    in.SlotMulai,
			SlotSelesai:  in.SlotSelesai,
			Jumlah:       it.Jumlah,
		})
		if err != nil {
			return paketHasil{}, http.StatusBadRequest, err
		}
		if dipakai[p.SumberDayaID] {
			return paketHasil{}, http.StatusBadRequest, errors.New("satu sumber daya tidak boleh muncul dua kali dalam satu paket")
		}
		dipakai[p.SumberDayaID] = true
		calon = append(calon, p)
	}

	for _, p := range calon {
		if s.adaBentrok(p, 0) {
			return paketHasil{}, http.StatusConflict, errBentrok
		}
	}

	paketID := s.nextPaketID
	s.nextPaketID++
	hasil := paketHasil{PaketID: paketID, Peminjaman: make([]Peminjaman, 0, len(calon))}
	for _, p := range calon {
		p.PaketID = paketID
		hasil.Peminjaman = append(hasil.Peminjaman, s.tambahPeminjaman(p))
	}
	return hasil, http.StatusCreated, nil
}

// hapusPaket menghapus semua peminjaman milik paket. false bila paket tidak ada.
func (s *Store) hapusPaket(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	ada := slices.ContainsFunc(s.peminjaman, func(p Peminjaman) bool { return p.PaketID == id })
	if !ada {
		return false
	}
	s.peminjaman = slices.DeleteFunc(s.peminjaman, func(p Peminjaman) bool { return p.PaketID == id })
	return true
}

func (s *Store) handleBuatPaket(w http.ResponseWriter, r *http.Request) {
	var in paketInput
	if !bacaBody(r, &in) {
		tulisError(w, http.StatusBadRequest, "body harus berupa JSON yang valid")
		return
	}
	hasil, kode, err := s.buatPaket(in)
	if err != nil {
		tulisError(w, kode, err.Error())
		return
	}
	tulisJSON(w, kode, hasil)
}

func (s *Store) handleHapusPaket(w http.ResponseWriter, r *http.Request) {
	id, ok := bacaID(r)
	if !ok {
		tulisError(w, http.StatusBadRequest, "id harus berupa angka")
		return
	}
	if !s.hapusPaket(id) {
		tulisError(w, http.StatusNotFound, "paket tidak ditemukan")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
