package main

import "context"

func main() {
	err := run(context.TODO())
	if err != nil {
		panic(err)
	}
}
