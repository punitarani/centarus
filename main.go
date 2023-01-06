package main

import "github.com/punitarani/centarus/cmd/centarus"

func main() {
	err := centarus.Run()
	if err != nil {
		panic(err)
	}
}
