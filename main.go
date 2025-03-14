package main

import (
	"fmt"
	"minesweeper/field"
	"minesweeper/input"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	start := time.Now()

	f := field.NewField(8)

	clear()

	endChannel := make(chan bool)

	osInterruptChannel := make(chan os.Signal)
	signal.Notify(osInterruptChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-osInterruptChannel

		endChannel <- true
	}()

	go func() {
		<-endChannel

		fmt.Printf("Your time is %v...\n", time.Since(start))

		fmt.Printf("Please press any key to continue...")
		input.GetInput()

		clear()
		os.Exit(0)
	}()

	for !f.IsFull() {
		f.Display()

		key := input.GetInput()
		if key == input.UP || key == input.LEFT || key == input.DOWN || key == input.RIGHT {
			f.Select(key)
		} else if key == input.UNCOVER {
			if !f.Uncover() {
				clear()
				f.Display()

				fmt.Printf("Game over\n")

				endChannel <- true

				for{}
			}
		} else if key == input.FLAG {
			f.Flag()
		} else {
			fmt.Printf("Wrong keypress\n")
		}

		clear()
	}

	f.Display()

	fmt.Printf("Good job!\n")

	endChannel <- true

	for{}
}
