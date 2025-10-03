# Golang-Blog-BE

Repo ini adalah project pembelajaran dari Backend API untuk aplikasi blog menggunakan Go.  
Mendukung autentikasi JWT, manajemen posting, kategori, dan komentar.

## Fitur

- Register dan login pengguna
- Proteksi endpoint dengan JWT
- CRUD posting blog
- CRUD kategori
- Komentar pada posting
- Middleware autentikasi dan validasi

## Instalasi

1. Clone repository
   ```bash
   git clone https://github.com/Dionisius951/Golang-Blog-BE.git
   cd Golang-Blog-BE
2. Install Dependency
   ```bash
   go mod tidy
3. Jalankan App
   ```bash
   go run cmd/main.go

## Route

| Method | Path                   | Deskripsi                    | Auth |
| ------ | ---------------------- | ---------------------------- | ---- |
| POST   | `/register`            | Register pengguna baru       | No   |
| POST   | `/login`               | Login pengguna               | No   |
| GET    | `/post`                | Ambil semua posting          | No   |
| GET    | `/post/{id}`           | Ambil posting by ID          | No   |
| POST   | `/post`                | Tambah posting baru          | Yes  |
| PUT    | `/post/{id}`           | Update posting               | Yes  |
| DELETE | `/post/{id}`           | Hapus posting                | Yes  |
| GET    | `/comment/{id}`        | Ambil komentar pada posting  | No   |
| POST   | `/comment/{id}`        | Tambah komentar pada posting | Yes  |
| PUT    | `/comment/{id}`        | Update komentar pada posting | No   |
| DELETE | `/comment/{id}`        | Delete komentar pada posting | Yes  |

