package model

type Service struct {
	Model
	Name        string
	Description string
	Port        int
	Server      *Server
	ServerID    uint64
	Config      map[string]interface{}
}
