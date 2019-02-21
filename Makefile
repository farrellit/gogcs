test:
	tmpfile=$$(mktemp); \
	trap "rm -v $$tmpfile"  EXIT; \
	dd bs=1m count=10 of=$$tmpfile if=/dev/urandom;\
	md5sum < $$tmpfile;\
	go run gogcs.go develop-203820-junk test <$$tmpfile; \
	gsutil cp gs://develop-203820-junk/test - |md5sum

linux:
	GOOS=linux GOARCH=amd64 go build gogcs.go

