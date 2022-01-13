package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"banana/mango"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Vegans!")
			mango.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}

