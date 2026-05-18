# cara install library
go get ...lib url...

# cara run inisiasi
cp .env.example .env
sesuaikan nilai .env
go run cli/seed/main.go

# cara tambah config
tambahkan key dan value ke .env
tambahkan ke ./config/config.go

# cara tambah seeder
tambahkan file di ./seeders
call func di ./cli/seed/main.go

# cara run watch *.go | run dev
~/go/bin/air

# cara run server
go run main.go
