package main

import (
	sgrpc "github.com/dielit66/microframework/pkg/grpc"
	"github.com/dielit66/microframework/pkg/http"
	ws "github.com/dielit66/microframework/pkg/websocket"
)

func main() {
	h := http.NewHTTPServer(":8080")
	g := sgrpc.NewGRPCServer()
	w := ws.NewWSServer(":8052")

	go h.Start()
	go g.Start(":8022")
	go w.Start()
	select {}
}
