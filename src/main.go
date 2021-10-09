package main


import (
	v8 "rogchap.com/v8go"
)


func main() {
	iso, _ := v8.NewIsolate() // create a new VM
	global, _ := v8.NewObjectTemplate(iso) // a template that represents a JS Object
	PrintCallback(iso, global);
	ctx, _ := v8.NewContext(iso, global) // new Context with the global Object set to our object template
	ctx.RunScript("print('Chucky')", "print.js") // will execute the Go callback with a single argunent 'foo'
}