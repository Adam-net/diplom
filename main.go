package main

import "diplom/pkg/server"

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
