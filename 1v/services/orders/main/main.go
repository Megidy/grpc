package main

func main() {
	httpServer := NewHTTPServer(":8080")
	go httpServer.Run()
	grpcServer := NewGRPCServer(":9000")
	grpcServer.Run()

}