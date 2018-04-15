package main

import (
	"github.com/jung-kurt/gofpdf"
	"os"
	"time"
)


func main() {
	data:=loadCSV(path())
	pdf := newReport()

	pdf = header(pdf, data[0])
	pdf = table(pdf, data[1:])

	pdf = image(pdf)



}

func loadCSV(path string) [][]string {

}

func path() string {
	if len(os.Args) < 2 {
		return "ordersReport.csv"
	}
	return os.Args[1]
}

func newReport() *gofpdf.Fpdf {
	//create pdf with New(), L = landscape P = portrait
	// paper format is Letter
	// last argument is path directory to font
	pdf := gofpdf.New("L", "mm", "Letter", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 28)
	pdf.Cell(40, 10, "Daily Report")

	pdf.Ln(12)
	pdf.SetFont("Times", "", 20)
	pdf.Cell(40, 10, time.Now().Format("Sun April 15, 2018"))
	pdf.Ln(20)

	return pdf
}




