<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>

> ✅ **Projek ini telah selesai di kembangkan ulang dengan menulis ulang kode menggunakan arsitektur [_hexagonal_](https://blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example/), serta menggunakan framework [Fiber](https://gofiber.io/).**
> **Jika ingin melihat kode sebelumnya bisa melihat [commit ini](https://github.com/afrizal423/Golang-Perpustakaan-Restful-API/tree/d451e99fd6cdb506accd16b969ffc901dfc81dac).**

# Sistem Informasi Perpustakaan Berbasis Rest API

Project sederhana membuat Sistem Informasi Perpustakaan berbasis Restful API menggunakan bahasa pemrograman Go dengan database MySql, mengimplementasikan arsitektur hexagonal, autentikasi JWT, serta modul CRUD untuk buku, peminjaman, dan denda. Alur projectnya kurang lebihnya seperti [ini](https://core.ac.uk/download/pdf/12347733.pdf). Password hash menggunakan argon2id

## ToDo List Pengerjaan Project :pushpin:

- [x] Create Migration table
- [x] Create Seed data akun admin
- [x] Login With JWT
- [x] CRUDF Buku
  - [x] CRUDF Jenis buku
  - [x] CRUDF Penulis buku
  - [x] CRUDF Penerbit buku
- [x] CRUDF peminjaman
  - [x] CRUDF Detail pinjam
  - [x] CRUDF Denda

## Project structure

```
.
├── cmd/
│   └── main.go                  # Titik masuk aplikasi
├── internal/
│   ├── application/             # Lapisan aplikasi (use case, layanan)
│   │   ├── usecase/
│   │   │   ├── auth/            # Layanan autentikasi
│   │   │   ├── buku/            # Layanan terkait buku
│   │   │   ├── denda/           # Layanan terkait denda
│   │   │   └── peminjaman/      # Layanan terkait peminjaman
│   │   └── mocks/               # Implementasi mock untuk pengujian
│   ├── core/
│   │   ├── domain/              # Entitas bisnis inti (model)
│   │   └── ports/               # Antarmuka untuk layanan aplikasi dan infrastruktur
│   └── infrastructure/          # Lapisan infrastruktur (adapter)
│       ├── config/              # Konfigurasi aplikasi
│       ├── database/            # Migrasi dan seeding database
│       ├── http/                # Server HTTP, handler, middleware
│       │   ├── middleware/      # Middleware untuk autentikasi, dll.
│       │   └── v1/              # Handler API versi 1
│       │       ├── auth/
│       │       ├── buku/
│       │       ├── denda/
│       │       └── peminjaman/
│       ├── repository/
│       │   └── mysql/           # Implementasi repositori khusus MySQL
│       │       ├── auth/
│       │       ├── buku/
│       │       ├── denda/
│       │       └── peminjaman/
│       └── utility/             # Utilitas umum (hashing, JWT, validasi)
└── web/
    └── storage/                 # File statis
```

## Features

- **Arsitektur Hexagonal:** Struktur proyek mengikuti prinsip-prinsip arsitektur heksagonal untuk pemisahan kepentingan dan pemeliharaan yang lebih baik.
- **Autentikasi Pengguna:** Login pengguna yang aman menggunakan JWT (JSON Web Tokens) untuk autentikasi.
- **Modul Buku:**
  - Full CRUD (Create, Read, Update, Delete) untuk Jenis Buku.
  - Full CRUD untuk Penulis Buku.
  - Full CRUD untuk Penerbit Buku.
- **Modul Peminjaman:**
  - Full CRUD untuk Peminjaman (catatan peminjaman).
  - Tampilan detail untuk catatan peminjaman individu, termasuk detail buku terkait.
- **Modul Denda:**
  - Full CRUD untuk Denda (catatan denda).
- **Kemampuan Pengujian:**
  - Semua layanan inti dirancang untuk kemampuan pengujian menggunakan `mockery` untuk pembuatan mock.

## Features

- **Arsitektur Hexagonal:** Struktur proyek mengikuti prinsip-prinsip arsitektur heksagonal untuk pemisahan kepentingan dan pemeliharaan yang lebih baik.
- **Autentikasi Pengguna:** Login pengguna yang aman menggunakan JWT (JSON Web Tokens) untuk autentikasi.
- **Modul Buku:**
  - Full CRUD (Create, Read, Update, Delete) untuk Jenis Buku.
  - Full CRUD untuk Penulis Buku.
  - Full CRUD untuk Penerbit Buku.
- **Modul Peminjaman:**
  - Full CRUD untuk Peminjaman (catatan peminjaman).
  - Tampilan detail untuk catatan peminjaman individu, termasuk detail buku terkait.
- **Modul Denda:**
  - Full CRUD untuk Denda (catatan denda).
- **Kemampuan Pengujian:**
  - Semua layanan inti dirancang untuk kemampuan pengujian menggunakan `mockery` untuk pembuatan mock.

## Dokumentasi API

#### Umum
- **Base URL:** `http://localhost:8001/api/v1`
- **Autentikasi:** Semua endpoint yang memerlukan autentikasi membutuhkan token JWT di header `Authorization: Bearer <token>`.

#### Modul Autentikasi

##### 1. Login Pengguna (Pegawai)
- **Endpoint:** `POST /login`
- **Deskripsi:** Mengautentikasi pengguna pegawai dan mengembalikan token JWT.
- **Body Permintaan (Request Body):**
  ```json
  {
    "username": "admin",
    "password": "admin_password"
  }
  ```
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Login successful",
    "data": {
      "username": "admin",
      "token": "eyJhbGciOiJIUzI1Ni...",
      "refresh_token": "eyJhbGciOiJIUzI1Ni..."
    }
  }
  ```
- **Respons Gagal (401 Unauthorized):**
  ```json
  {
    "error": true,
    "msg": "invalid credentials"
  }
  ```

#### Modul Buku (Publik)

##### 1. Mendapatkan Semua Buku
- **Endpoint:** `GET /api/v1/buku`
- **Deskripsi:** Mendapatkan daftar semua buku yang tersedia.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get all buku",
    "data": [
      {
        "id_buku": "01H019WHW4A45KH65AH9XW89PC",
        "isbn": "INIISBNKE1",
        "id_kategori_buku": "id_kategori_sample",
        "judul_buku": "Ini buku ke 1",
        "id_penulis_buku": "id_penulis_sample",
        "id_penerbit_buku": "id_penerbit_sample",
        "tahun_terbit": "2020",
        "stok_buku": 10,
        "rak_buku": "A01",
        "deskripsi_buku": "Ini deskripsi kali ya, bagian buku 1",
        "gambar_buku": null,
        "kondisi_buku": null,
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```

##### 2. Mendapatkan Detail Buku
- **Endpoint:** `GET /api/v1/buku/:id`
- **Deskripsi:** Mendapatkan detail lengkap buku berdasarkan ID-nya.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get buku by ID",
    "data": {
      "id_buku": "01H019WHW4A45KH65AH9XW89PC",
      "isbn": "INIISBNKE1",
      "id_kategori_buku": "id_kategori_sample",
      "judul_buku": "Ini buku ke 1",
      "id_penulis_buku": "id_penulis_sample",
      "id_penerbit_buku": "id_penerbit_sample",
      "tahun_terbit": "2020",
      "stok_buku": 10,
      "rak_buku": "A01",
      "deskripsi_buku": "Ini deskripsi kali ya, bagian buku 1",
      "gambar_buku": null,
      "kondisi_buku": null,
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

#### Modul Buku (Membutuhkan Autentikasi)

##### 1. Jenis Buku
- **Endpoint:** `GET /admin/buku/jenbuk`
- **Deskripsi:** Mendapatkan semua jenis buku. Dapat difilter dengan query `?q=keyword`.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get data buku",
    "data": [
      {
        "id": "01H019WHW4A45KH65AH9XW89PC",
        "jenis_buku": "History",
        "deskripsi": "Ini adalah jenis buku history",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```
- **Endpoint:** `GET /admin/buku/jenbuk/:id`
- **Deskripsi:** Mendapatkan jenis buku berdasarkan ID.
- **Respons Sukses (200 OK):**
  ```json
  {
    "status": "Success get data buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "jenis_buku": "History",
      "deskripsi": "Ini adalah jenis buku history",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `POST /admin/buku/jenbuk/create`
- **Deskripsi:** Membuat jenis buku baru.
- **Body Permintaan (Request Body):**
  ```json
  {
    "jenis_buku": "Fiksi Ilmiah",
    "deskripsi": "Buku-buku tentang sains fiksi"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "msg": "Success added jenis buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "jenis_buku": "Fiksi Ilmiah",
      "deskripsi": "Buku-buku tentang sains fiksi",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `PUT /admin/buku/jenbuk/update`
- **Deskripsi:** Memperbarui jenis buku yang ada.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id": "01H019WHW4A45KH65AH9XW89PC",
    "jenis_buku": "Fiksi Ilmiah Terbaru",
    "deskripsi": "Kategori untuk buku-buku sains fiksi terbaru"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "msg": "Success update jenis buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "jenis_buku": "Fiksi Ilmiah Terbaru",
      "deskripsi": "Kategori untuk buku-buku sains fiksi terbaru",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `DELETE /admin/buku/jenbuk/delete`
- **Deskripsi:** Menghapus jenis buku berdasarkan ID.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id": "01H019WHW4A45KH65AH9XW89PC"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "msg": "Success delete jenis buku"
  }
  ```

##### 2. Penulis Buku
- **Endpoint:** `GET /admin/buku/author`
- **Deskripsi:** Mendapatkan semua penulis buku. Dapat difilter dengan query `?q=keyword`.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get data Penulis buku",
    "data": [
      {
        "id": "01H019WHW4A45KH65AH9XW89PC",
        "penulis_buku": "Penulis 1",
        "alamat": "Kota 1",
        "email_penulis": "test1@test.com",
        "deskripsi": "",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```
- **Endpoint:** `GET /admin/buku/author/:id`
- **Deskripsi:** Mendapatkan penulis buku berdasarkan ID.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get data Penulis buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "penulis_buku": "Penulis 1",
      "alamat": "Kota 1",
      "email_penulis": "test1@test.com",
      "deskripsi": "",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `POST /admin/buku/author/create`
- **Deskripsi:** Membuat penulis buku baru.
- **Body Permintaan (Request Body):**
  ```json
  {
    "penulis_buku": "Penulis Baru",
    "alamat_penulis": "Alamat Penulis",
    "email_penulis": "baru@email.com",
    "deskripsi": "Deskripsi penulis baru"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "error": false,
    "msg": "Success insert data Penulis buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "penulis_buku": "Penulis Baru",
      "alamat": "Alamat Penulis",
      "email_penulis": "baru@email.com",
      "deskripsi": "Deskripsi penulis baru",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `PUT /admin/buku/author/update`
- **Deskripsi:** Memperbarui penulis buku yang ada.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id": "01H019WHW4A45KH65AH9XW89PC",
    "penulis_buku": "Penulis Diperbarui",
    "alamat_penulis": "Alamat Baru",
    "email_penulis": "update@email.com",
    "deskripsi": "Deskripsi penulis diperbarui"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "error": false,
    "msg": "Success update data Penulis buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "penulis_buku": "Penulis Diperbarui",
      "alamat": "Alamat Baru",
      "email_penulis": "update@email.com",
      "deskripsi": "Deskripsi penulis diperbarui",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `DELETE /admin/buku/author/delete`
- **Deskripsi:** Menghapus penulis buku berdasarkan ID.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id": "01H019WHW4A45KH65AH9XW89PC"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "msg": "Success delete Penulis buku"
  }
  ```

##### 3. Penerbit Buku
- **Endpoint:** `GET /admin/buku/penbuk`
- **Deskripsi:** Mendapatkan semua penerbit buku. Dapat difilter dengan query `?q=keyword`.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "status": "Success get data penerbit buku",
    "data": [
      {
        "id": "01H019WHW4A45KH65AH9XW89PC",
        "penerbit_buku": "Penerbit 1",
        "alamat_penerbit": "Kota 1",
        "telp_penerbit": null,
        "email_penerbit": "publisher1@test.com",
        "deskripsi": null,
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```
- **Endpoint:** `GET /admin/buku/penbuk/:id`
- **Deskripsi:** Mendapatkan penerbit buku berdasarkan ID.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "status": "Success get data penerbit buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "penerbit_buku": "Penerbit 1",
      "alamat_penerbit": "Kota 1",
      "telp_penerbit": null,
      "email_penerbit": "publisher1@test.com",
      "deskripsi": null,
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `POST /admin/buku/penbuk/create`
- **Deskripsi:** Membuat penerbit buku baru.
- **Body Permintaan (Request Body):**
  ```json
  {
    "penerbit_buku": "Penerbit Baru",
    "alamat_penerbit": "Alamat Penerbit Baru",
    "telp_penerbit": "081234567890",
    "email_penerbit": "penerbit@example.com",
    "deskripsi": "Deskripsi penerbit baru"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "error": false,
    "status": "Success insert data penerbit buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "penerbit_buku": "Penerbit Baru",
      "alamat_penerbit": "Alamat Penerbit Baru",
      "telp_penerbit": "081234567890",
      "email_penerbit": "penerbit@example.com",
      "deskripsi": "Deskripsi penerbit baru",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `PUT /admin/buku/penbuk/update`
- **Deskripsi:** Memperbarui penerbit buku yang ada.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id": "01H019WHW4A45KH65AH9XW89PC",
    "penerbit_buku": "Penerbit Diperbarui",
    "alamat_penerbit": "Alamat Baru",
    "telp_penerbit": "089876543210",
    "email_penerbit": "updated@example.com",
    "deskripsi": "Deskripsi diperbarui"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "error": false,
    "status": "Success update data penerbit buku",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "penerbit_buku": "Penerbit Diperbarui",
      "alamat_penerbit": "Alamat Baru",
      "telp_penerbit": "089876543210",
      "email_penerbit": "updated@example.com",
      "deskripsi": "Deskripsi diperbarui",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```
- **Endpoint:** `DELETE /admin/buku/penbuk/delete`
- **Deskripsi:** Menghapus penerbit buku berdasarkan ID.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id": "01H019WHW4A45KH65AH9XW89PC"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "msg": "Success delete penerbit buku"
  }
  ```

#### Modul Peminjaman (Membutuhkan Autentikasi)

##### 1. Mendapatkan Semua Peminjaman
- **Endpoint:** `GET /admin/peminjaman`
- **Deskripsi:** Mendapatkan semua catatan peminjaman.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get all peminjaman",
    "data": [
      {
        "id": "01H019WHW4A45KH65AH9XW89PC",
        "id_anggota": "id_anggota_sample",
        "tgl_pinjam": "2023-01-01T10:00:00Z",
        "tgl_hrs_kembali": "2023-01-08T10:00:00Z",
        "jaminan": "KTP",
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```

##### 2. Mendapatkan Peminjaman Berdasarkan ID
- **Endpoint:** `GET /admin/peminjaman/:id`
- **Deskripsi:** Mendapatkan catatan peminjaman berdasarkan ID.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get peminjaman by ID",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "id_anggota": "id_anggota_sample",
      "tgl_pinjam": "2023-01-01T10:00:00Z",
      "tgl_hrs_kembali": "2023-01-08T10:00:00Z",
      "jaminan": "KTP",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

##### 3. Mendapatkan Detail Peminjaman
- **Endpoint:** `GET /admin/peminjaman/detail/:id`
- **Deskripsi:** Mendapatkan catatan peminjaman lengkap dengan detail anggota dan buku yang dipinjam.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get detail peminjaman",
    "data": {
      "id": "01H019WHW4A45KH65AH9XW89PC",
      "anggota": {
        "id_anggota": "id_anggota_sample",
        "nama": "Nama Anggota"
      },
      "tgl_pinjam": "2023-01-01T10:00:00Z",
      "tgl_hrs_kembali": "2023-01-08T10:00:00Z",
      "jaminan": "KTP",
      "details": [
        {
          "id_detailpinjam": "id_detail_pinjam_sample",
          "id_buku": "id_buku_sample",
          "kondisi": "Baik"
        }
      ],
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

##### 4. Membuat Peminjaman Baru
- **Endpoint:** `POST /admin/peminjaman/create`
- **Deskripsi:** Membuat catatan peminjaman baru.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id_anggota": "id_anggota_sample",
    "tgl_pinjam": "2023-01-01T10:00:00Z",
    "tgl_hrs_kembali": "2023-01-08T10:00:00Z",
    "jaminan": "KTP"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "error": false,
    "msg": "Peminjaman created successfully"
  }
  ```

##### 5. Memperbarui Peminjaman
- **Endpoint:** `PUT /admin/peminjaman/update`
- **Deskripsi:** Memperbarui catatan peminjaman yang ada.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id_peminjaman": "01H019WHW4A45KH65AH9XW89PC",
    "id_anggota": "id_anggota_sample",
    "tgl_pinjam": "2023-01-01T10:00:00Z",
    "tgl_hrs_kembali": "2023-01-09T10:00:00Z",
    "jaminan": "SIM"
  }
  ```
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Peminjaman updated successfully"
  }
  ```

##### 6. Menghapus Peminjaman
- **Endpoint:** `DELETE /admin/peminjaman/delete`
- **Deskripsi:** Menghapus catatan peminjaman berdasarkan ID.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id_peminjaman": "01H019WHW4A45KH65AH9XW89PC"
  }
  ```
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Peminjaman deleted successfully"
  }
  ```

#### Modul Denda (Membutuhkan Autentikasi)

##### 1. Mendapatkan Semua Denda
- **Endpoint:** `GET /admin/denda`
- **Deskripsi:** Mendapatkan semua catatan denda.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get all denda",
    "data": [
      {
        "id_denda": "01H019WHW4A45KH65AH9XW89PC",
        "jumlah_denda": 5000,
        "tgl_pinjam": "2023-01-01T10:00:00Z",
        "tgl_hrs_kembali": "2023-01-08T10:00:00Z",
        "tgl_kembali": "2023-01-10T10:00:00Z",
        "id_peminjaman": "id_peminjaman_sample",
        "id_anggota": "id_anggota_sample",
        "created_at": "2023-01-01T00:00:00Z",
        "updated_at": "2023-01-01T00:00:00Z"
      }
    ]
  }
  ```

##### 2. Mendapatkan Denda Berdasarkan ID
- **Endpoint:** `GET /admin/denda/:id`
- **Deskripsi:** Mendapatkan catatan denda berdasarkan ID.
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Success get denda by ID",
    "data": {
      "id_denda": "01H019WHW4A45KH65AH9XW89PC",
      "jumlah_denda": 5000,
      "tgl_pinjam": "2023-01-01T10:00:00Z",
      "tgl_hrs_kembali": "2023-01-08T10:00:00Z",
      "tgl_kembali": "2023-01-10T10:00:00Z",
      "id_peminjaman": "id_peminjaman_sample",
      "id_anggota": "id_anggota_sample",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z"
    }
  }
  ```

##### 3. Membuat Denda Baru
- **Endpoint:** `POST /admin/denda/create`
- **Deskripsi:** Membuat catatan denda baru.
- **Body Permintaan (Request Body):**
  ```json
  {
    "jumlah_denda": 7500,
    "tgl_pinjam": "2023-01-01T10:00:00Z",
    "tgl_hrs_kembali": "2023-01-08T10:00:00Z",
    "tgl_kembali": "2023-01-12T10:00:00Z",
    "id_peminjaman": "id_peminjaman_sample",
    "id_anggota": "id_anggota_sample"
  }
  ```
- **Respons Sukses (201 Created):**
  ```json
  {
    "error": false,
    "msg": "Denda created successfully"
  }
  ```

##### 4. Memperbarui Denda
- **Endpoint:** `PUT /admin/denda/update`
- **Deskripsi:** Memperbarui catatan denda yang ada.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id_denda": "01H019WHW4A45KH65AH9XW89PC",
    "jumlah_denda": 10000,
    "tgl_pinjam": "2023-01-01T10:00:00Z",
    "tgl_hrs_kembali": "2023-01-08T10:00:00Z",
    "tgl_kembali": "2023-01-15T10:00:00Z",
    "id_peminjaman": "id_peminjaman_sample",
    "id_anggota": "id_anggota_sample"
  }
  ```
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Denda updated successfully"
  }
  ```

##### 5. Menghapus Denda
- **Endpoint:** `DELETE /admin/denda/delete`
- **Deskripsi:** Menghapus catatan denda berdasarkan ID.
- **Body Permintaan (Request Body):**
  ```json
  {
    "id_denda": "01H019WHW4A45KH65AH9XW89PC"
  }
  ```
- **Respons Sukses (200 OK):**
  ```json
  {
    "error": false,
    "msg": "Denda deleted successfully"
  }
  ```

1.  **Kloning repositori:**

    ```bash
    git clone https://github.com/afrizal423/Golang-Perpustakaan-Restful-API.git
    cd Golang-Perpustakaan-Restful-API
    ```

2.  **Instal Modul Go:**

    ```bash
    go mod tidy
    ```

3.  **Siapkan Variabel Lingkungan:**
    - Salin `.env.example` ke `.env`:
      ```bash
      cp .env.example .env
      ```
    - Edit `.env` dan isi kredensial database kamu serta rahasia JWT:

      ```
      # JWT settings:
      JWT_SECRET="your_jwt_secret"
      JWT_EXPIRES=15
      JWT_REFRESH_SECRET="your_jwt_refresh_secret"
      JWT_REFRESH_EXPIRES=10080

      # Server settings:
      SERVER_READ_TIMEOUT=60

      #DB Configuration
      DB_DRIVER=mysql
      DB_ADDRESS=localhost
      DB_PORT=3306
      DB_USERNAME=root
      DB_PASSWORD=
      DB_NAME=your_database_name
      ```

4.  **Siapkan Database:**
    - Pastikan server MySQL kamu berjalan.
    - Buat database yang ditentukan di `DB_NAME` dalam file `.env` kamu.

5.  **Jalankan Migrasi dan Seeder:**
    - File `main.go` saat ini memiliki baris `database.Migrate()` dan `database.Seeder(db)` yang dikomentari. Buka komentar baris-baris ini untuk sementara di `main.go` jika kamu ingin menjalankan migrasi dan mengisi data awal (pengguna admin, contoh buku).
    - Kemudian jalankan:
      ```bash
      go run main.go
      ```
    - Setelah dijalankan, komentari kembali baris tersebut untuk mencegah pengulangan saat setiap kali memulai.

6.  **Jalankan aplikasi:**

    ```bash
    go run main.go
    ```

    API akan tersedia di `http://localhost:8001`.

7.  **Jalankan Pengujian:**
    ```bash
    go test ./...
    ```
