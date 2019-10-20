package main

import "time"

func main() {
	done := make(chan int)
	go startServer()
	time.Sleep(time.Millisecond)
	go runClient()
	<-done
}
