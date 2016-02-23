package main

import(
	"fmt"
	"math/big"
	"regexp"
	"encoding/hex"
	"os"
	"github.com/fatih/color"
)

func test_bitLength(modulus_int *big.Int) (int) {
	fmt.Println(modulus_int.BitLen())
	return modulus_int.BitLen()
}

func test_safePrime(modulus_int *big.Int) (bool) {
	// q2 = p - 1
	b1 := new(big.Int)
	fmt.Sscan("1", b1)

	var q2 = new(big.Int)

	q2.Sub(modulus_int, b1)

	//fmt.Println("p-1", q2.String())

	// q2 % b2 == 0?
	b2 := new(big.Int)
	fmt.Sscan("2", b2)
	b0 := new(big.Int)
	fmt.Sscan("0", b0)
	mod := new(big.Int)

	if b0.Cmp(mod.Mod(q2, b2)) != 0 {
		return false
	}

	// q2 / 2 prime?
	q2.Div(q2, b2)

	if q2.ProbablyPrime(500) {
		return true
	} else {
		return false
	}

}

func main(){
	// something to read in stdin?
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("You need to pass a DH modulus in stdin")
		return
	}

	// read input
	var modulus_str string
	_, err := fmt.Scan(&modulus_str)
	if err != nil {
		fmt.Println("You need to pass a DH modulus in hex/int format in stdin")
		return
	}

	// convert int|hex -> big
	var modulus_int = new(big.Int)

	var int_regex = regexp.MustCompile(`^[0-9]+$`)
	var hex_regex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

	if int_regex.MatchString(modulus_str) {
		// int -> big
		_, err = fmt.Sscan(modulus_str, modulus_int)
	} else if hex_regex.MatchString(modulus_str) {
		// hex -> big
		input_bytes, err := hex.DecodeString(modulus_str)
		if err != nil {
			fmt.Println("Hexstring can't be parsed")
			return
		}
		modulus_int.SetBytes(input_bytes)
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
	fmt.Println("Taken input:", modulus_int.String())

	// test for Bitlength
	if bitlen := test_bitLength(modulus_int); bitlen >= 2048 {
		color.Green("Good modulus bitlength!")
		fmt.Println("(Mobulus is", bitlen, "bits)")
	} else {
		color.Red("Bad modulus bitlength! Should be at least 2048 bits")
		fmt.Println("(Mobulus is", bitlen, "bits)")
	}

	// test for safe prime
	if test_safePrime(modulus_int) {
		color.Green("The modulus is a safe prime!")
	} else {
		color.Red("The modulus is NOT a safe prime!")
	}
	
}
