// Copyright 2023 Northern.tech AS
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//	    http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.
package log

import (
	"fmt"
	"os"
)

var (
	verbose = false
)

func Setup(verb bool) {
	verbose = verb
}

func IsVerbose() bool {
	return verbose
}

func Err(msg string) {
	fmt.Fprintln(os.Stderr, msg)
}

func Errf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg, args...)
}

func Verb(msg string) {
	if verbose {
		fmt.Fprintln(os.Stdout, "VERBOSE "+msg)
	}
}

func Verbf(msg string, args ...interface{}) {
	if verbose {
		fmt.Fprintf(os.Stdout, "VERBOSE "+msg, args...)
	}
}

func Info(msg string) {
	fmt.Fprintln(os.Stdout, msg)
}

func Infof(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg, args...)
}
