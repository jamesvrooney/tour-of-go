package main

import "fmt"

type Vertex1 struct {
	X, Y int
}

func main() {
	v := Vertex1{3, 5}
	v.X = 45

	fmt.Println("x =", v.X, "y =", v.Y)
}
