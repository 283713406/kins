/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"fmt"
	"os"
	"strings"

)

const (
	// DefaultErrorExitCode defines exit the code for failed action generally
	DefaultErrorExitCode = 1
	// PreFlightExitCode defines exit the code for preflight checks
	PreFlightExitCode = 2
	// ValidationExitCode defines the exit code validation checks
	ValidationExitCode = 3
)

// fatal prints the message if set and then exits.
func fatal(msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}

// CheckErr prints a user friendly error to STDERR and exits with a non-zero
// exit code. Unrecognized errors will be printed with an "error: " prefix.
//
// This method is generic to the command in use and may be used by non-Kubectl
// commands.
func CheckErr(err error) {
	checkErr(err, fatal)
}

// preflightError allows us to know if the error is a preflight error or not
// defining the interface here avoids an import cycle of pulling in preflight into the util package
type preflightError interface {
	Preflight() bool
}

// checkErr formats a given error as a string and calls the passed handleErr
// func with that string and an exit code.
func checkErr(err error, handleErr func(string, int)) {
	switch err.(type) {
	case nil:
		return
	case preflightError:
		handleErr(err.Error(), PreFlightExitCode)

	default:
		handleErr(err.Error(), DefaultErrorExitCode)
	}
}

// FormatErrMsg returns a human-readable string describing the slice of errors passed to the function
func FormatErrMsg(errs []error) string {
	var errMsg string
	for _, err := range errs {
		errMsg = fmt.Sprintf("%s\t- %s\n", errMsg, err.Error())
	}
	return errMsg
}
