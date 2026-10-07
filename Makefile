APP := duckling
TMP := tmp
FLAGS := -trimpath -ldflags="-s -w"

.PHONY: run vet build-win64 build-win32 web android android-install pack size clean

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

android: vet
	~/go/bin/windows_amd64/gomobile-duckling build -androidapi 24 -o $(TMP)/$(APP).apk .

android-install:
	adb install -r $(TMP)/$(APP).apk

pack:
	$(TMP)/upxbin/upx.exe --best $(TMP)/$(APP)-win*.exe

size:
	ls -lh $(TMP)/$(APP)-win*.exe

clean:
	rm -f $(TMP)/$(APP)-win*.exe $(TMP)/web/$(APP).wasm
