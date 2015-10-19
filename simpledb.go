package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Struct to store commands
type cmd struct {
	command string
	name    string
	val     int
}

// Struct to store a transaction log
type transaction map[string]int

// Put a limit on size of names
const MAX_NAME_LEN int = 40

func main() {
	var scannee io.Reader
	if len(os.Args) >= 2 {
		// We are dealing with a file
		cmdFile := flag.String("f", "", "File with db commands to read")
		flag.Parse()

		file, err := os.Open(*cmdFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scannee = file
	} else {
		// No file, cmds are on stdin
		printGreeting()
		scannee = os.Stdin
	}

	// Create scanner on input scannee
	scanner := bufio.NewScanner(scannee)
	// Generate lines of possible input
	in := getScannerLines(scanner)
	// Validate lines, generate commands
	cmds := validateCmds(in)
	// Actually run the commands
	runCmds(cmds)
	for cmd := range cmds {
		fmt.Println(cmd)
	}
}

// printGreeting prints a greeting when program starts
func printGreeting() {
	fmt.Println("Simple db by Ryan Ernst")
	fmt.Println("For list of commands enter: HELP")
}

// printHelp prints a helpful message when requested by the user.
func printHelp() {
	fmt.Println("This is a help message")
}

// getScanner generates string arrays in a go routine and
// passes the results out through a channel to continue down the
// pipeline
func getScannerLines(scanner *bufio.Scanner) chan []string {
	out := make(chan []string)
	go func() {
		for scanner.Scan() {
			// Get tokens, remove extra white space
			tokens := strings.Fields(scanner.Text())
			cmd := strings.ToUpper(tokens[0])
			if cmd == "HELP" {
				printHelp()
			} else if cmd == "END" {
				break
			} else {
				out <- tokens
			}
		}
		close(out)
	}()
	return out
}

// validateCmds recieves a string array of tokens, validates them, and then
// packs each command into a cmd struct. The resulting cmd structs are
// then passed out via a channel.
func validateCmds(in chan []string) chan cmd {
	out := make(chan cmd)
	go func() {
		for tokens := range in {
			var c cmd
			var err error

			// Get command
			c.command = strings.ToUpper(tokens[0])
			switch c.command {
			case "BEGIN", "ROLLBACK", "COMMIT", "END":
				if len(tokens) > 1 {
					fmt.Printf("ERROR: %v does not take any arguments\n", c.command)
				} else {
					out <- c
				}

			case "SET":
				// Check number of args
				if len(tokens) != 3 {
					fmt.Println("ERROR: SET usage: SET name value")
					break
				}

				// Check length of name
				if len(tokens[1]) > MAX_NAME_LEN {
					fmt.Printf("ERROR: Max name length is %v\n", MAX_NAME_LEN)
					break
				} else {
					c.name = tokens[1]
				}

				// Try atoi, store if good else spit out error
				c.val, err = strconv.Atoi(tokens[2])
				if err != nil {
					fmt.Println("ERROR: Invalid integer value provided")
					break
				} else {
					out <- c
				}

			case "GET", "UNSET":
				// Check number of args
				if len(tokens) != 2 {
					fmt.Printf("ERROR: %v usage: %v name\n", c.command, c.command)
					break
				}

				// Check length of name
				if len(tokens[1]) > MAX_NAME_LEN {
					fmt.Printf("ERROR: Max name length is %v\n", MAX_NAME_LEN)
					break
				} else {
					c.name = tokens[1]
				}
				out <- c

			case "NUMEQUALTO":
				// Check number of args
				if len(tokens) != 2 {
					fmt.Printf("ERROR: %v usage: %v value\n", c.command, c.command)
					break
				}

				// Try atoi, store if good else spit out error
				c.val, err = strconv.Atoi(tokens[1])
				if err != nil {
					fmt.Println("ERROR: Invalid integer value provided")
					break
				} else {
					out <- c
				}

			default:
				fmt.Printf("ERROR: %v is not a recognized command\n", c.command)
			}

		}
		close(out)
	}()
	return out
}

// runCmds actually runs the commands.
func runCmds(in chan cmd) chan string {

}
