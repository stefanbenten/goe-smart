package webserver

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"

	"stefan-benten.de/goe-smart/pkg/webserver/template"
)

type Server struct {
	Address string
	Router  *mux.Router
}

func NewServer(addr string) (server Server, err error) {

	srv := Server{}

	router := mux.NewRouter()
	router.HandleFunc("/", srv.handleRoot)

	return Server{Router: router, Address: addr}, nil
}

func (s *Server) renderTemplate(w http.ResponseWriter, data interface{}) (err error) {

	buf := bytes.Buffer{}

	if err = template.Index.Execute(&buf, data); err != nil {
		return err
	}

	_, err = w.Write(buf.Bytes())
	return
}

func (s *Server) handleRoot(w http.ResponseWriter, r *http.Request) {
	_ = s.renderTemplate(w, struct {
		Text string
	}{"Dies ist ein Text"})
}

func (s *Server) Run() {
	_ = http.ListenAndServe(s.Address, s.Router)
}
