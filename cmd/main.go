package main

import "github.com/sfi2k7/picolog"

func main() {
	cl := bluelog.NewConsole("test")
	defer cl.Close()
	cl.Println("Console.lOgfo")

	lr := bluelog.New("test")
	defer lr.Close()
	lr.LogString("RRotate")
}
