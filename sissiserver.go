package main
import (
  "net/http"
  "strings"
	"bytes"
	"io/ioutil"
	"log"
	"os"
  "flag"
)


func check(err error) {
	if err != nil {
			panic(err)
	}
}

func handle(out http.ResponseWriter, req *http.Request) {
  path := req.URL.Path
	log.Println("request for "+path)
	handleFile(path, out, "\t")
}

func handleFile(path string, out http.ResponseWriter, indent string) {
  path=dir+strings.TrimPrefix(path, "/")
	if isDir(path) || path=="" {
		path=path+"index.html"
	}
	log.Println(indent+"- "+path)
	txt, err := ioutil.ReadFile(path)
	if err!=nil {
		log.Println(indent+"ERROR: Could not read "+path)
	}
	parts:=bytes.Split(txt,[]byte("<!--#include virtual=\""))
	out.Write([]byte(parts[0]))
	for i:=1; i<len(parts); i++ {
		pos:=bytes.Index(parts[i], []byte("-->"))
		includeFile:=parts[i][0:pos]
		pos2:=bytes.Index(includeFile,[]byte("\""))
	  handleFile(string(includeFile[:pos2]), out, indent+"\t")
		out.Write(parts[i][pos+3:])
	}
}

func getConfig(name string, defaultValue string) string {
	value:=os.Getenv(name)
	if len(value)>0 { return value}
	return defaultValue
}

var dir string
var port string
func main() {
  flag.StringVar(&dir, "dir", "html", "the directory with html files")
  flag.StringVar(&port, "port", "8090", "the port to listen to")
  flag.Parse()
  dir=dir+"/"
	log.Println("Listening on http://localhost:"+port+" to directory "+dir)
	http.HandleFunc("/", handle)
  err := http.ListenAndServe(":"+port, nil)
  check(err)
}

func isDir(path string) bool {
    fi, err := os.Stat(path)
    return err == nil && fi.Mode().IsDir()
}
