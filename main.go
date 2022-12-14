package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"jhpark.sinsiway.com/bootstrap/app"
)

func main() {
	m := app.MakeHandler()
	defer m.Close()

	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)
}
