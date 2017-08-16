package main

import (
	"runtime"
	"log"
	"flag"
	"os"
	"fmt"
	ss "github.com/vaetern/moxyServer/servingStrategies"
	"errors"
)

const MaxProcesses = 2

const ServingPort = "19501"

func main() {

	operationMode, operationPort, verboseMode := handleStartupSetting()

	runtime.GOMAXPROCS(MaxProcesses)

	servingStrategy, err := getServingStrategy(operationMode)

	if err != nil{
		log.Fatal(err)
		os.Exit(1)
	}

	servingStrategy.Start(operationPort, verboseMode)
}

func handleStartupSetting() (*string, *string, *bool) {

	fs := flag.NewFlagSet("moxy server", flag.ExitOnError)
	operationMode := fs.String("mode", "", "Operation mode (store/cache)")
	operationPort := fs.String("port", ServingPort, "Serving port (default :19501)")
	verboseMode := fs.Bool("verbose", true, "Verbose log")

	fs.Parse(os.Args[1:])

	fs.Usage = func() {
		fmt.Println("Usage: moxyServer.exe -mode=store -port=19501 -verbose=1")
	}

	if *operationMode == "" {
		fs.Usage()
		os.Exit(1)
	}

	return operationMode, operationPort, verboseMode
}

func getServingStrategy(operationMode *string) (strategy ss.ServingStrategy, err error){

	if *operationMode == "store"{
		strategy = ss.NewServeAndStoreStrategy()
	} else if *operationMode == "cache"{
		strategy = ss.NewServeFromCacheStrategy()
	} else{
		err = errors.New("Unexpected operation mode")
	}


	return strategy, err
}
