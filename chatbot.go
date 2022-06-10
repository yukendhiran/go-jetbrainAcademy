package main

import (
	"fmt"
	"strconv"
)

func greet(name, year string) {
	fmt.Println("Hello! My name is " + name + ".")
	fmt.Println("I was created in " + year + ".")
}

func showName() {
	var name string
	fmt.Println("Please, remind me your name.")
	fmt.Scan(&name)
	fmt.Println("What a great name you have, " + name + "!")
}

func guessAge() {
	var rem3, rem5, rem7, age int

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")
}

func count() {
	var n int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		fmt.Printf("%d!\n", i)
	}
}

func startQuiz() {
	fmt.Println("Let's test your programming knowledge.")
	fmt.Println("Why do we use methods?")
	fmt.Println(`1. To repeat a statement multiple times.
				2. To decompose a program into several small subroutines.
				3. To determine the execution time of a program.
				4. To interrupt the execution of a program.`)
	var ans int
	fmt.Scan(&ans)
	for i:=1; i<5; i++ {
		if(i == 2) {
			fmt.Println("Completed, have a nice day!")
		} else {
			fmt.Println("Please, try again.")
		}
	}
	fmt.Println("Completed, have a nice day!")
}

func sayGoodbye() {
	fmt.Println("Congratulations, have a nice day!")
}

func main() {
	greet("Aid", "2020") // change it as you need
	showName()
	guessAge()
	count()
	startQuiz()// call something here
	sayGoodbye()
}
