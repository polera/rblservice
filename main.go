package main

import (
	"github.com/polera/rblservice/search"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	//	"io"
	"net/http"
)

func lookup(c web.C, w http.ResponseWriter, r *http.Request) {
	res := search.Run(c.URLParams["host"])
	w.Write(res)

}

func main() {
	goji.Use(JSONResponse)
	goji.Get("/lookup/:host", lookup)
	goji.Serve()
}
