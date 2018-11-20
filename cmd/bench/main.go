package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var size = 1000000	
	if len(os.Args) > 1 {		
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid size", os.Args[1])
			return
		}
		size = n
	}
	
	fmt.Println("Running Keccak256 and Keccak512 on input of size", size)
	
	start := time.Now()
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = (byte)(i % 256)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Array time", elapsed)
	
	start = time.Now()
	hash := crypto.Keccak256(data)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Keccak256 time", elapsed)
	fmt.Println(hash)
	
	start = time.Now()
	hash = crypto.Keccak512(data)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Keccak512 time", elapsed)
	fmt.Println(hash)
}