# Test Diffie-Hellman Parameters

In need of testing a Diffie-Hellman implementation? Not sure the parameters are correct?

![test diffie hellman parameters](https://www.cryptologie.net/upload/Screen_Shot_2016-02-22_at_10.28_.42_PM_.png)

## This test will check for

* decent keysize (>=2048bits)

* safe primes (modulus has to be of the form `2q + 1` with `q` prime)

## How to run it?

### Using OSX?

If you are on OSX you can use directly `test_DHparams`:

* `cat socat_dh1024_p | ./test_DHparams`

### Else?

1. get [golang](https://golang.org/)

2. get dependencies:

* `go get github.com/fatih/color`

3. examples:

* `echo "52104230423" | go run ./test_DHparams.go`

* `cat socat_dh1024_p | go run ./test_DHparams.go`
