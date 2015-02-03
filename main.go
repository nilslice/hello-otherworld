package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

const code = `
package main

import (
    "fmt"
)

type Name struct {
    first string
    last  string
    full  string
}

func (n *Name) FullName() string {
   n.full = n.first + " " + n.last
   return n.full
}

func Greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    person := &Name{
        first: "Steve",
        last: "Manuel",
    }
    Greet(person.FullName())
}
`

func main() {
	program, err := template.New("code").Parse(code)
	if err != nil {
		fmt.Println("44 Error:", err)
		os.Exit(1)
	}

	err = os.Mkdir("hello", os.ModePerm)
	if err != nil {
		fmt.Println("50 Error:", err)
		os.Exit(1)
	}
	err = os.Chdir("hello")
	if err != nil {
		fmt.Println("55 Error:", err)
		os.Exit(1)
	}

	file, err := os.Create("hello.go")
	if err != nil {
		fmt.Println("61 Error:", err)
		os.Exit(1)
	}
	err = program.Execute(file, nil)
	if err != nil {
		fmt.Println("66 Error:", err)
		os.Exit(1)
	}
	file.Close()

	build := exec.Command("go", "build")
	err = build.Run()
	if err != nil {
		fmt.Println("74 Error:", err)
		os.Exit(1)
	}

	hello := exec.Command("./hello")
	hello.Stdout = os.Stdout
	hello.Stderr = os.Stderr
	err = hello.Run()
	if err != nil {
		fmt.Println("83 Error:", err)
		os.Exit(1)
	}

	clean := exec.Command("go", "clean")
	err = clean.Run()
	if err != nil {
		fmt.Println("90 Error:", err)
		os.Exit(1)
	}

	err = os.RemoveAll("../hello")
	if err != nil {
		fmt.Println("96 Error:", err)
		os.Exit(1)
	}

}
