package sim800c

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/tarm/serial"
)

// ClientTCP ...  Client TCP
type ClientTCP struct {
	PortName string
	IP       string
	Port     string
	APN      string
	User     string
	Password string
	IMEI     string
	BaudRate int
	PinKey   int
	PinSIM   string

	Connected bool
	Debug     bool

	// SerialPort
	serialPort *serial.Port

	// Slices of responses
	lines  []string
	inputs []string

	// Channel serial data
	data        chan string
	isConnected bool
	insertPIN   bool
}

// reading ... Receive messages
func (client *ClientTCP) reading() {
	go func() {
		line := []byte{}

		for {
			buffer := make([]byte, 64)
			n, err := client.serialPort.Read(buffer)
			if err != nil {
				if client.Debug {
					s := fmt.Sprintf("ClientTCP.SerialPort.Read.ERROR: %v", err)
					nErr := errors.New(s)

					printErrCmd(nErr)
				}

				client.Close()

				break
			}

			chunk := buffer[:n]
			size := len(chunk)
			for i := 0; i < size; i++ {
				line = append(line, chunk[i])

				if chunk[i] == bNL {
					var parts []string
					s := string(line)

					isCRNL := strings.Contains(s, sCRNL)
					if isCRNL {
						parts = strings.Split(s, sCRNL)
					} else {
						parts = strings.Split(s, sNL)
					}

					if len(parts) > 0 {
						first := parts[0]
						if first != "" {
							client.data <- first
						}

						line = []byte{}
					}
				}
			}
		}
	}()

	for line := range client.data {
		client.addJSON(line)

		client.lines = append(client.lines, line)
	}
}

// addJSON ... check if a message is JSON to get it with the GetJSON
func (client *ClientTCP) addJSON(s string) {
	o := map[string]interface{}{}
	a := []interface{}{}

	buffer := []byte(s)
	err := json.Unmarshal(buffer, &o)
	if err == nil {
		if client.Debug {
			printOutputCmd([]string{s})
		}

		client.inputs = append(client.inputs, s)
	}

	if err != nil {
		err := json.Unmarshal(buffer, &a)
		if err == nil {
			if client.Debug {
				printOutputCmd([]string{s})
			}

			client.inputs = append(client.inputs, s)
		}
	}
}

func (client *ClientTCP) commandExec(command string) error {
	/* ==== Command SIM800C ==== */
	if client.Debug {
		printInputCmd(command)
	}

	buffer := []byte(command)
	_, err := client.serialPort.Write(buffer)

	if err != nil {
		s := fmt.Sprintf("ClientTCP.SerialPort.Write.ERROR: %v", err)
		nErr := errors.New(s)

		printErrCmd(nErr)
	}

	return err
}

func (client *ClientTCP) isClosed(ch chan string) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}

// Open ... Configuration for connection serial port
func (client *ClientTCP) Open() (err error) {
	// Configuration for serial port
	options := &serial.Config{Name: client.PortName, Baud: client.BaudRate}

	// Open serial port
	client.serialPort, err = serial.OpenPort(options)
	if err != nil {
		if client.Debug {
			s := fmt.Sprintf("ClientTCP.SerialPort.Open.ERROR: %v", err)
			nErr := errors.New(s)

			printErrCmd(nErr)
		}

		time.Sleep(time.Second * 1)
	}

	client.data = make(chan string)
	client.isConnected = true

	go client.reading()

	return err
}

