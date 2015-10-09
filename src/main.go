package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/2", Cookie2)
	http.ListenAndServe(":8080", nil)
}
func Cookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:  "mycookie",
		Value: "helo",
		Path:  "/",
		//Domain: "localhost",
		MaxAge: 1200,
	}
	http.SetCookie(w, ck)

	ck2, err := r.Cookie("mycookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}
func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "mycookie",
		Value:  "helo",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 1200,
	}
	w.Header().Set("Set-Cookie", ck.String())
	ck2, err := r.Cookie("mycookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, ck2.Value)
}
