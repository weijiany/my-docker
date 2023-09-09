.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build main.go
.PHONY: up
up:
	vagrant up
.PHONY: reload
reload:
	vagrant reload
.PHONY: reload
halt:
	vagrant halt vm
.PHONY: generate-rootfs
generate-rootfs:
	@mkdir -p busybox
	@$(eval id=$(shell bash -c 'docker run --rm -d busybox sleep 100'))
	@docker export -o busybox.tar ${id}
	@tar -zxf busybox.tar -C busybox/
	@rm -rf busybox.tar
	@echo "copy rootfs from busybox successfully"
