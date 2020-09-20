package main

import (
	"fmt"
	"github.com/HAGARIHAYATO/docgen"
	"os"
	"os/exec"
)
func main() {
	var args []string
	if len(os.Args) > 2 {
		args = os.Args[1:]
	} else if len(os.Args) > 1  {
		fmt.Println("expected 2 arguments but give 1 argument")
	} else {
		fmt.Println("invalid input")
	}
	title := args[1]
	path := args[2] + "./"
	id := args[3]
	app, err := docgen.InitFireStore()
	if err != nil {
		panic(err)
	}
	docgen.GetDataByID(app, id)
	file, err := os.OpenFile(path + title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err != nil {
		panic(err)
	}
	_ = writeByres(file, []string{"ss"})
	out, err := exec.Command("ls", "-la").Output()
	fmt.Println(string(out))
}

func writeByres(file *os.File, array []string) error {
	for _, line := range array {
		b := []byte(line)
		_, err := file.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
