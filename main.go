package main

import (
	"fmt"
	"log"
)

func main() {
	if err := func2(); err != nil {
		log.Printf("error: %+v", err)
	}
}

func func1() error {
	return fmt.Errorf("err1")
}

func func2() error {
	if err := func1(); err != nil {
		return fmt.Errorf("func1: %+v", err)
	}
	return nil
}
