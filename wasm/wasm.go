// +build js,wasm
package main

import(
	"syscall/js"
)

func main() {
	listen := make(chan bool)

	boton := js.Global().Get("document").Call("getElementById", "boton")

	boton.Call("addEventListener", "click", js.FuncOf(Change))
	<-listen
}

func Change(this js.Value, args []js.Value) interface{} {
	js.Global().Get("document").Call("getElementById", "texto").Set("innerHTML", "Hello, WebAssembly!")
	return nil
}