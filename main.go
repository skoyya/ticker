package main

import ( "server"
	     "ticker"
	 )

func main() {
    t := ticker.NewTicker()
    t.RunTicker()

    s := server.NewServer(t)
	s.RunServer()
}

