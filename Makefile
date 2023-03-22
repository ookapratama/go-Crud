run:
	go run utama.go

tes:
	go mod tidy
	nodemon --exec go run utama.go --signal SIGTERM