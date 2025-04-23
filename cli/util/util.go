package util

import (
	"fmt"
	"log"
	"os"
)

func TouchFile(path, filename string) error {
	err := os.WriteFile(fmt.Sprintf("%s/%s", path, filename), []byte(""), 0755)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
