package main

import "fmt"

type errorPer struct {
	msgErr error
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("aquí está el error: %v", ep.msgErr)
}

func main() {
	ep := errorPer{
		msgErr: fmt.Errorf("Esto es un error desde main"),
	}

	foo(ep)
}

func foo(err error) {
	fmt.Println("foo corrió -", err, "\n")
}
