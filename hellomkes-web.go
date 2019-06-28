package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
  "time"
)

var Hostname, oserr = os.Hostname()

func HelloServer(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  dt := time.Now()
  fmt.Fprint(w,"[", dt.Format("01-02-2006 15:04:05.00"), " ", Hostname, "] Hello from hellomkes HTTPS web server.\n")
  //w.Write([]byte("Hello from hellomkes HTTPS web server.\n"))
  //fmt.Fprintf(w, "Hello from hellomkes HTTPS web server.\n")
}

func main() {
  if oserr != nil {
    fmt.Println("io.hostname error: ", oserr, " hostname value:", Hostname)
  }
  http.HandleFunc("/", HelloServer)
  err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
