package main

import "fmt"

func main() {
	first := 0
	second := 1
	for i := 0; i <= 10; i++ {

			fmt.Println(i)



			fmt.Println("starting the laser")
			var fibo = first + second
			fmt.Println(fibo)

			fmt.Println(second + fibo)
			fmt.Println((second+fibo) + second)
			fmt.Println((second+fibo+second) + (second+fibo))
			fmt.Println(((second+fibo+second) + (second+fibo)) + (second+fibo+second))


		}
}





/*

on 1 print 1 

on 2 print 2 

on 3 print 3

on 4 print 5

on 5 print 8


*/

