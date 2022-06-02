package main

import (
	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func main() {
	var err error
	pdf := gofpdf.New("P", "mm", "A4", "")
	filePaths := []string{
		"example-pdf-1.pdf",
		"example-pdf-2.pdf",
	}
	for _, path := range filePaths {
		addPdf(pdf, path)
	}
	err = pdf.OutputFileAndClose("example.pdf")
	if err != nil {
		panic(err)
	}
}

func addPdf(pdf *gofpdf.Fpdf, path string) {
	template := gofpdi.ImportPage(pdf, path, 1, "/MediaBox")
	pdf.AddPage()
	gofpdi.UseImportedTemplate(pdf, template, 0, 0, 210, 0)
}
