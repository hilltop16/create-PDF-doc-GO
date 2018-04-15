package main

import (
	"github.com/jung-kurt/gofpdf"
	"os"
	"time"
	"log"
)


func main() {
	data:=loadCSV(path())
	pdf := newReport()

	pdf = header(pdf, data[0])
	pdf = table(pdf, data[1:])

	pdf = image(pdf)

	if pdf.Err() {
		log.Fatalf("Creating pdf file has failed: %s\n", pdf.Error())
	}
	err:= savePDF(pdf)
	if err != nil {
		log.Fatalf("cannot save PDF: %s|n", err)
	}

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

func image(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	pdf.ImageOptions("stats.png", 255, 10, 25, 25,
		false, gofpdf.ImageOptions{ImageType:"PNG", ReadDpi: true}, 0, "")
	return pdf
}

func savePDF(pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose("report.pdf")
}


