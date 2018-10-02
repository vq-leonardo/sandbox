package api

import (
	"log"

	context "golang.org/x/net/context"
)

// Server struct
type Server struct {
}

// SayHello func
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive mesage %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
