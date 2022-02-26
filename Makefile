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
	./scripts/build.sh
	nfpm pkg --packager deb --target . -f ./nfpm.yaml