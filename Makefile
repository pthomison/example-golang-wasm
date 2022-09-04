run: wasm serve

wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm ./...
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" $(PWD)/wasm_exec.js

serve:
	python3 -m http.server 8080