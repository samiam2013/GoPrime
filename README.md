# GoPrime
a highly concurrent modulus-oriented arbitrary length prime number sieving algorithm

algorithm requires 16 - 32 gb of ram

implements a single function [`isBigPrime( n Int.Big ) bool { ... }`](

and then threads it with `go` threads and `make( chan bool ) <- go channels` 

the current implementation, based on your compiler is uses an algorithm 

`lowerTimeout := int64(math.Log2(float64(timeout)))`

`lowerRange := int64(math.Log2(float64(timeout)))`

`msPerTimeout := (lowerTimeout+1)*(lowerRange+1)*lR ^ (2 ^ lowerRange)`

`msPerTimeout := msPerTimout ^ timeout + 1`

I think this is like forward linear algebra, or "lambdas" or "a function" or "method" but I never took Differential Equations in College so I can't tell you.

it uses euler's modulus that I believe runs in O(n) time to elimiate all the prime factors from 3 to 11

[`func isModBig(n *big.Int) bool { ... }`](https://github.com/samiam2013/GoPrime/blob/801109614645e52d0245abaf189922833902306f/primeCheckerParallel.go#L44)

it skips checking anything that isn't a prime number by having a constant list of prime numbers to search and find the next

to estimate the amount of time it will take to finish a thread based on the prime numbers' length in binary numberspace 

I don't know why this is, don't ask me, 

and this floors it (hahahahhahahaha it's fast) on the fan function timeouts in milliseconds

and it skips all the hub-ub with checking it's mod 0 with anything by running a faster deterministic run of function first concurrently with the main algorithm call isBigPrime() in a Go `select`. 

I did notice that the graph of memory use in windows is very similar to the solution graph of the [Reimann-Zeta Hypothesis](https://en.wikipedia.org/wiki/Riemann_hypothesis) (which is an "imaginary" mathematical function) integrated into total work needed for positive integers over time, because the work gets easier as the prime numbers skew apart at higher numbers

I've thought a lot about prime numbers

this algorithm is not turing-complete because of the enourmous amount of ram you need to run it.

And I think Alan Turing did too. He tried to make a gear system that would solve the Reimann-Zeta function. Now why would you go about doing that if you didn't just really like prime numbers and solving cryptopgraphy?


