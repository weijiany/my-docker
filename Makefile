build:
	GOOS=linux GOARCH=amd64 go build main.go
up:
	vagrant up
reload:
	vagrant reload
halt:
	vagrant halt vm
