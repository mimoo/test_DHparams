# Test Diffie-Hellman Parameters

In need of testing a Diffie-Hellman implementation? Not sure the parameters are correct?

## This test will check for

* decent keysize (>=2048bits)

* safe primes (modulus has to be of the form `2q + 1` with `q` prime)

## How to run it?

examples:

* `echo "52104230423" | go run ./test_DHparams.go`

* `cat socat_dh1024_p | go run ./test_DHparams.go`

If you are on OSX you can use directly `test_DHparams`:

* `cat socat_dh1024_p | ./test_DHparams`