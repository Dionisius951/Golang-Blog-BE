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
| POST   | `/auth/register`       | Register pengguna baru       | No   |
| POST   | `/auth/login`          | Login pengguna               | No   |
| GET    | `/posts`               | Ambil semua posting          | No   |
| GET    | `/posts/{id}`          | Ambil posting by ID          | No   |
| POST   | `/posts`               | Tambah posting baru          | Yes  |
| PUT    | `/posts/{id}`          | Update posting               | Yes  |
| DELETE | `/posts/{id}`          | Hapus posting                | Yes  |
| GET    | `/categories`          | Ambil semua kategori         | No   |
| POST   | `/categories`          | Tambah kategori              | Yes  |
| GET    | `/posts/{id}/comments` | Ambil komentar pada posting  | No   |
| POST   | `/posts/{id}/comments` | Tambah komentar pada posting | Yes  |

