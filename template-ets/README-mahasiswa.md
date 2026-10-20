# PinjamKampus API

- Nama: Muhammad Alfaraldi Raihan 
- NRP: 5053241043
- Persona: D
- Cara menjalankan: `go run .` lalu buka http://localhost:8080

## Fitur
- [x] B1 Sumber daya
- [x] B2 Buat peminjaman
- [x] B3 Lihat peminjaman (daftar, filter, detail)
- [x] B4 Validasi dan error JSON
- [x] B5 Deteksi bentrok (409)
- [ ] B6 Ubah dan hapus
- [ ] B7 Aman untuk permintaan bersamaan
- [ ] Fitur khas persona
- [ ] Change request

## Yang belum selesai
- B6 (PUT dan DELETE /peminjaman/{id}): belum dikerjakan, endpoint-nya belum ada.
- B7 (aman untuk permintaan bersamaan): belum ada mutex, jadi request yang datang bersamaan belum dijamin aman.
- Fitur khas persona D (POST dan DELETE /paket-peminjaman): belum dikerjakan.
- Change request A (GET /ringkasan): belum dikerjakan.