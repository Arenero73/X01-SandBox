package main

import "github.com/Arenero73/X01-SandBox/golang_code/hello_world"

func main() {

	x := new(hello_world.HelloWorld)

	x.SetMessage("Arenero Says Hi!")

	x.PrintMessage()

}
