package ui

import "net/http"

type HTTPServer struct {

}

func NewHTTP() *HTTPServer {
	return &HTTPServer{}
}

func (HTTPServer) UseService(s Service){
}

func (HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/todos/" && r.URL.Path != "/todos"{
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}