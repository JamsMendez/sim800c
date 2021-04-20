package sim800c

import (
	"fmt"
	"strings"
	"time"
)

func (client *ClientTCP) waitHasReadyCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutReady * 1000 / delayTime
		if timeoutCount < timeoutMs {
			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				hasPIN := indexOf(client.lines, successPIN)
				isReady := indexOf(client.lines, pinReady)
				isOk := indexOf(client.lines, successOk)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(hasPINCmd, sNL, "")

					break

				} else if hasPIN && isOk {
					client.insertPIN = true
					break

				} else if isReady && isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(hasPINCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitPINCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutReady * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isReady := indexOf(client.lines, pinReady)
				isCall := indexOf(client.lines, callReady)
				isSMS := indexOf(client.lines, smsReady)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setPINCmd, sNL, "")

					break

				} else if isReady && isCall && isSMS {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setPINCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitGSNCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isIMEI := isIMEI(client.lines, client.IMEI)
				isOk := indexOf(client.lines, successOk)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setGSNCmd, sNL, "")

					break

				} else if isOk || isIMEI {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setGSNCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitHasCREGCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutReady * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isSuccess := indexOf(client.lines, successCREG)
				isOk := indexOf(client.lines, successOk)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(hasCREGCmd, sNL, "")

					break

				} else if isSuccess && isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(hasCREGCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitHasCGATTCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isOk := indexOf(client.lines, successOk)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(hasCGATTCmd, sNL, "")
					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(hasCGATTCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCGATTCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isOk := indexOf(client.lines, successOk)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCGATTCmd, sNL, "")
					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCGATTCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitHasCIPMODECmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {
			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isModeZero := indexOf(client.lines, successMTOFF)
				isModeOne := indexOf(client.lines, successMTON)
				isOk := indexOf(client.lines, successOk)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(hasCIPMODECmd, sNL, "")

					break

				} else if isOk && isModeOne {
					break

				} else if isOk && isModeZero {
					break
				}
			}

			timeoutCount = timeoutCount + 1
		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(hasCIPMODECmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIPMODECmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {
			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isOk := indexOf(client.lines, successOk)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCIPMODECmd, sNL, "")
					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1
		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPMODECmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIPSHUTCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutReady * 1000 / delayTime
		if timeoutCount < timeoutMs {
			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isSuccess := indexOf(client.lines, successCIPSHUT)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCIPSHUTCmd, sNL, "")

					break

				} else if isSuccess {
					break
				}
			}

			timeoutCount = timeoutCount + 1
		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPSHUTCmd, sNL, "")
			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIPSTATUSCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {
			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isClosed := indexOf(client.lines, successTCPCLOSED)
				isInitial := indexOf(client.lines, successIPINITIAL)
				isConnect := indexOf(client.lines, successCONNECTOK)
				isPDP := indexOf(client.lines, successPDPDEACT)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCIPSTATUSCmd, sNL, "")

					break

				} else if isInitial {
					message = successIPINITIAL

					break

				} else if isConnect {
					message = successCONNECTOK

					break

				} else if isClosed {
					message = successTCPCLOSED

					break

				} else if isPDP {
					message = successPDPDEACT

					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPSTATUSCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitAPNCmd(timer *time.Ticker, response chan<- string, setAPN string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isOk := indexOf(client.lines, successOk)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setAPN, sNL, "")

					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setAPN, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIICRCmd(timer *time.Ticker, response chan<- string) {
	var message string

	for range timer.C {
		length := len(client.lines)
		if length > 0 {
			isError := indexOf(client.lines, successError)
			isOk := indexOf(client.lines, successOk)

			if isError {
				message = msgErrCmd + strings.ReplaceAll(setCIICRCmd, sNL, "")

				break

			} else if isOk {
				break
			}
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIFSRCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isSuccess := isIP(client.lines)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCIFSRCmd, sNL, "")

					break

				} else if isSuccess {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIFSRCmd, sNL, "")
			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIPSTARTCmd(timer *time.Ticker, response chan<- string, nCIPSTARTCmd string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutReady * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isFail := indexOf(client.lines, successCONNECTFAIL)
				isSuccess := indexOf(client.lines, successCIPSTART)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCIPSTARTCmd, sNL, "")
					break

				} else if isSuccess {
					break

				} else if isFail {
					message = successCONNECTFAIL + ": " + strings.ReplaceAll(nCIPSTARTCmd, sNL, "")

					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(nCIPSTARTCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIPSENDCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutSend * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isClosed := indexOf(client.lines, successCLOSED)
				isSuccess := indexOf(client.lines, successCIPSEND)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCIPSENDCmd, sNL, "")
					break

				} else if isSuccess {
					break

				} else if isClosed {
					message = msgErrCmd + successCLOSED

					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPSENDCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitENTERCmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutSend * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isClosed := indexOf(client.lines, successCLOSED)
				isSuccess := indexOf(client.lines, successSENDOK)

				if isError {
					message = msgErrCmd + "ENTER SEND JSON"
					break

				} else if isSuccess {
					break

				} else if isClosed {
					message = msgErrCmd + successCLOSED

					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setENTERCmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIPCLOSECmd(timer *time.Ticker, response chan<- string) {
	timeoutCount := 0
	var message string

	for range timer.C {
		timeoutMs := timeoutCmd * 1000 / delayTime
		if timeoutCount < timeoutMs {

			length := len(client.lines)
			if length > 0 {
				isError := indexOf(client.lines, successError)
				isClosed := indexOf(client.lines, successCLOSEOK)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCIPCLOSECmd, sNL, "")

					break

				} else if isClosed {
					break

				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPCLOSECmd, sNL, "")

			break
		}

		if !client.isConnected {
			message = fmt.Sprintf("%s %s", msgErrCmd, "serial port is closed")

			break
		}
	}

	response <- message
}
