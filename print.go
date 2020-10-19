package main

import (
	"os"

	"github.com/fatih/color"
)

var errorPrintf = color.New(color.FgRed).PrintfFunc()
var successPrintf = color.New(color.FgGreen).PrintfFunc()

func PrintError(errMsg string, a ...interface{}) {
	errorPrintf(errMsg, a...)
}

func PrintFatalError(errMsg string, a ...interface{}) {
	PrintError(errMsg, a...)
	os.Exit(1)
}

func PrintSuccess(msg string, a ...interface{}) {
	successPrintf(msg, a...)
}
