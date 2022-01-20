
all: server client

server: $(shell find server -type f)
	cd server; go build -o ../bin/server
test:
	cd server; go test -v
client:

dev: server client
	cp configs/server_dev.yaml bin/server.yaml
	./bin/server
doc: $(shell find server/handler -type f)
	cd server; swag init

.PHONY: server
