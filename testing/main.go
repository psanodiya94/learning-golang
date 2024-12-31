package main

import (
	"errors"
	"log"
)

func devide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return 0, errors.New("Cannot devide by zero")
	}
	result = x / y
	return result, nil
}

func main() {

	result, err := devide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Result: ", result)
}
