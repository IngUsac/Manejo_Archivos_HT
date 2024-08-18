package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bienvenido al simulador de comandos. Escriba 'execute archi_script.txt'  o  'exit' para salir.")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if input == "exit" {
			break
		}
		processCommand(input)
	}
}

func processCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	command := strings.ToLower(parts[0])
	args := parts[1:]

	switch command {
	case "execute":
		executeCommand(args)
	case "mkdisk":
		mkdiskCommand(args)
	case "rep":
		repCommand(args)
	default:
		fmt.Println("Comando no reconocido:", command)
	}
}
