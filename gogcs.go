/* 
  goal: read standard input and write to gcs
*/

package main

import(
  "log"
  "context"
  "os"
  "io"
  "cloud.google.com/go/storage"
)

func main() {
  var gcsc *storage.Client
  bucket := os.Args[1]
  dest := os.Args[2]
  ctx := context.Background()
  var err error
  if gcsc, err = storage.NewClient(ctx); err != nil {
    log.Fatal("GCS Client could not be created: %s", err.Error())
  }
  obj := gcsc.Bucket(bucket).Object(dest)
  w := obj.NewWriter(ctx)
  defer w.Close()
  bytes, err := io.Copy(w, os.Stdin)
  if err != nil {
    log.Fatal(err)
  } else {
    log.Printf("Wrote %d bytes to gs://%s/%s", bytes, bucket, dest)
  }
}
