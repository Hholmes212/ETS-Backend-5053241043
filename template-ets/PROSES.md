# PROSES

## 1. Persona dan masalahnya
Panitia acara kampus yang butuh beberapa sumber daya sekaligus untuk satu acara, misal aula, proyektor, sound system. Masalahnya, kalau dipesan satu per satu, bisa saja aulanya sudah terpesan tetapi proyektornya ternyata sudah dipinjam orang lain, sehingga acara batal dan aulanya terlanjur terpakai. Fitur peminjaman paket menyelesaikan ini dengan memesan semua item dalam satu permintaan: semuanya berhasil, atau tidak ada satu pun yang dibuat.

## 2. Rancangan struktur data
type SumberDaya struct {
    ID    int    `json:"id"`
    Nama  string `json:"nama"`
    Jenis string `json:"jenis"`
    Stok  int    `json:"stok"`
}

type Peminjaman struct {
    ID           int    `json:"id"`
    SumberDayaID int    `json:"sumber_daya_id"`
    Peminjam     string `json:"peminjam"`
    Tanggal      string `json:"tanggal"`
    SlotMulai    int    `json:"slot_mulai"`
    SlotSelesai  int    `json:"slot_selesai"`
    Jumlah       int    `json:"jumlah"`
    Status       string `json:"status"`
    PaketID      int    `json:"paket_id"` // 0 kalau dibuat lewat POST /peminjaman biasa
}

type Store struct {
    mu           sync.Mutex
    sumberDaya   []SumberDaya
    peminjaman   []Peminjaman
    nextSDID     int // id berikutnya untuk sumber daya
    nextPinjamID int // id berikutnya untuk peminjaman
    nextPaketID  int // paket_id berikutnya, hanya naik kalau paket berhasil
}

Alasan:
Semua data saya taruh di satu struct Store supaya cukup satu mutex yang
menjaga semuanya. Peminjaman dan sumber daya saya simpan di slice. Id-nya
bertambah satu-satu dan datanya selalu ditambah di belakang, jadi daftar
otomatis urut dan tidak perlu di-sort lagi.

Paket tidak saya buatkan struct sendiri. Setiap item dalam paket tetap
satu peminjaman biasa, bedanya cuma diberi paket_id yang sama. Dengan
begitu cek bentrok yang sudah ada langsung bisa dipakai, dan peminjaman
dari paket ikut muncul di GET /peminjaman. Untuk cek bentrok saya buat
satu fungsi yang dipakai bersama oleh POST, PUT, dan paket.

Perubahan dari rancangan awal: (isi kalau ada, atau "tidak ada")


## 3. Satu keputusan desain
(Satu keputusan yang kamu ambil, pilihan lain yang kamu pertimbangkan, dan alasanmu memilih yang ini. Contoh: kenapa cek bentrok per slot, bukan per rentang.)

## 4. Bug atau kendala (minimal dua)
### Bug 1
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

### Bug 2
- Gejala:
- Cara menemukan:
- Penyebab dan solusi:

## 5. Change request
(Apa yang diminta, bagian kode mana yang kamu ubah, dan satu keputusan desain yang kamu ambil.)
