package debugger

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"time"

	simutex "github.com/nfwGytautas/Simutex/Simutex"
)

type Debugger struct {
	modelUnderDebug simutex.Model
}

func NewDebugger(model simutex.Model) Debugger {
	return Debugger{
		modelUnderDebug: model,
	}
}

func (d *Debugger) Tick() {
	d.modelUnderDebug.Tick()
}

func (d *Debugger) Start(address string) error {
	// Start a socket server
	srv, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Println("Debugger waiting for connection...")

	conn, err := srv.Accept()
	if err != nil {
		return err
	}

	d.handleConnection(conn)

	return nil
}

func (d *Debugger) handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("Starting debugger")

	// Read commands from the connection
	// and execute them

	for {
		// get message, output
		log.Println("Waiting for instruction...")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		if message == "exit\n" {
			log.Println("Exiting debugger")
			break
		}

		if message == "state\n" {
			log.Println("Sending model state")
			state := d.modelUnderDebug.GetState()

			e := json.NewEncoder(conn)

			err = e.Encode(state)
			if err != nil {
				log.Println("Error marshaling model state")
				break
			}
		}

		if message == "tick\n" {
			log.Println("Ticking model")
			d.modelUnderDebug.Tick()

			state := d.modelUnderDebug.GetState()

			e := json.NewEncoder(conn)

			err = e.Encode(state)
			if err != nil {
				log.Println("Error marshaling model state")
				break
			}
		}

		if message == "run\n" {
			log.Println("Running model, 1s = 1tick")
			timer := time.NewTimer(1 * time.Second)
			for {
				select {
				case <-timer.C:
					d.modelUnderDebug.Tick()
					state := d.modelUnderDebug.GetState()

					e := json.NewEncoder(conn)

					err = e.Encode(state)
					if err != nil {
						log.Println("Error marshaling model state")
						return
					}
				}
			}
		}
	}
}
