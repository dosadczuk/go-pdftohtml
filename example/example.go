package main

import (
	"fmt"
	"log"

	"github.com/dosadczuk/pdftohtml"
)

func main() {
	cmd := pdftohtml.NewCommand(
		pdftohtml.WithOutdirOverwrite(),
		pdftohtml.WithMeta(),
		pdftohtml.WithFormFields(),
		pdftohtml.WithEmbededFonts(),
	)

	err := cmd.Run("./example.pdf", "./html")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Done")
}
