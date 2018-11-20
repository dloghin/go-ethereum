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
	var bench = "Keccak256"
	if len(os.Args) > 2 {
		bench = os.Args[1]
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid size", os.Args[2])
			return
		}
		size = n
	} else {
		fmt.Println("Usage:",os.Args[0]," <Keccak256|Keccak512> <size>")
	}

	fmt.Println("Running", bench, "on input of size", size)

	start := time.Now()
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = (byte)(i % 256)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Array time", elapsed)

	if bench == "Keccak256" {
		start = time.Now()
		hash := crypto.Keccak256(data)
		t = time.Now()
		elapsed = t.Sub(start)
		fmt.Println("Keccak256 time", elapsed)
		fmt.Println(hash)
	} else if bench == "Keccak512" {
		start = time.Now()
		hash := crypto.Keccak512(data)
		t = time.Now()
		elapsed = t.Sub(start)
		fmt.Println("Keccak512 time", elapsed)
		fmt.Println(hash)
	} else {
		fmt.Println("Invalid bench", bench)
	}
}
