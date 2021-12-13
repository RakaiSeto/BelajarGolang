package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat string
	long string
	err error
}

func (e sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", e.lat, e.long, e.err)
}

func main() {
	_, err := sqrt(-10.2)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0{
		e := errors.New("can't square numbers less than 0")
		return 0, sqrtError{"8.2312", "10.1235", e}
	}
	return 42, nil
}