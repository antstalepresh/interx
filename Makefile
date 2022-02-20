install:
	./scripts/install.sh

start:
	go run main.go

publish:
	go build -o ./interxd
	nfpm pkg --packager deb --target . -f ./nfpm.yaml
