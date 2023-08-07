build:
	GOOS=linux GOARCH=amd64 go build main.go
up:
	vagrant up
reload:
	vagrant reload
halt:
	vagrant halt vm
generate-rootfs:
	@mkdir -p busybox
	@$(eval id=$(shell bash -c 'docker run --rm -d busybox sleep 100'))
	@docker export -o busybox.tar ${id}
	@tar -zxf busybox.tar -C busybox/
	@rm -rf busybox.tar
	@echo "copy rootfs from busybox successfully"
