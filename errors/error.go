package main

import (
	"errors"
	"fmt"
)

func f1(num int) (int, error) {
	if num == 2 {
		return -1, errors.New("error for 2")
	}
	return num, nil
}

func main() {
	for _, i := range []int{1, 2, 3} {
		r, e := f1(i)
		if e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 success:", r)
		}
	}

}
