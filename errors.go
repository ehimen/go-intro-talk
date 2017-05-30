package main

import (
	"errors"
	"fmt"
)

func main() {
	// Build a map of our two functions.
	// Quick & dirty polymorphism.
	fns := map[string]func() (error, string){
		"doSomethingValid":   doSomethingValid,
		"doSomethingInvalid": doSomethingInvalid,
	}

	for fnName, fn := range fns {
		// Check to see if our function returns
		// and error. Only use str if err is nil
		if err, str := fn(); err != nil {
			fmt.Printf("%s: Something went wrong! %s\n", fnName, err)
		} else {
			fmt.Printf("%s: Success! %s\n", fnName, str)
		}
	}
}

func doSomethingValid() (error, string) {
	return nil, "Valid"
}

func doSomethingInvalid() (error, string) {
	var str string

	return errors.New("Invalid thing"), str
}
