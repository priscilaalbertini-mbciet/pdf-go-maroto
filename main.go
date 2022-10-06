package main

import (
	"fmt"
	"pdf-go-maroto/body"
	"pdf-go-maroto/header"

	"os"

	"github.com/google/uuid"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	header.Build(m)
	body.Build(m)

	err := m.OutputFileAndClose("pdfs/" + uuid.New().String() + ".pdf")
	if err != nil {
		fmt.Println("Could not sabe PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}
