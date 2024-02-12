package lib

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type HttpHandler struct {
	Engine *httprouter.Router
}

func NewHttpHandler() HttpHandler {
	return HttpHandler{
		Engine: httprouter.New(),
	}
}

func (h HttpHandler) Run() error {
	srv := http.Server{
		Addr:         ":8000",
		Handler:      h.Engine,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return srv.ListenAndServe()
}
