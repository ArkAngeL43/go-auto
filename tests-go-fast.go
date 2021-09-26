package main 


/// testing module 

import (
    "fmt"
    "log"
    "net/http"
	"os/exec"
)


func exit() {
	out, err := exec.Command("exit").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println("[-] Exiting..", output)
}

func Command() {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println("[+] Current Dir", output)
}

func clear() {
	out, err := exec.Command("clear").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println("[+] Clearning...", output)
}

func secondtest() {
	resp, err := http.Get("http://golangcode.com")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Response -> ", resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Println("Returned Error => ", resp.StatusCode)
	}
    if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
        fmt.Println("HTTP Returned in range of 200 ")
    } else {
        fmt.Println("Broken...")
		exit()
    }
}

func main() {
	clear()
	Command()
	secondtest()
}