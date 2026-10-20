# AI-LOG

Catat setiap penggunaan AI yang menghasilkan kode. Jika tidak memakai AI sama sekali, tulis satu baris: **tidak memakai AI**.

Log yang jujur tidak mengurangi nilai, sebanyak apa pun AI dipakai. Yang mengurangi nilai adalah log yang tidak sesuai dengan kode, atau isi log yang tidak bisa dijelaskan.

## B1-B3: sumber daya dan peminjaman dasar
- Prompt: Minta bantuan menjalankan project (error `go run` dan `go mod`) dan membuat daftar request Postman untuk menguji endpoint B1-B3.
- Dipakai: Daftar request Postman untuk menguji POST/GET sumber daya, POST peminjaman, serta GET daftar, filter, dan detail peminjaman. Kode B1-B3 (`sumber_daya.go`, `peminjaman.go`, `store.go`, `validasi.go`) saya tulis sendiri.
- Diubah: Tidak ada yang diubah dari hasil AI, karena AI tidak menghasilkan kode untuk bagian ini.

## B4-B5: validasi dan deteksi bentrok
- Prompt: Minta dilengkapi validasi error JSON dan deteksi bentrok slot berdasarkan stok.
- Dipakai: `adaBentrok()` di `store.go`, pemanggilannya di `handleBuatPeminjaman` (`peminjaman.go`), dan route `"/"` untuk 404 JSON di `main.go`.
- Diubah: Hasil AI dipecah menjadi versi bertahap sampai B5 supaya sesuai progres saya. Fungsi `adaBentrok` menghitung pemakaian unit per slot, bukan hanya ada atau tidaknya peminjaman lain.
