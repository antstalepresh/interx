proto-gen:
	./scripts/proto-gen.sh

install:
	./scripts/proto-gen.sh
	./scripts/install.sh

start:
	go run main.go

publish:
	./scripts/proto-gen.sh
	./scripts/install.sh
	nfpm pkg --packager deb --target . -f ./nfpm.yaml