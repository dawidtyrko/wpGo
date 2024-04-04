package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Name     string
	Age      int
	Password string
}

func main() {
	scan := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your name: ")
	scan.Scan()
	name := scan.Text()

	fmt.Print("Enter your age: ")
	scan.Scan()
	ageString := scan.Text()
	ageNum, err := strconv.Atoi(ageString)
	if err != nil {
		fmt.Println("Invalid input: ", err)
	}

	fmt.Print("Enter new password: ")
	scan.Scan()
	firstPass := scan.Text()
	firstPass = encryption(firstPass)

	fmt.Print("Confirm your password: ")
	scan.Scan()
	secondPass := scan.Text()
	secondPass = encryption(secondPass)

	fL := strings.ToUpper(string(name[0]))
	lL := name[1:]
	name = fL + lL

	if firstPass == secondPass {
		data := Data{Name: name, Age: ageNum, Password: secondPass}
		fmt.Println("Great success!\n-----------")
		fmt.Println("Name: ", data.Name)
		fmt.Println("Age: ", data.Age)
		fmt.Println("Encrypted password: ", data.Password)
	} else {
		fmt.Println("wrong password")
		return
	}
}

func encryption(password string) string {
	p := sha256.New()
	p.Write([]byte(password))
	//fmt.Printf("%x", p.Sum(nil))
	return hex.EncodeToString(p.Sum(nil))
}
