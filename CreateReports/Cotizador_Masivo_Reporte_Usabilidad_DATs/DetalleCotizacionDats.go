package main

import (
	"Reportes_T/CreateReports"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"time"
)

func CreateReport() {

	var nameReporte = "DetalleCotizacionDats.xlsx"
	var _, err = os.Stat(nameReporte)
	if os.IsNotExist(err) {
		var fechaUt = getDate()
		var file = excelize.NewFile()
		index := file.NewSheet(fechaUt)
		file.DeleteSheet("Sheet1")

		style, err := file.NewStyle(`{"font":{"bold": true, "italic":false, "size": 10, "color":"#FFFFFF", "family": "Arial"}, "fill":{"type":"pattern","color":["#010080"],"pattern":1}}`)

		if err != nil {
			fmt.Println(err)
		}

		for key, value := range titlesDetalleCotizacionDats {
			file.SetColWidth(fechaUt, key, key, float64(value))
		}
		for key, value := range titlesDetalleCotizacionDatsValuesTitles {
			file.SetCellValue(fechaUt, key+"1", value)
		}
		err = file.SetCellStyle(fechaUt, "A1", "Z1", style)
		file.SetActiveSheet(index)

		if err := file.SaveAs(nameReporte); err != nil {
			fmt.Println(err)
		}
	}
}
func main() {
	CreateReports.ReadTxt()
	//CreateReport()
}

func getDate() (dateStr string) {
	fecha := time.Now()
	ayer := fecha.AddDate(0, 0, -1)
	//.Format("2006-02-01")
	return ayer.Format("2006-01-02")
}

var titlesDetalleCotizacionDats = map[string]int{"A": 15, "B": 12, "C": 12, "D": 14, "E": 8, "F": 12, "G": 18, "H": 22, "I": 22, "J": 26, "K": 24,
	"L": 20, "M": 28, "N": 27, "O": 8, "P": 14, "Q": 15, "R": 20, "S": 20, "T": 20, "U": 20, "V": 21, "W": 9, "X": 5, "Y": 9, "Z": 9}

var titlesDetalleCotizacionDatsValuesTitles = map[string]string{
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
