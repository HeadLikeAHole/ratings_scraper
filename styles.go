package main

import "github.com/xuri/excelize/v2"

var (
	font12 = &excelize.Font{
		Size: 12,
	}

	fontBold12 = &excelize.Font{
		Bold: true,
		Size: 12,
	}

	fontBold14 = &excelize.Font{
		Bold: true,
		Size: 14,
	}

	alignCenter = &excelize.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	}

	border = []excelize.Border{
		{Type: "left", Color: "#000000", Style: 1},
		{Type: "right", Color: "#000000", Style: 1},
		{Type: "top", Color: "#000000", Style: 1},
		{Type: "bottom", Color: "#000000", Style: 1},
	}
)
