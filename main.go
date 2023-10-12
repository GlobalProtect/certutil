package main

import (
	"./internal/certutil"
	"fmt"
)

func main() {
	if err := certutil.Run(); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
