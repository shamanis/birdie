SERVICE_NAME="birdie"

build-proto:
		scripts/build_proto.sh
bench:
		go test ./... --bench=. -benchmem
test:
		go test ./... -cover
