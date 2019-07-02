package main

import (
	"net/http"
	"strings"
)

func main() {
	//very bad funcs
	//exec.Command
	//By default, the http.Dir filesystem has directory indexing enabled. For example,
	//let's say you have a .git/ folder at the root of the folder you're serving.
	//If someone were to request your_url/.git/, the contents of the folder would be listed.
	// use https://github.com/jordan-wright/unindexed

	//http.FileServer()
	//http.ServeFile()
}

//StaticFile get static files from dist folder
func StaticFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p := r.URL.Path[5:]
		// very bad, not setted folder path
		/*
			hacking sample
				/etc/passwd
				..%2F..%2F..%2Fetc%2F
		*/
		http.ServeFile(w, r, "./myfiles/"+p)
		return
	}
}

//RmFilePath remove bad chars
func RmFilePath(p string) string {
	p = strings.Replace(p, `@`, "", -1)
	p = strings.Replace(p, `$`, "", -1)
	p = strings.Replace(p, `:`, "", -1)
	p = strings.Replace(p, `\`, "", -1)
	p = strings.Replace(p, "../", "", -1)
	p = strings.Replace(p, "~/", "", -1)
	p = strings.Replace(p, "%", "", -1)
	p = strings.Replace(p, " ", "", -1)
	p = strings.Replace(p, "~", "", -1)
	return p
}
