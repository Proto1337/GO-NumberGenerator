package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear") // clear POSIX
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // clear nt
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// tidy console
func ClearConsole() {
	value, check := clear[runtime.GOOS]
	if check {
		value()
	} else { //unsupported platform
		panic("Platform is not supported!")
	}
}

// check if int i is in slice s
func contains(s []int, i int) bool {
	for _, x := range s {
		if x == i {
			return true
		}
	}
	return false
}

func main() {
	ClearConsole()

	fmt.Println("--------------------------------")
	fmt.Println("--- program written by proto ---")
	fmt.Println("--------------------------------")

	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	counter := 0

	// upper border for numbers
	fmt.Print("Set max number to generate :: ")
	input, _, _ := reader.ReadLine()
	in := string(input)
	number, _ := strconv.Atoi(in) // console in -> int
	current := 0                  // current helper
	var pulled []int              // slice documents pulled numbers
	ClearConsole()

	for {
		fmt.Println()
		fmt.Println("What to do?")
		fmt.Println("' .exit ' to end program")
		fmt.Println("' clear ' to clear number history")
		fmt.Println("' history ' to view generated numbers")
		fmt.Println("' generate ' to generate numbers")
		fmt.Println()

		fmt.Print("$> ")
		input, _, _ := reader.ReadLine()
		in = string(input)

		if in == ".exit" {
			// end Programm
			fmt.Println("|- ending program -|")
			time.Sleep(time.Second)
			ClearConsole()
			os.Exit(0)
		} else if in == "generate" {
			// generate numbers if possible
			for {
				fmt.Println()
				fmt.Println("Continuing deletes all current output!")
				fmt.Println("Enter to continue...")
				_, _, _ = reader.ReadLine()
				ClearConsole()
				// generate
				fmt.Println("How many numbers shall be generated?")
				fmt.Println("' .exit ' to go back to main menu")
				fmt.Print(":-> ")

				input, _, _ := reader.ReadLine()
				in = string(input)

				// cancel for
				if in == ".exit" {
					break
				}

				counter, _ = strconv.Atoi(in)

				for iterator := 0; iterator < counter; iterator++ {
					if len(pulled) == number {
						// slice full -> end
						fmt.Println("! All possible numbers have been generated !")
						break
					} else {
						// generate rand number
						// if in list already -> again
						current = rand.Intn(number) + 1
						if contains(pulled, current) {
							iterator--
						} else {
							fmt.Println("Number :: ", current)
							pulled = append(pulled, current)
						}
					}
				}
			}

			fmt.Println("Enter to continue...")
			_, _, _ = reader.ReadLine()

			ClearConsole()

		} else if in == "history" {
			// show generated numbers in their order
			ClearConsole()

			fmt.Println("Generated numbers:")
			fmt.Println(pulled)
			fmt.Println()
			fmt.Println("Enter to continue...")
			_, _, _ = reader.ReadLine()

			ClearConsole()
		} else if in == "clear" {
			// clear slice
			ClearConsole()
			pulled = nil
			ClearConsole()
		} else {
			// unknown input -> again
			ClearConsole()

			fmt.Println("invalid input")

			fmt.Println("Enter to continue...")
			_, _, _ = reader.ReadLine()

			ClearConsole()
		}

	}

}
