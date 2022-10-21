package main

import "fmt"

func main() {
	// computes recursively the fibonacci sequence
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d ", fibb(i))
	}
	fmt.Println(" ")

	// computes recursively the factorial of n
	fmt.Println(fact(5))

	// Prints all natural numbers up to 50 recursively
	fmt.Println(natural(50))

	// Sum of numbers from 1 to n recursively
	fmt.Println(sum(5))

	//print the array elements using recursion.
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Print("The elements in the array are: ")
	array1(arr1, 0, 6)
	fmt.Println(" ")

	//Print the sum of digits for 25
	fmt.Printf("The sum of digits for 25 is %d\n", sumDigits(25))
}

func fibb(n int) int {
	if n <= 1 {
		return 1
	}
	return fibb(n-1) + fibb(n-2)
}

func fact(n int) int {
	if n <= 1 {
		return n
	}
	return n * fact(n-1)
}

func natural(n int) int {
	if n <= 1 {
		return n
	}
	fmt.Printf("%d ", n)
	return natural(n - 1)
}

func sum(n int) int {
	if n <= 1 {
		return n
	}
	return n + sum(n-1)
}

func array1(arr1 [6]int, pos, length int) {
	if pos >= length {
		return
	}
	fmt.Printf("%d ", arr1[pos])
	array1(arr1, pos+1, length)
}

func sumDigits(n int) int {
	if n <= 1 {
		return n
	}
	return (n % 10) + sumDigits(n/10)
}
