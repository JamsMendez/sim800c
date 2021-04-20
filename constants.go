package sim800c

// Commands SIM800
const (
	hasPINCmd       = "AT+CPIN?\n"
	setPINCmd       = "AT+CPIN=%s\n"
	setGSNCmd       = "AT+GSN\n"
	hasCREGCmd      = "AT+CREG?\n"
	hasCGATTCmd     = "AT+CGATT?\n"
	setCGATTCmd     = "AT+CGATT=1\n"
	setCIPSHUTCmd   = "AT+CIPSHUT\n"
	hasCIPMODECmd   = "AT+CIPMODE?\n"
	setCIPMODECmd   = "AT+CIPMODE=0\n"
	setCIPSTATUSCmd = "AT+CIPSTATUS\n"
	setCIPMUXCmd    = "AT+CIPMUX=0\n"
	setAPNCmd       = "AT+CSTT="
	setCIICRCmd     = "AT+CIICR\n"
	setCIFSRCmd     = "AT+CIFSR\n"
	setCIPSTARTCmd  = "AT+CIPSTART=\"TCP\","
	setCIPSENDCmd   = "AT+CIPSEND\n"
	setCIPCLOSECmd  = "AT+CIPCLOSE\n"
	setJSONCmd      = "[JSON]\r\n"
	setENTERCmd     = "\x1A"

	// Response commands SIM800
	successOk    = "Ok"
	successError = "ERROR"

	successPIN         = "+CPIN: SIM PIN"
	successCREG        = "+CREG: 0,1"
	successCREGZero    = "+CREG: 0,0"
	sucessCGATTON      = "+CGATT: 1"
	sucessCGATTOFF     = "+CGATT: 0"
	successCIPSHUT     = "SHUT OK"
	successMTON        = "+CIPMODE: 1"
	successMTOFF       = "+CIPMODE: 0"
	successIPINITIAL   = "STATE: IP INITIAL"
	successCONNECTOK   = "STATE: CONNECT OK"
	successTCPCLOSED   = "STATE: TCP CLOSED"
	successPDPDEACT    = "STATE: PDP DEACT"
	successCIPSTART    = "CONNECT OK"
	successCIPSEND     = "AT+CIPSEND"
	successSENDOK      = "SEND OK"
	successCLOSEOK     = "CLOSE OK"
	successCLOSED      = "CLOSED"
	successCONNECTFAIL = "CONNECT FAIL"
	pinReady           = "+CPIN: READY"
	callReady          = "Call Ready"
	smsReady           = "SMS Ready"

	msgTimeoutCmd = "COMMAND TIMEOUT: "
	msgErrCmd     = "COMMAND ERROR: "

	sCRNL = "\r\n"
	sNL   = "\n"
)

// New line
const bNL = 10

// values in seconds
const timeoutSend = 60 * 3
const timeoutReady = 90
const timeoutCmd = 30

// value in milliseconds
const delayTime = 250
