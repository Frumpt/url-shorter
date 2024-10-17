package main

import (
	"fmt"
	"log"
	"os"
	"url-sorter/internal/config"
)

func main() {
	if err := os.Setenv("CONFIG_PATH", ".\\config\\local.yaml"); err != nil {
		log.Fatalf("can not set env %s", err)
	}
	cnf := config.MustLoad()
	fmt.Println(cnf)
}
