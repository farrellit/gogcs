# gogcs

copy data to gcs by piping it to this program's standard input.  

## Motivation

This program achieves the following goals:

* standalone binary should run without addtional dependencies
* standard credentials resolution
* data is streamed so arbitrary amounts can be supported (real limits undefined)


## Example 

```
$ make
tmpfile=$(mktemp); \
	trap "rm -v $tmpfile"  EXIT; \
	dd bs=1m count=10 of=$tmpfile if=/dev/urandom;\
	md5sum < $tmpfile;\
	go run gogcs.go develop-203820-junk test <$tmpfile; \
	gsutil cp gs://develop-203820-junk/test - |md5sum
10+0 records in
10+0 records out
10485760 bytes transferred in 0.667804 secs (15701853 bytes/sec)
ddbc966b406f14817fc0da61ffe9513c  -
2019/02/21 00:31:23 Wrote 32768 bytes in 803.198µs
2019/02/21 00:31:35 Wrote 8421376 bytes in 12.430819679s
2019/02/21 00:31:35 Wrote 10485760 bytes to gs://develop-203820-junk/test
ddbc966b406f14817fc0da61ffe9513c  -
/var/folders/x1/bth4wnzj7fs8xts8_qkkfwyr0000gn/T/tmp.wT1t7ePH
```
