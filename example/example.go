package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dosadczuk/pdftohtml"
)

func main() {
	cmd := pdftohtml.NewCommand(
		pdftohtml.WithOutdirOverwrite(),
		pdftohtml.WithEmbedMetaTags(),
		pdftohtml.WithEmbedFormFields(),
		pdftohtml.WithEmbedFonts(),
	)

	err := cmd.Run(context.Background(), "./example.pdf", "./html")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Done")
}
