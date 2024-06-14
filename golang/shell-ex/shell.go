package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		in, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		clear_in := strings.TrimRight(in, "\n")
		cmds := strings.Split(clear_in, " ")
		switch cmds[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println(strings.Join(cmds[1:], " "))
		case "type":
			switch cmds[1] {
			case "exit", "echo", "type":
				fmt.Printf("%s is a shell builtin\n", cmds[1])
			default:
				paths := strings.Split(os.Getenv("PATH"), ":")
				isFound := false
				for _, path := range paths {
					fullPath := filepath.Join(path, cmds[1])
					if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
						fmt.Printf("%s is %s\n", cmds[1], fullPath)
						isFound = true
						break
					}
				}
				if !isFound {
					fmt.Printf("%s: not found\n", cmds[1])
				}
			}

		default:
			command := exec.Command(cmds[0], cmds[1:]...)
			command.Stderr = os.Stderr
			command.Stdout = os.Stdout
			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", cmds[0])
			}
		}
	}
}
