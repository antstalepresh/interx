.PHONY: proto-gen build publish install start test test-local network-start network-stop

proto-gen:
	./scripts/proto-gen.sh

install:
	./scripts/proto-gen.sh
	./scripts/build.sh

build:
	./scripts/build.sh

start:
	go run main.go

publish:
	./scripts/proto-gen.sh
	./scripts/publish.sh

test:
	./scripts/test.sh

test-local:
	./scripts/test-local.sh

network-start:
	./scripts/test-local/network-stop.sh
	./scripts/test-local/network-start.sh

network-stop:
	./scripts/test-local/network-stop.sh