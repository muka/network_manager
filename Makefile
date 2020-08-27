
.PHONY: setup clean copy

setup:
	sudo apt install network-manager-dev --no-install-recommends
	go get -u github.com/amenzhinsky/dbus-codegen-go

clean:
	rm -rf ./interfaces

copy: clean
	mkdir -p ./interfaces
	cp /usr/share/dbus-1/interfaces/org.freedesktop.NetworkManager.* ./interfaces/
	$(eval VERSION := $(shell apt list network-manager-dev 2>/dev/null  | grep installed | awk '{print $$2}'))
	echo "package network_manager\n\nconst Version=\"${VERSION}\"" > version.go	

generate:
	$(eval LIST := $(shell ls ./interfaces/*.xml))
	dbus-codegen-go -system=true -prefix=org.freedesktop -client-only -gofmt=true -package=network_manager < $(LIST) > network_manager.go

all: setup copy generate