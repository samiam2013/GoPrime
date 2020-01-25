package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

var zN *big.Int = big.NewInt(0)
var oN *big.Int = big.NewInt(1)
var tN *big.Int = big.NewInt(2)
var pL = [25]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, //13
	43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97} //12

//how many solving routines to start
var gBN int64 = 1000000

func setBuf(n int64) bool {
	gBN = n
	return true
}

func timeDelay(tMs int64) {
	t := rand.Int63n(tMs)
	time.Sleep(time.Duration(t) * time.Millisecond)
}

func selectPrimeBig(n *big.Int) bool {
	//fmt.Println("selectPrimeBig(",candidate,") called.")
	sC := make(chan bool)
	if isModBig(n) {
		defer fmt.Println("mod false", n)
	}
	go primeIt(n, sC)
	return <-sC
}

func testMod(n int64) bool {
	return isModBig(big.NewInt(n))
}

func isModBig(n *big.Int) bool {
	//fmt.Println("isModBig(",n,",",sPrime,") : ")
	i := int64(2)
	for i < 13 {
		resMod := new(big.Int).Mod(n, tN)
		resCmp := resMod.Cmp(zN)
		if resCmp == 0 {
			return false
		}
		i = nextIntPrime(i)
	}
	return true
}

func nextIntPrime(n int64) int64 {
	//fmt.Println("nextIntPrime(", n, ")")
	switch n {
	case 2:
		return 3
	case 3:
		return 5
	case 5:
		return 7
	case 7:
		return 11
	case 11:
		return 13
	}
	i := 13
	for i < (len(pL) - 1) {
		if pL[i] == n {
			return pL[i+1]
		}
		i = i + 1
	}
	return 2
}

func isPrimeBig(n *big.Int) bool {
	c2 := new(big.Int).SetInt64(2).Sqrt(n)
	p2 := new(big.Int).Add(c2, oN)
	mF := big.NewInt(2)
	mP := big.NewInt(1)
	for p2.Cmp(mF) > 0 {
		new(big.Int).DivMod(n, mF, mP)
		if mP.Cmp(zN) == 0 {
			return false
		}
		mF.Add(mF, oN)
	}
	return true
}

func checkStart(n *big.Int, c chan bool) bool {
	c <- isPrimeBig(n)
	return true
}

func rangePrimeCheck(f *big.Int, c *big.Int, ch chan bool) {
	//fmt.Println("rangePrimeCheck(",floor,",",ceil,",",chReturn,")")
	//range limited by int64 to max(2^64)-1 for bufLimit
	var nL int64 = 1000000 + 1
	r := new(big.Int).Sub(c, f)
	if r.Cmp(new(big.Int).SetInt64(nL)) > 0 {
		fmt.Println("err: range request exceeds f() internal buffer limit")
		return
	}
	rC := make(chan bool, nL)
	i := new(big.Int).Set(f)
	for i.Cmp(c) < 0 {
		go checkStart(i, rC)
		if <-rC {
			fmt.Println(i, "is prime")
		}
		i.Add(i, tN)
	}
	close(rC)
	ch <- true
	return
}

func primeFan(f *big.Int, c chan bool) {
	//fmt.Println("primeFan(",floor,", c )")
	//floor number has to be odd
	fN := new(big.Int).Set(f)
	cN := new(big.Int).Add(fN, big.NewInt(gBN))
	fC := make(chan bool)
	go rangePrimeCheck(fN, cN, fC)
	c <- <-fC
	return
}

func sendIt(n *big.Int) {
	c := make(chan bool)
	if new(big.Int).Mod(n, tN).Cmp(zN) == 0 {
		n = new(big.Int).Add(n, oN)
	}
	go primeFan(n, c)
	if <-c {
		close(c)
	}
	return
}

func primeIt(n *big.Int, c chan bool) bool {
	sendIt(n)
	c <- true
	return true
}

func primeInt(n int64) bool {
	//fmt.Println("primeInt(",n,")")
	c := make(chan bool)
	primeIt(big.NewInt(n), c)
	defer close(c)
	return <-c
}

func primeSeq(n int64, t int64, c chan bool) {
	gBN = t
	go primeInt(n)
	c <- true
	return
}

func testPrimeParallel(n, t, tMs int64) {
	c := make(chan bool)
	go primeSeq(n, t, c)
	timeDelay(tMs) // milliseconds tMilliSec
	fmt.Println("testPrimeParallel(): ", <-c)
	fmt.Println("primeSeq() timeout is working!")
	close(c)
	return
}

func testSelectPrime(n, ms int64) bool {
	//fmt.Println("testSelectPrime(", candidate, ")")
	r := selectPrimeBig(big.NewInt(n))
	timeDelay(ms)
	return r
}

func testPrimeCheckers(candidate, threads, milliseconds int64) bool {
	if testMod(candidate) {
		//fmt.Println(candidate, "?")
		return testSelectPrime(candidate, milliseconds)
	}
	return false
}

func mainTest(n, t, tF int64) chan bool {
	//fmt.Println("mainTest: ", candidate)
	rC := make(chan bool)
	msF := tF * t
	//derive boundaries and timeOuts
	lT := int64(math.Log2(float64(t)))
	lR := int64(math.Log2(float64(t)))
	msPT := (lT+1)*(lR+1)*lR ^ (2 ^ lR)
	msPF := msPT ^ t + 1
	if msPF < tF {
		msPT = msF
	} else if msPF < tF {
		msPT = msF
	}
	//now compute inside the boundaries
	i := int64(0)
	for i < t {
		i += 2
		rC <- testPrimeCheckers((n + i), t, msPT)
	}
	timeDelay(msPT)
	return rC
}

func mainTestFan(n int64) {
	fmt.Print(<-mainTest(n, 1, 100), ":", n, "\n")
	return
}

func mainFun(n int64) {
	i := int64(0)
	r := int64(1)
	for i < r {
		//fmt.Println("mainFun(", candidate+i, ")")
		i += 2
		mainTestFan(n)
		timeDelay(1000)
	}
	return
}

func main() {
	iterations := int64(10000)
	j := int64(8191)
	for j < iterations {
		j += 2
		go mainFun(j)
	}
	timeDelay(60 * 60 * 1000)
}
