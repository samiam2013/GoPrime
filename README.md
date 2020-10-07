# GoPrime
a highly concurrent modulus-oriented arbitrary length prime number sieving algorithm

algorithm requires 16 - 32 gb of ram

implements a single function `isBigPrime( n Int.Big ) bool { ... }` with sub-functions

and then threads it with `go` threads and `make( chan bool ) <- go channels` 

the current implementation, based on your compiler is uses an algorithm 

`lowerTimeout := int64(math.Log2(float64(timeout)))`

`lowerRange := int64(math.Log2(float64(timeout)))`

`msPerTimeout := (lowerTimeout+1)*(lowerRange+1)*lR ^ (2 ^ lowerRange)`

interestingly, if you run this code in linux the smoothed process graph looks a lot like the Reimann-Zeta function.


