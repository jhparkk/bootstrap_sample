package main

import (
	"github.com/urfave/negroni"
	"jhpark.sinsiway.com/bootstrap/app"
	"net/http"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)
}
