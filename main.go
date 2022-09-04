package main

import (
    "fmt"
    "syscall/js"
)

func main() {
    fmt.Println("Go Web Assembly")
    js.Global().Set("invokeWASM", jsInvokable())

    <-make(chan bool)
}

func jsInvokable() js.Func {

    function := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

        if len(args) != 1 {
            return "Invalid no of arguments passed"
        }

        inputString := args[0].String()
        fmt.Printf("input %s\n", inputString)
        return inputString
    })

    return function
}
