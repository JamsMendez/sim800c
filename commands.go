package sim800c

import (
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
					message = msgErrCmd + strings.ReplaceAll(hasPINCmd, "\n", "")

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
			message = msgTimeoutCmd + strings.ReplaceAll(hasPINCmd, "\n", "")

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
					message = msgErrCmd + strings.ReplaceAll(setPINCmd, "\n", "")

					break

				} else if isReady && isCall && isSMS {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setPINCmd, "\n", "")

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
				isOk := isIMEI(client.lines, client.IMEI)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setGSNCmd, "\n", "")

					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setGSNCmd, "\n", "")

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
					message = msgErrCmd + strings.ReplaceAll(hasCREGCmd, "\n", "")

					break

				} else if isSuccess && isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(hasCREGCmd, "\n", "")

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
					message = msgErrCmd + strings.ReplaceAll(hasCGATTCmd, "\n", "")
					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(hasCGATTCmd, "\n", "")

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
					message = msgErrCmd + strings.ReplaceAll(setCGATTCmd, "\n", "")
					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCGATTCmd, "\n", "")

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
					message = msgErrCmd + strings.ReplaceAll(hasCIPMODECmd, "\n", "")

					break

				} else if isOk && isModeOne {
					break

				} else if isOk && isModeZero {
					break
				}
			}

			timeoutCount = timeoutCount + 1
		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(hasCIPMODECmd, "\n", "")

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
					message = msgErrCmd + strings.ReplaceAll(setCIPMODECmd, "\n", "")
					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1
		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPMODECmd, "\n", "")

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
					message = msgErrCmd + strings.ReplaceAll(setCIPSHUTCmd, "\n", "")

					break

				} else if isSuccess {
					break
				}
			}

			timeoutCount = timeoutCount + 1
		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPSHUTCmd, "\n", "")
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
					message = msgErrCmd + strings.ReplaceAll(setCIPSTATUSCmd, "\n", "")

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
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPSTATUSCmd, "\n", "")

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
					message = msgErrCmd + strings.ReplaceAll(setAPN, "\n", "")

					break

				} else if isOk {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setAPN, "\n", "")

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
				message = msgErrCmd + strings.ReplaceAll(setCIICRCmd, "\n", "")

				break

			} else if isOk {
				break
			}
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
					message = msgErrCmd + strings.ReplaceAll(setCIFSRCmd, "\n", "")

					break

				} else if isSuccess {
					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIFSRCmd, "\n", "")
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
					message = msgErrCmd + strings.ReplaceAll(setCIPSTARTCmd, "\n", "")
					break

				} else if isSuccess {
					break

				} else if isFail {
					message = successCONNECTFAIL + ": " + strings.ReplaceAll(nCIPSTARTCmd, "\n", "")

					break
				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(nCIPSTARTCmd, "\n", "")
			break
		}
	}

	response <- message
}

func (client *ClientTCP) waitCIPSENDCmd(timer *time.Ticker, response chan<- string) {
	var message string

	for range timer.C {
		length := len(client.lines)
		if length > 0 {
			isError := indexOf(client.lines, successError)
			isClosed := indexOf(client.lines, successCLOSED)
			isSuccess := indexOf(client.lines, successCIPSEND)

			if isError {
				message = msgErrCmd + strings.ReplaceAll(setCIPSENDCmd, "\n", "")
				break

			} else if isSuccess {
				break

			} else if isClosed {
				message = msgErrCmd + successCLOSED

				break
			}
		}
	}

	response <- message
}

func (client *ClientTCP) waitENTERCmd(timer *time.Ticker, response chan<- string) {
	var message string

	for range timer.C {
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
				isClosed := indexOf(client.lines, successCLOSED)

				if isError {
					message = msgErrCmd + strings.ReplaceAll(setCIPCLOSECmd, "\n", "")

					break

				} else if isClosed {
					break

				}
			}

			timeoutCount = timeoutCount + 1

		} else {
			message = msgTimeoutCmd + strings.ReplaceAll(setCIPCLOSECmd, "\n", "")

			break
		}
	}

	response <- message
}
