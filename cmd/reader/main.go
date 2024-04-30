package main

import (
	"github.com/DiLRandI/my-finance/pdf-parser/pdf"
)

func main() {
	doc, err := pdf.Open("file.pdf")
	if err != nil {
		panic(err)
	}

	defer doc.Close()
}
