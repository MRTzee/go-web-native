# GO Web Native 🌐

Aplikasi web sederhana yang dibangun dengan Go menggunakan pendekatan native.

## 🚀 Setup dan Instalasi

1. **Persiapan Docker**
   ```bash
   docker-compose up --build
   ```

2. **Jalankan Migrasi Database**
   ```bash
   migrate -database "mysql://root:password@tcp(localhost:3306)/mydb" -path db/migrations up
   ```

3. **Aplikasi Siap Digunakan! 🎉**

## 📋 Prasyarat

- Docker dan docker-compose terinstall
- Go-migrate CLI terinstall
- MySQL

## 🗄️ Struktur Database

Database menggunakan MySQL dengan konfigurasi berikut:
- Host: localhost
- Port: 3306
- Username: root
- Password: password
- Database: mydb

## 🛠️ Tools yang Digunakan

- Docker untuk containerization
- Go-migrate untuk database migration
- MySQL sebagai database

## 👨‍💻 Developed By
[MRTzee](https://github.com/mrtzee)

## 📄 License
MIT License

---
⭐ Jangan lupa beri bintang jika project ini membantu Anda!
