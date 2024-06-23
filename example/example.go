package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dosadczuk/go-pdftohtml"
)

func main() {
	cmd, err := pdftohtml.NewCommand(
		pdftohtml.WithOutdirOverwrite(),
		pdftohtml.WithEmbedMetaTags(),
		pdftohtml.WithEmbedFormFields(),
		pdftohtml.WithEmbedFonts(),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Run(context.Background(), "./example.pdf", "./html")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Done")
}
