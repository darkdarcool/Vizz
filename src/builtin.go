package main

import "fmt"
import v8 "rogchap.com/v8go"

func PrintCallback(iso *v8.Isolate, global *v8.ObjectTemplate) {
	printfn, _ := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
		fmt.Printf("Hi %v!\n", info.Args()[0]) // when the JS function is called this Go callback will execute
		return nil // you can return a value back to the JS caller if required
	})
	global.Set("print", printfn) // sets the "print" property of the Object to our function
}

