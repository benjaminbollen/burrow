// Copyright 2015, 2016 Eris Industries (UK) Ltd.
// This file is part of Eris-RT

// Eris-RT is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Eris-RT is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Eris-RT.  If not, see <http://www.gnu.org/licenses/>.

package commands

import (
	"fmt"

	cobra "github.com/spf13/cobra"

	log	"github.com/eris-ltd/eris-logger"

	definitions "github.com/eris-ltd/eris-db/definitions"
	version	    "github.com/eris-ltd/eris-db/version"
)

// Global Do struct
var doTest *definitions.Do

var ErisTestCmd = &cobra.Command {
	Use:   "eris-db-test",
	Short: "Eris-DB-test runs provided tests against an existing chain and reports on failures.",
	Long:  `Eris-DB-test runs provided tests against an existing chain and reports on failures.

Made with <3 by Eris Industries.

Complete documentation is available at https://docs.erisindustries.com
` + "\nVERSION:\n " + version.VERSION,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

	log.SetLevel(log.WarnLevel)
	if doTest.Verbose {
	  log.SetLevel(log.InfoLevel)
	} else if doTest.Debug {
	  log.SetLevel(log.DebugLevel)
	}

	// if WorkDir was not set by a flag or by $ERIS_DB_WORKDIR
	// NOTE [ben]: we can consider an `Explicit` flag that eliminates
	// the use of any assumptions while starting Eris-DB
	if doTest.WorkDir == "" {
	  if currentDirectory, err := os.Getwd(); err != nil {
		log.Fatalf("No directory provided and failed to get current working directory: %v", err)
		os.Exit(1)
	  } else {

		doTest.WorkDir = currentDirectory
	  }
	}
	if !util.IsDir(do.WorkDir) {
	  log.Fatalf("Provided working directory %s is not a directory", do.WorkDir)
	}

  },
  Run: RunTestSuite,
}

func ExecuteTest() {
	InitErisDbTest()
	fmt.Printf("TESTING")
}

func InitErisDbTest() {
  // initialise an empty do struct for command execution
	doTest = definitions.NowDo()
}

func AddGlobalFlags() {
	ErisDbCmd.PersistentFlags().BoolVarP(&doTest.Verbose, "verbose", "v", defaultVerbose(), "verbose output; more output than no output flags; less output than debug level; default respects $ERIS_DB_VERBOSE")
	ErisDbCmd.PersistentFlags().BoolVarP(&doTest.Debug, "debug", "d", defaultDebug(), "debug level output; the most output available for eris-db; if it is too chatty use verbose flag; default respects $ERIS_DB_DEBUG")
}

//------------------------------------------------------------------------------
// Test Suite

fun RunTestSuite(cmd *cobra.Command, args []string) {
	
}



//------------------------------------------------------------------------------
// Defaults

// defaultVerbose is set to false unless the ERIS_DB_VERBOSE environment
// variable is set to a parsable boolean.
func defaultVerbose() bool {
  return setDefaultBool("ERIS_DB_VERBOSE", false)
}

// defaultDebug is set to false unless the ERIS_DB_DEBUG environment
// variable is set to a parsable boolean.
func defaultDebug() bool {
  return setDefaultBool("ERIS_DB_DEBUG", false)
}

// setDefaultBool returns the provided default value if the environment variab;e
// is not set or not parsable as a bool.
func setDefaultBool(environmentVariable string, defaultValue bool) bool {
	value := os.Getenv(environmentVariable)
	if value != "" {
		if parsedValue, err := strconv.ParseBool(value); err == nil {
		return parsedValue
	}
	}
	return defaultValue
}

func setDefaultString(envVar, def string) string {
	env := os.Getenv(envVar)
	if env != "" {
		return env
	}
	return def
}

func setDefaultStringSlice(envVar string, def []string) []string {
	env := os.Getenv(envVar)
	if env != "" {
		return strings.Split(env, ",")
	}
	return def
}
