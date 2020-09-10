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
	path := args[1] + "/"
	file, err := os.OpenFile(path + "Dockerfile", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	comFile, err := os.OpenFile(path + "docker-compose.yml", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	defer comFile.Close()

	fw, err := argsAnalyzer(args[0])
	if err != nil {
		panic(err)
	}
	_ = writeByres(file, fw["Doc"])
	_ = writeByres(comFile, fw["ComposeDoc"])
	out, err := exec.Command("ls", "-la").Output()
	fmt.Println(string(out))
}

func argsAnalyzer(arg string) (map[string][]string, error) {
	switch arg {
	case "rails":
		return docgen.Rails, nil
	default:
		return nil, errors.New("invalid frame work name")
	}
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
