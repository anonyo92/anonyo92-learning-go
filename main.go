package main

import (
	"fmt"
	"os"
	"strconv"
	"learning-go/basics"
	"learning-go/ioPkg"
	"learning-go/concurrent"
	"learning-go/misc"
	"learning-go/webPkg"
)

func main() {
	// This is a single-line comment
	/* This is a block comment */
	
	clArgs := os.Args
	if len(clArgs) < 2 {
		fmt.Println("No option entered. Exiting program...")
		os.Exit(0)
	}

	execChapter := clArgs[1]

	switch execChapter {
	// Basics
	case "variables", "vars":
		basics.RunVariables()
	case "arrays": 
		basics.RunArrays()
	case "typeAssertions":
		basics.RunTypeAssertions()
	case "typeSwitch":
		basics.RunTypeSwitch(23)
		basics.RunTypeSwitch(65.3)
		basics.RunTypeSwitch(false)
		basics.RunTypeSwitch("false")
	case "accessScopes":
		basics.RunAccessScopes()
		fmt.Println(`A variable can also be exported for use in another package in the same way as functions, 
		such as variable 'P3' of the 'basics' package, whose value is`, basics.P3, "as accessed from the 'main' package.")
		fmt.Println(`To import packages from other modules however, the other module needs to be added as
		a dependency in this module - a require() clause in the go.mod file of this module.`)


	// IO
	case "fileRead", "readFile":
		ioPkg.RunFileRead()
	case "fileWrite", "writeFile":
		ioPkg.RunFileWrite()
	case "lineFilter":
		ioPkg.RunLineFilter()
	case "filePaths", "paths":
		ioPkg.RunFilePaths()
	case "directories", "dirs":
		ioPkg.RunDirectories()
	case "tmpFileDirs", "tmpFiles", "tmpDirectories", "tmpDirs":
		ioPkg.RunTmpFilesAndDirectories()


	// Data-Type specific Utils

	// Concurrent Programming
	case "goroutines":
		concurrent.RunGoroutines()
	case "channels":
		concurrent.RunChannels()
	case "channelBuffering", "chanBuffer":
		concurrent.RunChannelBuffering()
	case "channelDirections", "chanDir":
		concurrent.RunChannelDirections()
	case "channelSync", "chanSync":
		concurrent.RunChannelSync()
	case "channelSelect", "chanSelect":
		concurrent.RunChannelSelect()
	case "timeouts":
		concurrent.RunTimeouts()
	case "channelClose", "closeChannel":
		concurrent.RunCloseChannel()
	case "channelIter":
		concurrent.RunChannelIteration()
	case "timers":
		concurrent.RunTimers()
	case "tickers":
		concurrent.RunTickers()
	case "workerPool":
		concurrent.RunWorkerPool()
	case "waitGroup":
		concurrent.RunWaitGroup()
	case "rateLimit":
		concurrent.RunRateLimiting()
	case "atomicCounter":
		concurrent.RunAtomicCounter()
	case "mutex":
		concurrent.RunMutex()
	case "statefulGoroutine":
		concurrent.RunStatefulGoroutine()

	
	// Miscellaneous
	case "defer":
		misc.RunDefer()
	case "panic":
		misc.RunPanic()
	case "recover":
		misc.RunRecover()
	case "exit":
		misc.RunExit()
	case "envVar":
		misc.RunEnvVars()
	

	// Web and Database
	case "httpClient":
		var method, url, payload string = "", "", ""
		if len(clArgs) < 3 || (len(clArgs) == 3 && clArgs[2] == "GET") {
			method = "GET"
		} else if len(clArgs) == 4 && (clArgs[2] == "GET" || clArgs[2] == "DELETE") {
			method = clArgs[2]
			url = clArgs[3]
		} else if len(clArgs) > 4 {
			method = clArgs[2]
			url = clArgs[3]
			payload = clArgs[4]
		} else {
			fmt.Println("Invalid or Missing arguments.")
			os.Exit(3)
		}
		webPkg.RunHttpClient(method, url, payload)

	case "httpServer":
		serverPort := 8090
		var err error
		if len(clArgs) >= 3 {
			serverPort, err = strconv.Atoi(clArgs[2])
			if err != nil {
				panic(err)
			}
		}
		webPkg.RunHttpServer(serverPort)
	case "context", "ctx":
		serverPort := 8090
		var err error
		if len(clArgs) >= 3 {
			serverPort, err = strconv.Atoi(clArgs[2])
			if err != nil {
				panic(err)
			}
		}
		webPkg.RunContexts(serverPort)

	
	default:
		fmt.Println("Invalid option. Try again!")
	}
}