// Connect ... Configuration to create a TCP Client
func (client *ClientTCP) Connect() (err error) {
	if !client.isConnected {
		message := "serial port isn't open"
		err = errors.New(message)
		return err
	}

	// === Inicialize SIM800C and/or registerPIN ===
	client.lines = make([]string, 0)
	err = client.commandExec(hasPINCmd)
	if err != nil {
		return err
	}

	response := make(chan string)
	timer := time.NewTicker(time.Millisecond * delayTime)

	defer close(response)

	go client.waitHasReadyCmd(timer, response)

	message := <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.insertPIN {
		if client.Debug {
			s := "Insert PIN starting ..."
			printOutputCmd([]string{s})
		}

		if client.Debug {
			printOutputCmd(client.lines)
		}

		// === SET PIN COMMAND ===
		client.lines = make([]string, 0)

		nPINCmd := fmt.Sprintf(setPINCmd, client.PinSIM)
		err = client.commandExec(nPINCmd)
		if err != nil {
			return err
		}

		timer := time.NewTicker(time.Millisecond * delayTime)

		go client.waitPINCmd(timer, response)

		message := <-response

		timer.Stop()

		if message != "" {
			if client.Debug {
				printOutputCmd(client.lines)
			}

			return errors.New(message)
		}
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	// === GSN COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(setGSNCmd)
	if err != nil {
		return err
	}

	timer = time.NewTicker(time.Millisecond * delayTime)

	go client.waitGSNCmd(timer, response)

	message = <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	// === CREG COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(hasCREGCmd)
	if err != nil {
		return err
	}

	timer = time.NewTicker(time.Millisecond * delayTime)

	go client.waitHasCREGCmd(timer, response)

	message = <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	// === CGATT COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(hasCGATTCmd)
	if err != nil {
		return err
	}

	timer = time.NewTicker(time.Millisecond * delayTime)

	go client.waitHasCGATTCmd(timer, response)

	message = <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	// === SET CGATT COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(setCGATTCmd)
	if err != nil {
		return err
	}

	timer = time.NewTicker(time.Millisecond * delayTime)

	go client.waitCGATTCmd(timer, response)

	message = <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	// === SET CIPSTATUS COMMAND ===
	var hasStatusTCPClosed bool

	client.lines = make([]string, 0)
	err = client.commandExec(setCIPSTATUSCmd)
	if err != nil {
		return err
	}

	timer = time.NewTicker(time.Millisecond * delayTime)

	go client.waitCIPSTATUSCmd(timer, response)

	message = <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		if message == successCONNECTOK {
			return err
		}

		if message == successTCPCLOSED {
			hasStatusTCPClosed = true
		}

	} else {
		s := fmt.Sprintf("%s RESPONSE EMPTY", msgErrCmd)
		err = errors.New(s)
	}

	if !hasStatusTCPClosed {
		// === SET CIPSHUT COMMAND ===
		client.lines = make([]string, 0)
		err = client.commandExec(setCIPSHUTCmd)
		if err != nil {
			return err
		}

		timer = time.NewTicker(time.Millisecond * delayTime)

		go client.waitCIPSHUTCmd(timer, response)

		message = <-response

		timer.Stop()

		if message != "" {
			if client.Debug {
				printOutputCmd(client.lines)
			}

			return errors.New(message)
		}

		if client.Debug {
			printOutputCmd(client.lines)
		}

		// === SET HAS CIPMODE COMMAND ===
		client.lines = make([]string, 0)
		err = client.commandExec(hasCIPMODECmd)
		if err != nil {
			return err
		}

		timer = time.NewTicker(time.Millisecond * delayTime)

		go client.waitHasCIPMODECmd(timer, response)

		message = <-response

		timer.Stop()

		if message != "" {
			if client.Debug {
				printOutputCmd(client.lines)
			}

			return errors.New(message)
		}

		if client.Debug {
			printOutputCmd(client.lines)
		}

		// === SET CIPMODE COMMAND ===
		client.lines = make([]string, 0)
		err = client.commandExec(setCIPMODECmd)
		if err != nil {
			return err
		}

		timer = time.NewTicker(time.Millisecond * delayTime)

		go client.waitCIPMODECmd(timer, response)

		message = <-response

		timer.Stop()

		if message != "" {
			if client.Debug {
				printOutputCmd(client.lines)
			}

			return errors.New(message)
		}

		if client.Debug {
			printOutputCmd(client.lines)
		}

		// === SET CIPSTATUS COMMAND ===
		client.lines = make([]string, 0)
		err = client.commandExec(setCIPSTATUSCmd)
		if err != nil {
			return err
		}

		timer = time.NewTicker(time.Millisecond * delayTime)

		go client.waitCIPSTATUSCmd(timer, response)

		message = <-response

		timer.Stop()

		if message != "" {
			if client.Debug {
				printOutputCmd(client.lines)
			}

			if message != successIPINITIAL {
				err = errors.New(message)

				return err
			}

		} else {
			s := fmt.Sprintf("%s RESPONSE EMPTY", msgErrCmd)
			err = errors.New(s)
		}

		// === SET APN COMMAND	===
		setAPN := fmt.Sprintf("%s\"%s\",\"%s\",\"%s\"%s", setAPNCmd, client.APN, client.User, client.Password, sNL)
		client.lines = make([]string, 0)
		err = client.commandExec(setAPN)
		if err != nil {
			return err
		}

		timer = time.NewTicker(time.Millisecond * delayTime)

		go client.waitAPNCmd(timer, response, setAPN)

		message = <-response

		timer.Stop()

		if message != "" {
			if client.Debug {
				printOutputCmd(client.lines)
			}

			return errors.New(message)
		}

		if client.Debug {
			printOutputCmd(client.lines)
		}

		// === SET CIICR COMMAND ===
		client.lines = make([]string, 0)
		err = client.commandExec(setCIICRCmd)
		if err != nil {
			return err
		}

		timer = time.NewTicker(time.Millisecond * delayTime)

		go client.waitCIICRCmd(timer, response)

		message = <-response

		timer.Stop()

		if message != "" {
			if client.Debug {
				printOutputCmd(client.lines)
			}

			return errors.New(message)
		}

		if client.Debug {
			printOutputCmd(client.lines)
		}

		// === SET CIFSR COMMAND ===
		client.lines = make([]string, 0)
		err = client.commandExec(setCIFSRCmd)
		if err != nil {
			return err
		}

		timer = time.NewTicker(time.Millisecond * delayTime)

		go client.waitCIFSRCmd(timer, response)

		message = <-response

		timer.Stop()

		if message != "" {
			if client.Debug {
				printOutputCmd(client.lines)
			}

			return errors.New(message)
		}

		if client.Debug {
			printOutputCmd(client.lines)
		}
	}

	// === SET CIPSTART COMMAND ===
	client.lines = make([]string, 0)
	nCIPSTART := fmt.Sprintf("%s\"%s\",\"%s\"%s", setCIPSTARTCmd, client.IP, client.Port, sNL)
	err = client.commandExec(nCIPSTART)
	if err != nil {
		return err
	}

	timer = time.NewTicker(time.Millisecond * delayTime)

	go client.waitCIPSTARTCmd(timer, response, nCIPSTART)

	message = <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	return err
}

