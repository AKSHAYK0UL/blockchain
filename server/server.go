package server

type Server struct {
	Port uint16 `json:"port"`
}

func NewServer(port uint16) *Server {
	return &Server{
		Port: port,
	}
}

// get port number
func (ser *Server) PortNumber() uint16 {
	return ser.Port
}

//handler
