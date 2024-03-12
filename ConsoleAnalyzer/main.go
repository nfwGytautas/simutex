package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	simutex "github.com/nfwGytautas/Simutex/Simutex"
)

func main() {
	// Start a socket client
	conn, err := net.Dial("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	for {
		// Wait for user input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Command: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Write to server
		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Fatal(err)
		}

		if text == "run" {
			for {
				// Clear
				fmt.Print("\033[H\033[2J")

				d := json.NewDecoder(conn)

				state := simutex.FrozenState{}
				err = d.Decode(&state)
				if err != nil {
					log.Fatal(err)
				}

				for _, value := range state.Entries {
					fmt.Println(value)
				}
			}
		} else {
			// Clear
			fmt.Print("\033[H\033[2J")

			d := json.NewDecoder(conn)

			state := simutex.FrozenState{}
			err = d.Decode(&state)
			if err != nil {
				log.Fatal(err)
			}

			for _, value := range state.Entries {
				fmt.Println(value)
			}
		}
	}
}
