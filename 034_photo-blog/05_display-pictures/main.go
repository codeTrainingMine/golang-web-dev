package main

import (
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
	"strings"
	"fmt"
	"os"
	"path"
	"log"
	"io"
	"crypto/sha1"
	"io/ioutil"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/pictures/", http.StripPrefix("/pictures", http.FileServer(http.Dir("./pictures"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//fmt.Print("in index")
	c := getCookie(w, req)
	if req.Method == http.MethodPost {
		fmt.Println("post")
		ff, fh, err := req.FormFile("myfileupload")
		if err != nil {
			log.Fatalln(err)
		}
		defer ff.Close()

		//h := sha1.New()
		b, _ := ioutil.ReadAll(ff)
		//fmt.Printf("%x%s\n", sha1.Sum(b), filepath.Ext(fh.Filename))
		filename := fmt.Sprintf("%x%s", sha1.Sum(b), filepath.Ext(fh.Filename))

		wd, err := os.Getwd()
		fullFilename := path.Join(wd, "pictures", filename)
		nf, err := os.Create(fullFilename)
		if err != nil {
			log.Fatalln(err)
		}
		defer nf.Close()

		ff.Seek(io.SeekStart, io.SeekStart)
		io.Copy(nf, ff)

		c = appendValue(w, c, filename)
	}


	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:])
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, filename string) *http.Cookie {
	// append
	s := c.Value
	if !strings.Contains(s, filename) {
		s += "|" + filename
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
