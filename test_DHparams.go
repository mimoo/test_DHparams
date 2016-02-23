package main

import(
	"fmt"
	"math/big"
	"regexp"
	"encoding/hex"
	"os"
)

func main(){
	// something to read in stdin?
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("You need to pass a DH modulus in stdin")
		return
	}

	// read input
	var input_string string
	_, err := fmt.Scan(&input_string)
	if err != nil {
		fmt.Println("You need to pass a DH modulus in hex/int format in stdin")
		return
	}

	// convert int|hex -> big
	var int_test = new(big.Int)

	var int_regex = regexp.MustCompile(`^[0-9]+$`)
	var hex_regex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	if int_regex.MatchString(input_string) {
		// int -> big
		_, err = fmt.Sscan(input_string, int_test)
	} else if hex_regex.MatchString(input_string) {
		// hex -> big
		input_bytes, err := hex.DecodeString(input_string)
		if err != nil {
			fmt.Println("Hexstring can't be parsed")
			return
		}
		int_test.SetBytes(input_bytes)
	}	else {
		// ?
		fmt.Println("input number should be either decimal or hexstring")
		return
	}

	// test for error
	if err != nil {
		fmt.Println("Couldn't understand the input number")
		return
	}

	//
	fmt.Println("Taken input:", int_test.String())

	//
	// test for safe prime
	//
	
	// q2 = p - 1
	b1 := new(big.Int)
	fmt.Sscan("1", b1)
	var q2 = new(big.Int)

	q2.Sub(int_test, b1)

	fmt.Println("p-1", q2.String())

	// q2 % b2 == 0?
	b2 := new(big.Int)
	fmt.Sscan("2", b2)
	b0 := new(big.Int)
	fmt.Sscan("0", b0)
	mod := new(big.Int)

	if b0.Cmp(mod.Mod(q2, b2)) != 0 {
		fmt.Println(mod == b0)
		fmt.Println(mod.String(), b0)
		fmt.Println("not a safe prime.")
		return
	}

	// q2 / 2 prime?
	q2.Div(q2, b2)

	fmt.Println("(p-1)/2", q2.String())

	if q2.ProbablyPrime(500) {
		fmt.Println("safe prime!")
	} else {
		fmt.Println("not a safe prime.")
	}
}
