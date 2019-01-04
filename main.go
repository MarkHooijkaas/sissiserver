package main
import (
  "net/http"
  "strings"
	"bytes"
	"io/ioutil"
	"log"
	"os"
)


func check(err error) {
	if err != nil {
			panic(err)
	}
}

func handle(out http.ResponseWriter, req *http.Request) {
  path := req.URL.Path
  path = strings.TrimPrefix(path, "/")
	if path!="favicon.ico" {
  	handleFile(path, out)
	}
}

func handleFile(path string, out http.ResponseWriter) {
	if isDir(path) || path=="" {
		path=path+"index.html"
	}
	log.Println("Loading: "+path)
	txt, err := ioutil.ReadFile(path)
	if err!=nil {
		log.Println("ERROR: Could not load "+path)
	}
	parts:=bytes.Split(txt,[]byte("<!--#include file=\""))
	out.Write([]byte(parts[0]))
	for i:=1; i<len(parts); i++ {
		pos:=bytes.IndexAny(parts[i], "-->")
		includeFile:=parts[i][0:pos]
		pos2:=bytes.IndexAny(includeFile,"\"")
		//out.Write([]byte("***"))
		//out.Write(includeFile[:pos2])
		//out.Write([]byte("***"))
	  handleFile(string(includeFile[:pos2]), out)
		out.Write(parts[i][pos+3:])
	}
}

func main() {
  http.HandleFunc("/", handle)
  err := http.ListenAndServe(":8090", nil)
  check(err)
}

func isDir(path string) bool {
    fi, err := os.Stat(path)
    return err == nil && fi.Mode().IsDir()
}
