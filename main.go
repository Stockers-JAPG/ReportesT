/*package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"

	_ "github.com/godror/godror"
	"log"
	"os"
	"time"
)

func main() {
	readFile, err := os.Open("./consultasReportes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	readFile.Close()

	db, err := createConnection()
	fmt.Println('\n')
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}
	fmt.Println("\n")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("\n")
		fmt.Printf("Error conectando: %v", err)
		return
	}

	fmt.Println("\n Conectado Correctamente")
	fmt.Println("\n Empezando con las consultas")

	crearArchivosReportes("JOSE", "dato1")

}

func createConnection() (db *sql.DB, e error) {
	db, err := sql.Open("godror", "CUDGADM/Telcel01#@10.191.209.174:1522/CUDG")
	if err != nil {
		fmt.Println(err)
		return
	}
	return db, nil
}

func crearArchivosReportes(nombre string, datos string) (archivoSalida string) {
	fecha := time.Now().Format("2006-02-01")

	var nameFull = nombre + "_" + fecha + ".xlsx"
	fmt.Println(nameFull)
	var _, err = os.Stat(nameFull)
	if os.IsNotExist(err) {

		titlesDetalleCotizacionDats := map[string]int{"A": 15, "B": 12, "C": 12, "D": 14, "E": 8, "F": 12, "G": 18, "H": 22, "I": 22, "J": 26, "K": 24,
			"L": 20, "M": 28, "N": 27, "O": 8, "P": 14, "Q": 15, "R": 20, "S": 20, "T": 20, "U": 20, "V": 21, "W": 9, "X": 5, "Y": 9, "Z": 9}

		titlesDetalleCotizacionDatsValuesTitles := map[string]string{
			"A": "ID_COTIZACION",
			"B": "ID_USUARIO",
			"C": "IP",
			"D": "FORMA_PAGO",
			"E": "PRECIO",
			"F": "PRECIO_IVA",
			"G": "PRECIO_SUBSIDIO",
			"H": "PRECIO_SUBSIDIO_IVA",
			"I": "PRECIO_SIN_SUBSIDIO",
			"J": "PRECIO_SIN_SUBSIDIO_IVA",
			"K": "FECHA_HORA_REGISTRO",
			"L": "CORREO_ENVIO",
			"M": "ID_OFERTAPRIMARIA",
			"N": "ID_OFERTASUPLEMENTARIA",
			"O": "REGION",
			"P": "IVA",
			"Q": "SOBREPRECIO",
			"R": "CARGO_EQUIPO",
			"S": "CARGO_EQUIPO_IVA",
			"T": "IDENTIFICADOR_UNO",
			"U": "IDENTIFICADOR_DOS",
			"V": "IDENTIFICADOR_TRES",
			"W": "PLAZO",
			"X": "SKU",
			"Y": "COLOR",
			"Z": "ENVIADO"}

		var file = excelize.NewFile()

		index := file.NewSheet(fecha)
		file.DeleteSheet("Sheet1")

		file.SetCellValue(fecha, "A1", "JOSE ANTONIO")

		style, err := file.NewStyle(`{"font":{"bold": true, "italic":false, "size": 10, "color":"#FFFFFF", "family": "Arial"}, "fill":{"type":"pattern","color":["#000DA1"],"pattern":1}}`)

		if err != nil {
			fmt.Println(err)
		}

		for key, value := range titlesDetalleCotizacionDats {
			file.SetColWidth(fecha, key, key, float64(value))
		}

		for key, value := range titlesDetalleCotizacionDatsValuesTitles {
			file.SetCellValue(fecha, key+"1", value)
		}
		err = file.SetCellStyle(fecha, "A1", "Z1", style)
		file.SetActiveSheet(index)

		if err := file.SaveAs(nameFull); err != nil {
			fmt.Println(err)
		}

	}

	return nameFull
}
*/