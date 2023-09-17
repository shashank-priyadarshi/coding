// print primes till n
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type stdReader struct {
	reader  *bufio.Reader
	data    int
	retries int
}

func main() {
	reader := stdReader{
		reader:  bufio.NewReader(os.Stdin),
		data:    0,
		retries: 0,
	}
	fmt.Print("Enter the value of n: ")
	if err := reader.readFromStdInput(); err != nil {
		fmt.Println(fmt.Errorf("%v: retried %d times", err, reader.retries))
		os.Exit(1)
	}
	reader.getPrimesBrute()
	reader.getPrimesSieveOfEratosthenes()
	reader.getPrimesSieveOfSundaram()
	reader.getPrimesSieveOfAtkin()
	reader.getPrimesSegmentedSieves()
	reader.getPrimesMillerRabin()
}

func (s *stdReader) getPrimesBrute() {
	//	iterate till n
	//	for each n, check if it is prime: iterate till sqrt(n)
	//	check number of divisors
	for i := 2; i <= s.data; i++ {
		count := 0
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				count++
				break
			}
		}
		if count < 1 {
			fmt.Printf("%v: %v\n", i, count)
		}
	}
}
func (s *stdReader) getPrimesSieveOfEratosthenes() {
	x := 101
	primeMap := [101]bool{}
	// loop till x, for each number that is false, make all multiples as true
	for key, value := range primeMap {
		newVal := key + 2
		if !value {
			for i := newVal * 2; i <= x; i += newVal {
				fmt.Printf("key: %v, value: %v, newVal: %v, arrayI: %v, primeMap[i]: %v \n", key, value, newVal, i, primeMap[i-2])
				primeMap[i-2] = true
			}
		}
	}
	fmt.Println("------------------------------------------------------------")
	fmt.Println(primeMap)
	for key, value := range primeMap {
		if !value {
			fmt.Println(key+2, key)
		}
	}
	// wg := sync.WaitGroup{}
	//	wg.Add(10)
	//	for i := 0; i < 10; i++ {
	//		go func(i int) {
	//			fmt.Println(i)
	//			wg.Done()
	//		}(i)
	//	}
	//	wg.Wait()
}
func (s *stdReader) getPrimesSieveOfSundaram() {

}
func (s *stdReader) getPrimesSieveOfAtkin() {
}
func (s *stdReader) getPrimesSegmentedSieves() {
}
func (s *stdReader) getPrimesMillerRabin() {
}

func (s *stdReader) readFromStdInput() error {
	if s.retries >= 5 {
		return fmt.Errorf("unable to read your input")
	}
	n, err := s.reader.ReadString(' ')
	fmt.Println("input::: ", n)
	if err != nil {
		fmt.Println("Could not read input. Please try again: ", err)
		s.retries = s.retries + 1
		s.readFromStdInput()
	}
	if s.data, err = strconv.Atoi(n); err != nil {
		fmt.Println("please enter a valid positive number: ", err)
		s.retries = s.retries + 1
		s.readFromStdInput()
	}
	return nil
}
