SERVE_DIR = public
WASM_TARGET = public/main.wasm
GO_ARCH = wasm
GO_OS = js
GO_SRC = src/main.go
.PHONY: all build clean serve

all: build

build:
	GOARCH=$(GO_ARCH) GOOS=$(GO_OS) go build -o $(WASM_TARGET) $(GO_SRC)

clean:
	rm -f $(WASM_TARGET)

serve:
	cd $(SERVE_DIR) && python3 -m http.server

