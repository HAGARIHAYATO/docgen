package main

import (
	"errors"
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
	path := args[0] + "/"
	title := args[1]
	id := args[2]
	app, err := docgen.InitFireStore()
	if err != nil {
		panic(err)
	}
	data, err := docgen.GetDataByID(app, id)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(path + title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err != nil {
		panic(err)
	}
	_ = writeByres(file, data["text"])
	out, err := exec.Command("ls", "-la").Output()
	fmt.Println(string(out))
}

func writeByres(file *os.File, iter interface{}) error {
	switch iter := iter.(type) {
	case string:
		b := []byte(iter)
		_, err := file.Write(b)
		if err != nil {
			return err
		}
		return nil
	case []string:
		for _, line := range iter {
			b := []byte(line)
			_, err := file.Write(b)
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return errors.New("invalid type")
	}
}
