package header

import (
	"fmt"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func Build(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/nezuko.jpg", props.Rect{
					Center:  true,
					Percent: 50,
				})
				if err != nil {
					fmt.Println("Image file was not loaded:", err)
				}
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Prepared for you by the Div Rhino Fruit Company", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
