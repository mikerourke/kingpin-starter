build-darwin:
	go build -o out/manage cmd/manage.go

build-windows:
	go build -o out/manage.exe cmd/manage.go