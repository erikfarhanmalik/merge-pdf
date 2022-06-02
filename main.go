package main

import (
	"log"

	"github.com/dslipak/pdf"
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func main() {
	var err error
	pdfResult := gofpdf.New("P", "mm", "A4", "")
	filePaths := []string{
		"example-pdf-1.pdf",
		"example-pdf-2.pdf",
	}
	for _, path := range filePaths {
		f, err := pdf.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		addPdf(pdfResult, path, f.NumPage())
	}
	err = pdfResult.OutputFileAndClose("example.pdf")
	if err != nil {
		panic(err)
	}
}

func addPdf(pdf *gofpdf.Fpdf, path string, pageNumber int) {
	for i := 1; i <= pageNumber; i++ {
		template := gofpdi.ImportPage(pdf, path, i, "/MediaBox")
		pdf.AddPage()
		gofpdi.UseImportedTemplate(pdf, template, 0, 0, 210, 0)
	}
}
