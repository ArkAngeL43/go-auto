package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "os/exec"
)


func clear() {
	out, err := exec.Command("clear").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:]) 
	fmt.Printf("", output)
}


func sleep() {
	out, err := exec.Command("sleep 6").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println("[-] Exiting..", output)
}

func exit() {
	out, err := exec.Command("exit").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println("[-] Exiting..", output)
}

func main() {
  clear()
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("---------------------")

  for {
    fmt.Print("Password >>> ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)

    if strings.Compare("PSQu013dn", text) == 0 {
      runit()
    } else {
		fmt.Println("[-] Incorrect Password ")
		fmt.Println("[!] Try Again [!] ")
		sleep()
		clear()
	}

  }

}





//func one() {
//	var age int//
//	fmt.Println(" Options ")
	//_, err := fmt.Scan(&age)
//
//	if age == 0{
//		fmt.Println("hello")
//	}
//}

//func main() {
//	one()
//}

