package spass

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"spass/shamir"
)

//
// Call the shamir split/combine functions to manage a secret.
//

// read a secret from stdin; split it into parts
func Split(parts int, threshold int) {
	secretBuf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read stdin: %v\n", err)
		os.Exit(1)
	}
	byteParts, err := shamir.Split(secretBuf, parts, threshold)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to split secret: %v\n", err)
		os.Exit(1)
	}
	for _, bytePart := range byteParts {
		fmt.Printf("%x\n", bytePart)
	}
}

// read the secret parts (one per line) and combine to get secret back
func Combine() {
	var hexParts []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		hexParts = append(hexParts, scanner.Text())
	}
	var byteParts [][]byte
	for _, hexPart := range hexParts {
		b, err := hex.DecodeString(hexPart)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to decode %q: %v\n", hexPart, err)
			os.Exit(1)
		}
		byteParts = append(byteParts, b)
	}
	secretBytes, err := shamir.Combine(byteParts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to combine secret: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(secretBytes))
}
