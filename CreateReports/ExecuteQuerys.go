package CreateReports

import (
	"Reportes_T/Config"
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadTxt() {
	readFile, err := os.Open("./consultasReportes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var linesTextFile []string
	for fileScanner.Scan() {
		linesTextFile = append(linesTextFile, fileScanner.Text())
	}
	readFile.Close()
	db, err := Config.CreateConnection()

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error de conexion al hacer Ping: \t %v", err)
		return
	}

	for key, value := range linesTextFile {
		fmt.Printf("Ejecutando query %k, %v", key, value)

	}

}
