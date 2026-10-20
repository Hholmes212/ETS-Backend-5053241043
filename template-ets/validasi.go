package main

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func validasiTeks(nilai *string, label string, maks int) (string, error) {
	if nilai == nil {
		return "", errors.New(label + " wajib diisi")
	}
	t := strings.TrimSpace(*nilai)
	if t == "" {
		return "", errors.New(label + " tidak boleh kosong")
	}
	if utf8.RuneCountInString(t) > maks {
		return "", errors.New(label + " terlalu panjang (maksimal " + strconv.Itoa(maks) + " karakter)")
	}
	return t, nil
}

func validasiTanggal(nilai *string) (string, error) {
	if nilai == nil {
		return "", errors.New("tanggal wajib diisi")
	}
	t := strings.TrimSpace(*nilai)
	if _, err := time.Parse("2006-01-02", t); err != nil {
		return "", errors.New("tanggal harus berformat YYYY-MM-DD dan benar-benar ada")
	}
	return t, nil
}

func validasiSumberDaya(in sumberDayaInput) (SumberDaya, error) {
	nama, err := validasiTeks(in.Nama, "nama", 50)
	if err != nil {
		return SumberDaya{}, err
	}
	if in.Jenis == nil {
		return SumberDaya{}, errors.New("jenis wajib diisi")
	}
	jenis := strings.TrimSpace(*in.Jenis)
	if jenis != "ruang" && jenis != "alat" {
		return SumberDaya{}, errors.New(`jenis harus "ruang" atau "alat"`)
	}
	if in.Stok == nil {
		return SumberDaya{}, errors.New("stok wajib diisi")
	}
	if *in.Stok < 1 {
		return SumberDaya{}, errors.New("stok minimal 1")
	}
	return SumberDaya{Nama: nama, Jenis: jenis, Stok: *in.Stok}, nil
}

func (s *Store) validasiPeminjaman(in peminjamanInput) (Peminjaman, error) {
	var p Peminjaman

	if in.SumberDayaID == nil {
		return p, errors.New("sumber_daya_id wajib diisi")
	}
	sd, ada := s.cariSumberDaya(*in.SumberDayaID)
	if !ada {
		return p, errors.New("sumber_daya_id tidak ditemukan")
	}

	peminjam, err := validasiTeks(in.Peminjam, "peminjam", 50)
	if err != nil {
		return p, err
	}
	tanggal, err := validasiTanggal(in.Tanggal)
	if err != nil {
		return p, err
	}

	if in.SlotMulai == nil || in.SlotSelesai == nil {
		return p, errors.New("slot_mulai dan slot_selesai wajib diisi")
	}
	if *in.SlotMulai < 1 || *in.SlotMulai > 8 || *in.SlotSelesai < 1 || *in.SlotSelesai > 8 {
		return p, errors.New("slot harus antara 1 sampai 8")
	}
	if *in.SlotMulai > *in.SlotSelesai {
		return p, errors.New("slot_mulai tidak boleh lebih besar dari slot_selesai")
	}

	if in.Jumlah == nil {
		return p, errors.New("jumlah wajib diisi")
	}
	if *in.Jumlah < 1 || *in.Jumlah > sd.Stok {
		return p, errors.New("jumlah harus antara 1 dan stok sumber daya")
	}

	return Peminjaman{
		SumberDayaID: sd.ID,
		Peminjam:     peminjam,
		Tanggal:      tanggal,
		SlotMulai:    *in.SlotMulai,
		SlotSelesai:  *in.SlotSelesai,
		Jumlah:       *in.Jumlah,
		Status:       "aktif",
	}, nil
}
