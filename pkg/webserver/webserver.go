package webserver

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"

	"stefan-benten.de/goe-smart/pkg/handler"
	"stefan-benten.de/goe-smart/pkg/webserver/template"
)

type Server struct {
	Address string
	Router  *mux.Router

	Data *handler.Data
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
	_ = s.renderTemplate(w, s.Data)
}

func (s *Server) Run(data *handler.Data) {
	s.Data = data
	_ = http.ListenAndServe(s.Address, s.Router)
}
