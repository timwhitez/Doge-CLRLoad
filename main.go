package main

import (
	"fmt"
	"github.com/timwhitez/Doge-CLRLoad/load"
	"io/ioutil"
	"log"
)

func main() {
	debug := true

	helloWorldPath := `testbin/hello.exe`
	helloWorld111Path := `testbin/hello111.exe`
	// Get hello
	helloBytes, err := ioutil.ReadFile(helloWorldPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("there was an error reading in the HelloWorld file from %s:\n%s", helloWorldPath, err))
	}
	if debug {
		fmt.Printf("[-] Ingested %d assembly bytes\n", len(helloBytes))
	}

	// Get hello111
	hello11Bytes, err := ioutil.ReadFile(helloWorld111Path)
	if err != nil {
		log.Fatal(fmt.Sprintf("there was an error reading in the hello111 file from %s:\n%s", helloWorld111Path, err))
	}

	if debug {
		fmt.Printf("[-] Ingested %d assembly bytes\n", len(hello11Bytes))
	}

	// Load assembly into default AppDomain
	stdout, e := load.LoadBin(helloBytes, []string{"arg111"}, "v4", debug)
	if e != nil {
		fmt.Printf("[DEBUG] Returned STDOUT/STDERR: \n%s\n", stdout)
	}

	// Execute assembly from default AppDomain x2
	if debug {
		fmt.Println("[-] Executing the helloworld x2...")
	}

	stdout, e = load.LoadBin(helloBytes, []string{""}, "v4", debug)
	if e != nil {
		fmt.Printf("[DEBUG] Returned STDOUT/STDERR: \n%s\n", stdout)
	}

	// Load assembly into default AppDomain
	stdout, e = load.LoadBin(hello11Bytes, []string{"11111"}, "v4", debug)
	if e != nil {
		fmt.Printf("[DEBUG] Returned STDOUT/STDERR: \n%s\n", stdout)
	}

	load.CleanCLR(debug)
}
