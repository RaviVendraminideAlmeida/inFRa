package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type DbParams struct {
	user     string
	password string
	port     int
}

// PORT | USER | DBName
func HandleDBParameters() DbParams {

	var err error

	params := &DbParams{
		user:     "root",
		password: "root",
		port:     5432,
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("*- Configuração BD -*\n\n")

	fmt.Print("Usuário do BD > ")
	params.user, err = reader.ReadString('\n')
	fmt.Println()

	if err != nil {
		log.Fatal(err)
	}

	params.user = strings.TrimSpace(params.user)

	fmt.Print("Senha do BD > ")
	params.password, err = reader.ReadString('\n')
	fmt.Println()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Porta do BD > ")
	params.password = strings.TrimSpace(params.password)

	_, err = fmt.Scanf("%d", &params.port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	return *params

}

func TouchFile(path, filename string) error {
	err := os.WriteFile(fmt.Sprintf("%s/%s", path, filename), []byte(""), 0755)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
