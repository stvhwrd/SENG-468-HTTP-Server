package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	commonlib "github.com/kurtd5105/SENG-468-Common-Lib"
)

// DebugType contains all the information of user commands, in
// addition to an optional debug message
func buildLogDebug(logParameters commonlib.LogCommandParameter) (uint8, commonlib.LogCommandParameter) {
	// Required fields: Timestamp, Server, TransactionNum, Command
	// Optional fields: Username, StockSymbol, Filename, Funds, DebugMessage

	log.Println(logParameters.DebugMessage)

	logParameters.Timestamp = getUnixTimestampAsString()
	logParameters.Server = "Web"

	// logParameters := commonlib.LogCommandParameter{
	// 	Timestamp:      getUnixTimestampAsString(),
	// 	Server:         "Web",
	// 	TransactionNum: "0001",
	// 	Command:        "QUOTE",
	// 	DebugMessage:   debugMessage,
	// }

	commandID := uint8(commonlib.DebugType)

	return commandID, logParameters
}

// ErrorEventType contains all the information of user commands, in
// addition to an optional error message
func buildLogErrorEvent() (uint8, commonlib.LogCommandParameter) {
	// Required fields: Timestamp, Server, TransactionNum, Command
	// Optional fields: Username, StockSymbol, Filename, Funds, ErrorMessage

	logParameters := commonlib.LogCommandParameter{
		Timestamp:      getUnixTimestampAsString(),
		Server:         "Web",
		TransactionNum: "0002",
		Command:        "SET_BUY_TRIGGER",
	}

	logCommandID := uint8(commonlib.ErrorEventType)

	return logCommandID, logParameters
}

// SystemEventTypes can be current user commands, interserver communications,
// or the execution of previously set triggers.
func buildLogSystemEvent() (uint8, commonlib.LogCommandParameter) {
	// Required fields: Timestamp, Server, TransactionNum, Command
	// Optional fields: Username, StockSymbol, Filename, Funds
	logParameters := commonlib.LogCommandParameter{
		Timestamp:      getUnixTimestampAsString(),
		Server:         "Web",
		TransactionNum: "0003",
		Command:        "DUMPLOG",
	}

	logCommandID := uint8(commonlib.SystemEventType)

	return logCommandID, logParameters
}

// UserCommandType comes from the user command files via Workload Generator
// or from manual entries in the UI
func buildLogUserCommand() (uint8, commonlib.LogCommandParameter) {
	// Required fields: Timestamp, Server, TransactionNum, Command
	// Optional fields: Username, StockSymbol, Filename, Funds
	logParameters := commonlib.LogCommandParameter{
		Timestamp:      getUnixTimestampAsString(),
		Server:         "Web",
		TransactionNum: "0004",
		Command:        "SELL",
	}

	logCommandID := uint8(commonlib.UserCommandType)

	return logCommandID, logParameters
}

// sendLog sends the given log as a message to the logging server
func sendLog(logCommandID uint8, logParameters commonlib.LogCommandParameter) {
	request := commonlib.GetSendableLogCommand(logCommandID, logParameters)
	if len(request) == 0 {
		fmt.Println("GetSendableLogCommand returned a 0-length byte slice")
	}

	replyBody, err := commonlib.SendCommand("POST", "application/json", state.loggingServerAddressAndPort, request)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf(replyBody)
}

// getUnixTimestampAsString returns a string representing the number of milliseconds (UTC) since the UNIX epoch
func getUnixTimestampAsString() string {
	timeNano := time.Now().UTC().UnixNano() // debugging
	fmt.Printf("timeNano:  %d\n", timeNano)

	ratio := int64(time.Nanosecond / time.Millisecond) // debugging
	fmt.Printf("ratio:  %d\n", ratio)

	convertedTime := timeNano * ratio // debugging
	fmt.Printf("convertedTime:  %d\n", convertedTime)

	timeString := strconv.FormatInt(convertedTime, 10) // debugging
	fmt.Printf("timeString:  %s\n", timeString)

	// return timeString
	return "1546300800780" // debugging
}