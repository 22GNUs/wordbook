package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/22GNUs/wordbook/cfg"
	"github.com/22GNUs/wordbook/eduic"
	"github.com/fatih/color"
)

type cmd string

const (
	add cmd = "add"
	lst cmd = "lst"
	del cmd = "del"
)

var authProvider = func() string {
	c := cfg.Read()
	return c.Auth
}

func main() {
	// get command line arguements
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: workbook <cmd> <args>")
		return
	}
	cmd := cmd(args[0])
	c := eduic.NewClient(authProvider)
	switch cmd {
	case add:
		words := args[1:]
		if len(words) == 0 {
			fmt.Println("Please input least one word to add")
			return
		}
		msg, err := c.AddWords(args[1:]...)
		if err != nil {
			fmt.Printf("Error occurd when addWords: %s\n", err)
			return
		}

		fmt.Println(msg)
	case del:
		words := args[1:]
		if len(words) == 0 {
			fmt.Println("Please input least one word to add")
			return
		}
		err := c.DelWords(args[1:]...)
		if err != nil {
			fmt.Printf("Error occurd when delWords: %s\n", err)
			return
		}
	case lst:
		extraArgs := args[1:]
		page, size := 0, 10
		var err error
		if len(extraArgs) > 0 {
			page, err = strconv.Atoi(extraArgs[0])
			if err != nil {
				fmt.Println("Please input 'page' as number")
				return
			}
		}
		if len(extraArgs) > 1 {
			size, err = strconv.Atoi(extraArgs[1])
			if err != nil {
				fmt.Println("Please input 'pageSize' as number")
				return
			}
		}
		words, err := c.ListWords(page, size)
		if err != nil {
			fmt.Printf("Error occurd when ListWords: %s\n", err)
			return
		}
		printWords(words)
	default:
		fmt.Printf("Cmd: %s is not support", cmd)
	}
}

func printWords(words []eduic.Explain) {
	for _, w := range words {
		color.Magenta("  " + w.Word)
		exps := strings.Split(w.Exp, "<br>")
		for _, e := range exps {
			if len(e) != 0 {
				// right icon: fb0c
				color.Yellow("﬌ " + e)
			}
		}
	}
}
