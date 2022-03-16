package main

import (
	"fmt"
	"log"
	"net/http"
	"archive/zip"
	"io"
	"os"
	"time"
	"runtime"
)

func getArchiveFile(w http.ResponseWriter, r *http.Request){
	t := time.Now()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")

	zw := zip.NewWriter(w)
	zf, err := zw.Create("archive.zip");
	if err != nil {
    	fmt.Print(err)
    }
	f1, _ := os.Open("2GB.bin")
	defer f1.Close()

	_, err = io.Copy(zf, f1)
	err = zw.Close()
	if err != nil {
    	fmt.Print(err)
    }
	// w.Write(<-archiveFile("archive"))
	fmt.Println(time.Since(t).Seconds());
	
}

func handleRequest() {
	http.HandleFunc("/zip", getArchiveFile);
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	runtime.GOMAXPROCS(1)
	handleRequest()
}