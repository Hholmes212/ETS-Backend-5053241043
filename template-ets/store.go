package main

import "sync"

type Store struct {
	mu          sync.Mutex
	sumberDaya  []SumberDaya
	peminjaman  []Peminjaman
	nextSDID    int
	nextPemID   int
	nextPaketID int
}

func NewStore() *Store {
	return &Store{
		sumberDaya:  []SumberDaya{},
		peminjaman:  []Peminjaman{},
		nextSDID:    1,
		nextPemID:   1,
		nextPaketID: 1,
	}
}

func (s *Store) tambahSumberDaya(sd SumberDaya) SumberDaya {
	sd.ID = s.nextSDID
	s.nextSDID++
	s.sumberDaya = append(s.sumberDaya, sd)
	return sd
}

func (s *Store) cariSumberDaya(id int) (SumberDaya, bool) {
	for _, sd := range s.sumberDaya {
		if sd.ID == id {
			return sd, true
		}
	}
	return SumberDaya{}, false
}

func (s *Store) tambahPeminjaman(p Peminjaman) Peminjaman {
	p.ID = s.nextPemID
	s.nextPemID++
	s.peminjaman = append(s.peminjaman, p)
	return p
}

func (s *Store) cariPeminjaman(id int) (Peminjaman, bool) {
	if i := s.indeksPeminjaman(id); i >= 0 {
		return s.peminjaman[i], true
	}
	return Peminjaman{}, false
}

func (s *Store) indeksPeminjaman(id int) int {
	for i, p := range s.peminjaman {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func (s *Store) filterPeminjaman(sdID *int, tanggal, status string) []Peminjaman {
	hasil := []Peminjaman{}
	for _, p := range s.peminjaman {
		if sdID != nil && p.SumberDayaID != *sdID {
			continue
		}
		if tanggal != "" && p.Tanggal != tanggal {
			continue
		}
		if status != "" && p.Status != status {
			continue
		}
		hasil = append(hasil, p)
	}
	return hasil
}

func (s *Store) adaBentrok(p Peminjaman, abaikanID int) bool {
	sd, ada := s.cariSumberDaya(p.SumberDayaID)
	if !ada {
		return false
	}
	var terpakai [9]int // indeks 1..8
	for _, q := range s.peminjaman {
		if q.ID == abaikanID || q.Status != statusAktif ||
			q.SumberDayaID != p.SumberDayaID || q.Tanggal != p.Tanggal {
			continue
		}
		for slot := q.SlotMulai; slot <= q.SlotSelesai; slot++ {
			terpakai[slot] += q.Jumlah
		}
	}
	for slot := p.SlotMulai; slot <= p.SlotSelesai; slot++ {
		if terpakai[slot]+p.Jumlah > sd.Stok {
			return true
		}
	}
	return false
}
