// // package main tells go that this is an application and not a library hence look for the main function to start
// package main

// // fmt is Go's standard formatting package, that help you print to the console
// // net/http is the http server and client library, defines request handlers and start an http server
// import (
// 	"fmt"
// 	"net/http"
// 	"log"
// )

// // the line below declares a new type called server, and in it addr is a string
// // by storing the address in a struct, it not only stores the address but also anything else the handlers need like DB connection, logger instance, etc.
// type server struct {
// 	addr string
// }

// // in the function below the server type is attached to the ServeHTTP function, hence *server is a valid HTTP handler
// // the function below is the main entry point for every incoming request, whenever someone does sth like GET http://localhost:8081/foo, GO's HTTP server will call the serverHttp method, and pass in the response writer w to write back headers and body, also the r Request object describing the incoming request
// // the method below is similar to app.use((req, res))
// // any function that looks like the two below, those are handlers, and they are callbacks that HTTP server invokes for every incoming request
// // func(w http.ResponseWriter, r *http.Request) OR
// // ServeHTTP(w http.ResponseWriter, r *http.Request)

// func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// HTTP sends raw bytes over the wire hence convert GO string to []byte a format that is required
// 	w.Write([]byte("Hello Server running on port " + s.addr))
// 	fmt.Println("Hello World")
// }

// func main() {

// 	// instance of the server, with value ":8081" from addr variable
// 	s := &server{addr: ":8081"}
// 	// pass the s in place of the nil handler. Go sees the s has a serveHttp method, so it'll call the s.ServeHTTP for every request
// 	err :=http.ListenAndServe(s.addr, s)
// 	if err !=nil{
// 		log.Fatal(err)
// 	}
// }

//  ********* REFACTOR 2 *********

// package main

// import "net/http"

// type api struct{
// 	addr string
// }

// func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	switch r.Method{
// 	case http.MethodGet:
// 		switch r.URL.Path{
// 		case "/":
// 			w.Write([]byte("Get Method"))
// 		case "/index":
// 			w.Write([]byte("Index"))
// 		}
// 	case http.MethodPost:
// 		w.Write([]byte("Post Method"))
// 	}
// }

// // additional handler methods
// func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Create User"))
// }

// func (a *api) getUsershandler(w http.ResponseWriter, r *http.Request){
// 	w.Write([]byte("Get User..."))
// }

// func main(){
// 	api := &api{addr:":8081"} // creating instance of api
// 	// a mux is a router (request mutiplexer)
// 	mux :=http.NewServeMux()
// 	// creating a server struct with configs
// 	srv := &http.Server{
// 		Addr: api.addr,
// 		Handler: mux,
// 	}

// 	mux.HandleFunc("GET /users", api.getUsershandler)
// 	//mux needs a string and a handlerFunc as its arguments
// 	mux.HandleFunc("POST /users", api.createUsersHandler)

// 	srv.ListenAndServe()
// }

// ******** MORE FINE REFACTOR ************
package main

import ("net/http"
		"log"

)

// type api struct{
// 	addr string
// }

func main(){

	api := &api{addr: ":8081"}
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)

	err :=srv.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}


}