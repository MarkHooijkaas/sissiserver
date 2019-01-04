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
	log.Println("request for "+path)
	handleFile(path, out)
}

func handleFile(path string, out http.ResponseWriter) {
  path=dir+strings.TrimPrefix(path, "/")
	if isDir(path) || path=="" {
		path=path+"index.html"
	}
	log.Println("Reading: "+path)
	txt, err := ioutil.ReadFile(path)
	if err!=nil {
		log.Println("ERROR: Could not read "+path)
	}
	parts:=bytes.Split(txt,[]byte("<!--#include file=\""))
	out.Write([]byte(parts[0]))
	for i:=1; i<len(parts); i++ {
		pos:=bytes.IndexAny(parts[i], "-->")
		includeFile:=parts[i][0:pos]
		pos2:=bytes.IndexAny(includeFile,"\"")
	  handleFile(string(includeFile[:pos2]), out)
		out.Write(parts[i][pos+3:])
	}
}

func getConfig(name string, defaultValue string) string {
	value:=os.Getenv(name)
	if len(value)>0 { return value}
	return defaultValue
}

var dir string
func main() {
	port:=getConfig("port","8090")
	dir=getConfig("dir","site")+"/"
	log.Println("Listening on port "+port+" to directory "+dir)
	http.HandleFunc("/", handle)
  err := http.ListenAndServe(":"+port, nil)
  check(err)
}

func isDir(path string) bool {
    fi, err := os.Stat(path)
    return err == nil && fi.Mode().IsDir()
}
