GO=tinygo

.PHONY: json draganddrop fetch websocket node
all: json draganddrop fetch websocket node

node:
	$(GO) build  -o ./example/static/node.wasm --no-debug -target wasm example/node/main.go

json:
	$(GO) build  -o ./example/static/json.wasm --no-debug -target wasm example/json/main.go

fetch:
	$(GO) build  -o ./example/static/fetch.wasm --no-debug -target wasm example/fetch/main.go

draganddrop:
	$(GO) build  -o ./example/static/draganddrop.wasm --no-debug -target wasm example/draganddrop/main.go

websocket:
	$(GO) build  -o ./example/static/websocket.wasm --no-debug -target wasm example/websocket/main.go