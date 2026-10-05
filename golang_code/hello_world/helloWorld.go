package hello_world

import "fmt"

type HelloWorld struct {
	Message string
}

func (hw *HelloWorld) GetMessage() string {
	return hw.Message
}

func (hw *HelloWorld) SetMessage(message string) {
	hw.Message = message
}

func (hw *HelloWorld) PrintMessage() {
	fmt.Printf("%s\n", hw.Message)
}
