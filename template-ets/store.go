package main

type Store struct {
	sumberDaya []SumberDaya
	peminjaman []Peminjaman
	nextSDID   int
	nextPemID  int
}

func NewStore() *Store {
	return &Store{
		sumberDaya: []SumberDaya{},
		peminjaman: []Peminjaman{},
		nextSDID:   1,
		nextPemID:  1,
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
	for _, p := range s.peminjaman {
		if p.ID == id {
			return p, true
		}
	}
	return Peminjaman{}, false
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
