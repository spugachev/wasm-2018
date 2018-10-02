package main

import (
	"fmt"
	"syscall/js"
)

var document js.Value

func init() {
	document = js.Global().Get("document")
}

func main() {
	fmt.Println("Hello, from Go!")

	div := document.Call("getElementById", "target")
	node := document.Call("createElement", "div")
	node.Set("innerHTML", "<h1>Hello from GO!</h1>")
	div.Call("appendChild", node)
}
