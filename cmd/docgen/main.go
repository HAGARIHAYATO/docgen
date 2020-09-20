package main

import (
	"errors"
	"fmt"
	"github.com/HAGARIHAYATO/docgen"
	"log"
	"os"
)
func main() {

	args := os.Args[1:]

	path, title, id, err := argsHandle(args)
	if err != nil {
		log.Fatal(err)
	}

	text, err := docgen.GetDataByID(id)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(path + title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	_ = writeByres(file, text)
}

func argsHandle(args []string) (string, string, string, error) {
	var err string
	switch len(args) {
	case 3:
		path := args[0] + "/"
		title := args[1]
		id := args[2]
		return path, title, id, nil
	default:
		err = fmt.Sprintf(
			"invalid arguments counts. expected to get 3 arguments, but got %v. \n ex... docgen <path> <file-title> <docID>",
			len(args),
			)
		return "", "", "", errors.New(err)
	}
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
