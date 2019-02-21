package main

import(
  "log"
  "context"
  "os"
  "io"
  "cloud.google.com/go/storage"
  "time"
  //"flag"
)

type ProgressWriter struct {
    io.Writer
    total uint64
    increment time.Duration
    start time.Time
    last time.Time
}

func (pt *ProgressWriter) Write(p []byte) (int, error) {
  if pt.total == 0 {
    pt.start = time.Now()
  }
  n, err := pt.Writer.Write(p)
  pt.total += uint64(n)
  if err == nil {
    now := time.Now()
    if now.Sub(pt.last) > ( 10 * time.Second ) {
      log.Println("Wrote", pt.total, "bytes in", now.Sub(pt.start) )
      pt.last = now
    }
  }
  return n, err
}


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
  bytes, err := io.Copy(&ProgressWriter{Writer:w, increment: 1 * time.Second }, os.Stdin)
  if err != nil {
    log.Fatal(err)
  } else {
    log.Printf("Wrote %d bytes to gs://%s/%s", bytes, bucket, dest)
  }
}
