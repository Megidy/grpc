package main

func main() {
	httpServer := NewHTTPServer(":1000")
	httpServer.Run()
}
