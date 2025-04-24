package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Comando %v non esistente", os.Args[1])
	}
}

func compare(password, hash string) {
	fmt.Printf("Compare %q to %q", password, hash)
}

func hash(password string) {
	fmt.Printf("Hash this password:%q", password)
}
