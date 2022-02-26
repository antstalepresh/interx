install:
	./scripts/install.sh

start:
	go run main.go

publish:
	./scripts/install.sh
	nfpm pkg --packager deb --target . -f ./nfpm.yaml