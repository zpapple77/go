package main

import "fmt"

func funcRecover() error {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("panic recover")
		}
	}
}