// IsConnect... Get TCP Client is connected
func (client *ClientTCP) IsConnect(msg string) (isOk bool, err error) {
	if !client.isConnected {
		message := "serial port isn't open"
		err = errors.New(message)
		return isOk, err
	}

	// === SET CIPCLOSE COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(setCIPSTATUSCmd)
	if err != nil {
		return isOk, err
	}

	response := make(chan string)
	defer close(response)

	timer := time.NewTicker(time.Millisecond * delayTime)

	go client.waitCIPSTATUSCmd(timer, response)

	message := <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		if message == successCONNECTOK {
			isOk = true

			return isOk, err
		}

		if message == successIPINITIAL || message == successCONNECTFAIL {
			return isOk, err
		}

		err = errors.New(message)

	} else {
		s := fmt.Sprintf("%s RESPONSE EMPTY", msgErrCmd)
		err = errors.New(s)
	}

	return isOk, err
}

// Disconnect... Close connection TCP client
func (client *ClientTCP) Disconnect() (err error) {
	if !client.isConnected {
		message := "serial port isn't open"
		err = errors.New(message)
		return err
	}

	// === SET CIPCLOSE COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(setCIPCLOSECmd)
	if err != nil {
		return err
	}

	response := make(chan string)
	defer close(response)

	timer := time.NewTicker(time.Millisecond * delayTime)

	go client.waitCIPCLOSECmd(timer, response)

	message := <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	return err
}

// Reset... reset connection and configuration TCP client
func (client *ClientTCP) Reset() (err error) {
	if !client.isConnected {
		message := "serial port isn't open"
		err = errors.New(message)
		return err
	}

	// === SET CIPCLOSE COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(setCIPCLOSECmd)
	if err != nil {
		return err
	}

	response := make(chan string)
	defer close(response)

	timer := time.NewTicker(time.Millisecond * delayTime)

	go client.waitCIPCLOSECmd(timer, response)

	message := <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		// Error ignored because there may be no connection
	}

	// === SET CIPSHUT COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(setCIPSHUTCmd)
	if err != nil {
		return err
	}

	timer = time.NewTicker(time.Millisecond * delayTime)

	go client.waitCIPSHUTCmd(timer, response)

	message = <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	return err
}

// Send ... Send data packet to TCP server
func (client *ClientTCP) Send(msg string) (err error) {
	if !client.isConnected {
		message := "serial port isn't open"
		err = errors.New(message)
		return err
	}

	// === SET CIPSEND COMMAND ===
	client.lines = make([]string, 0)
	err = client.commandExec(setCIPSENDCmd)
	if err != nil {
		return err
	}

	response := make(chan string)
	defer close(response)

	timer := time.NewTicker(time.Millisecond * delayTime)

	go client.waitCIPSENDCmd(timer, response)

	message := <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	// === Insert data >_ ===
	setJSON := strings.ReplaceAll(setJSONCmd, "[JSON]", msg)
	client.lines = make([]string, 0)
	err = client.commandExec(setJSON)
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * delayTime)

	// === ENTER send data ===
	client.lines = make([]string, 0)
	err = client.commandExec(setENTERCmd)
	if err != nil {
		return err
	}

	timer = time.NewTicker(time.Millisecond * delayTime)

	go client.waitENTERCmd(timer, response)

	message = <-response

	timer.Stop()

	if message != "" {
		if client.Debug {
			printOutputCmd(client.lines)
		}

		return errors.New(message)
	}

	if client.Debug {
		printOutputCmd(client.lines)
	}

	return err
}

// Close ... Close serial port
func (client *ClientTCP) Close() {
	if client.isConnected {
		client.isConnected = false
	}

	if !client.isClosed(client.data) {
		if client.data != nil {
			close(client.data)
		}
	}

	if client.serialPort != nil {
		err := client.serialPort.Close()
		if err != nil {
			if client.Debug {
				fmt.Println("ClientTCP.SerialPort.Close.ERROR: ", err)
			}
		}
	}
}

// GetJSON ... Gets the JSONs received from server TCP
func (client *ClientTCP) GetJSON() []string {
	defer func() {
		client.inputs = []string{}
	}()

	return client.inputs
}
