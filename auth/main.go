package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == "test" && password == "123456" {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func roleMD(h httprouter.Handle, roleb string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println(time.Now(), "role")
		if roleb == "ok" {
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func jwtAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Token")
		fmt.Println(time.Now(), "jwt")
		if token == "hashtoken" {
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "index route is Not protected!\n")
}

func protected(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "protected route is jwt protected!\n")
}

func roleProtected(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "protected route is jwt protected!\n")
}

func basic(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Basic Auth Protected!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/basic", basicAuth(basic))
	//first check jwt token
	//second check role of user
	router.GET("/protecteda", jwtAuth(roleMD(protected, "ok")))
	router.GET("/protectedb", jwtAuth(roleMD(roleProtected, "ok")))

	log.Fatal(http.ListenAndServe(":8080", router))
}
