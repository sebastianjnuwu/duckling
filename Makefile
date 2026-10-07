# Makefile — Duckling (Go + Ebiten, só Windows + web)
# Uso: mingw32-make <alvo> (ou make no MSYS/Linux).
# Builds magros: -trimpath corta paths locais, -s -w tira
# símbolos/debug. UPX depois (~-45%).

APP := duckling
TMP := tmp
FLAGS := -trimpath -ldflags="-s -w"

.PHONY: run vet build-win64 build-win32 web pack size clean

run:
	go run .

vet:
	go vet ./...

build-win64: vet
	GOARCH=amd64 go build $(FLAGS) -o $(TMP)/$(APP)-win64.exe .

build-win32: vet
	go build $(FLAGS) -o $(TMP)/$(APP)-win32.exe .

web: vet
	GOOS=js GOARCH=wasm go build $(FLAGS) -o $(TMP)/web/$(APP).wasm .

pack:
	$(TMP)/upxbin/upx.exe --best $(TMP)/$(APP)-win*.exe

size:
	ls -lh $(TMP)/$(APP)-win*.exe

clean:
	rm -f $(TMP)/$(APP)-win*.exe $(TMP)/web/$(APP).wasm
