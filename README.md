# gogcs
copy data to gcs by piping it to this program's standard input

## Example 
```
$ echo -e "a\nb\nc\nd" | go run ./gogcs.go develop-203820-junk test
2019/02/20 18:04:12 Wrote 8 bytes to gs://develop-203820-junk/test

$ gsutil cp gs://develop-203820-junk/test -
a
b
c
d
```
