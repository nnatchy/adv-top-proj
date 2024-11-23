package main

import (
	"github.com/nnatchy/adv-top-proj/internal/containers"
)

func main() {
	c := containers.NewContainer()
	if err := c.Run().Error; err != nil {
		panic(err)
	}
	// c.Run()
}
