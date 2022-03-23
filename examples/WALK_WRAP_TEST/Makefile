32:
	GOARCH=386 go build -ldflags '-s -w -H windowsgui' -o WALK_WRAP32.exe
64:
	GOARCH=amd64 go build -ldflags '-s -w -H windowsgui' -o WALK_WRAP64.exe
rsrc:
	rsrc -manifest WALK_WRAP.manifest -ico test.ico -o WALK_WRAP.syso
all:
	make 32 64