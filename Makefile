# Makefile — Duckling (Go + Ebiten: Windows, Linux + web)
# Uso: mingw32-make <alvo> (ou make no MSYS/Linux).
# Builds magros: -trimpath corta paths locais, -s -w tira
# símbolos/debug. Um arquivo por plataforma. UPX depois (~-45%).
# Linux cruza daqui mas pede toolchain C (zig cc); no Linux
# nativo, os mesmos alvos funcionam direto.

APP := duckling
TMP := tmp
FLAGS := -trimpath -ldflags="-s -w"

.PHONY: run vet syso build-win64 build-win32 build-linux64 build-linux32 web pack size clean

run:
	go run .

vet:
	go vet ./...

syso: rsrc_windows_386.syso rsrc_windows_amd64.syso

rsrc_windows_%.syso: content/images/logo.png
	go run github.com/tc-hib/go-winres@latest simply --icon $< --manifest none --arch $* --out rsrc

build-win64: vet rsrc_windows_amd64.syso
	GOARCH=amd64 go build $(FLAGS) -o $(TMP)/$(APP)-win64.exe .

build-win32: vet rsrc_windows_386.syso
	go build $(FLAGS) -o $(TMP)/$(APP)-win32.exe .

build-linux64: vet
	GOOS=linux GOARCH=amd64 go build $(FLAGS) -o $(TMP)/$(APP)-linux64 .

build-linux32: vet
	GOOS=linux GOARCH=386 go build $(FLAGS) -o $(TMP)/$(APP)-linux32 .

web: vet
	GOOS=js GOARCH=wasm go build $(FLAGS) -o $(TMP)/web/$(APP).wasm .

pack:
	$(TMP)/upxbin/upx.exe --best $(TMP)/$(APP)-win*.exe $(TMP)/$(APP)-linux*

size:
	ls -lh $(TMP)/$(APP)-win*.exe $(TMP)/$(APP)-linux*

clean:
	rm -f $(TMP)/$(APP)-win*.exe $(TMP)/$(APP)-linux* $(TMP)/web/$(APP).wasm
