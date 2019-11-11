package service 
 
import ( 
	// use martini framework
	"github.com/go-martini/martini" 
) 

//the server struct
type Server struct{
	handle * martini.ClassicMartini
}

//run the server
func (server * Server)Run(port string){
	// call martini.Martini.RunOnAddr()
	server.handle.RunOnAddr(port)
}
func NewServer() *Server {
	// get the ClassicMartini, a struct consisting a new Router and Martini 
	server := &Server{
		handle : martini.Classic(),
	}
	// call martini.Router.Get(),the action when server is access in the required form
	server.handle.Get("/:name", func(params martini.Params) string { 
		return "Hello " + params["name"] +"\n"
	}) 
	return server
}

