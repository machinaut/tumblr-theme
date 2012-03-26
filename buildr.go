package main
import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func readfile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error Open("+filename+"):",err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error ReadAll("+filename+"):",err)
	}

	return string(buf)
}

type Tvmblr struct {
	JS,CSS string
}

func main() {
	tvmblrHTML := readfile("tvmblr.html")
	//js := readfile("application.js")+readfile("_functions.js")+readfile("functions.js")+readfile("posts.js")
	js := ""
	css := readfile("application.css")+readfile("normalize.css")+readfile("posts.css")

	tvmblrTemplate := template.Must(template.New("tvmblr").Parse(tvmblrHTML))
	tvmblr := Tvmblr{js,css}

	err := tvmblrTemplate.Execute(os.Stdout,tvmblr)
	if err != nil {
		log.Fatal("Error Template.Execute():",err)
	}
}
