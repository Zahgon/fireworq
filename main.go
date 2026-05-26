//go:generate go-assets-builder -p main -o assets.go LICENSE AUTHORS CREDITS
package main

import (
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/fireworq/fireworq/config"
	"github.com/fireworq/fireworq/dispatcher"
	logwriter "github.com/fireworq/fireworq/log"
	"github.com/fireworq/fireworq/web"
)

func main() {
	out := os.Stderr

	initDefaultConfig()

	args, err := parseCmdArgs(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	if args.showVersion {
		fmt.Fprintln(out, versionString(" "))
		os.Exit(0)
	}

	if args.showLicense {
		fmt.Println(licenseText)
		os.Exit(0)
	}

	if args.showCredits {
		fmt.Println(creditsText)
		os.Exit(0)
	}

	for _, k := range config.Keys() {
		config.Set(k, *args.settings[k])
	}

	accessLog := initLogging(syscall.SIGUSR1)
	initProcess()
	dispatcher.Init()
	web.Init()

	startServer(accessLog)
}

type cmdArgs struct {
	showVersion bool
	showLicense bool
	showCredits bool
	settings    map[string]*string
}

func parseCmdArgs(args []string) (*cmdArgs, error) { _ = "STUB: not implemented"; return nil, nil }

func initDefaultConfig() { _ = "STUB: not implemented"; return }

func initProcess() { _ = "STUB: not implemented"; return }

func initLogging(sig syscall.Signal) (accessLog logwriter.Writer) {
	_ = "STUB: not implemented"
	// Access log
	return *new(logwriter.Writer)
}

// Error log

// Queue log

// Reopening log files (for logrotate)

func startServer(accessLogWriter io.Writer) { _ = "STUB: not implemented"; return }

func versionString(sep string) string { _ = "STUB: not implemented"; return "" }

func mustAssetString(path string) string { _ = "STUB: not implemented"; return "" }

var (
	licenseText = mustAssetString("/LICENSE") + `
   Copyright (c) 2017 The Fireworq Authors. All rights reserved.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

` + mustAssetString("/AUTHORS")
	creditsText = mustAssetString("/CREDITS")

	helpText = `Usage: fireworq [options]

  A lightweight, high performance job queue system.

Options:

  --version, -v  Show the version string.
  --license      Show the license text.
  --credits      Show the library dependencies and their licenses.
  --help, -h     Show the help message.
`
)
