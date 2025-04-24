package main

import (
	"fmt"
	"os"
	"golang.org/x/crypto/bcrypt"
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
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	if err!=nil{
		fmt.Printf("Password is invalid %v\n",password)
		return
	}
	fmt.Println("Password is correct")
}

func hash(password string) {
	hashedBytes,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		fmt.Printf("error hashing: %v\n",password)
		return
	}
	fmt.Println(string(hashedBytes))
}
