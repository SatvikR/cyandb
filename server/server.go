package server

import "os"

const (
	// Use this for default db functionality
	DefaultDBPath = "/data/cyan/"
	dataFileName  = "db.dat"
)

// Server is the struct definition for the server
// I will add more to this once websockets are introduced
type Server struct {
	Location string
}

// CreateServer creates a server struct
func CreateServer(path string) *Server {
	// Make sure path ends in '/'
	if path[len(path)-1:] != "/" {
		path = path + "/"
	}

	// Create db directory if doesn't exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, 0777)
	}

	server := &Server{Location: path + dataFileName}

	return server
}
