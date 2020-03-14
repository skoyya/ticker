package server

import (
    "log"
    "net/http"
	"ticker"
	"fmt"
)

type Server struct{
	ticker *ticker.Ticker
}

func NewServer(ticker *ticker.Ticker) *Server {
	return &Server{ticker}
}

func (s *Server) RunServer() {
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(":8088", nil))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "PUT":
		tickerName, tickerNameExists := r.URL.Query()["tickerName"]
		fmt.Println(tickerName)
        if !tickerNameExists || len(tickerName[0]) < 1 {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(`{"message": "Url Param 'tickerName' is missing"}`))
            log.Println("Url Param 'tickerName' is missing")
			return
		}
		tickerValue, tickerValueExists := r.URL.Query()["tickerValue"]
		fmt.Println(tickerValue)
        if !tickerValueExists || len(tickerValue[0]) < 1 {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(`{"message": "Url Param 'tickerValue' is missing"}`))
            log.Println("Url Param 'tickerValue' is missing")
			return
		}
		if s.ticker.IsRunning() {
			s.ticker.UpdateTicker(tickerName[0], tickerValue[0])
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Update successful"}`))
		}else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "Ticker stopped already, you can not update"}`))
		}
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "Failed, Not Supported operation"}`))
    }
}
