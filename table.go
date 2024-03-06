package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func createTable() error {
	err := styleTable()
	if err != nil {
		return err
	}

	err = setColsWidth()
	if err != nil {
		return err
	}

	err = setRowsHeight()
	if err != nil {
		return err
	}

	err = createHeaders()
	if err != nil {
		return err
	}

	return nil
}

func styleTable() error {
	mainHeaderStyle, err := xl.NewStyle(&excelize.Style{
		Font:      fontBold14,
		Alignment: alignCenter,
	})
	if err != nil {
		return err
	}

	err = xl.SetCellStyle("Sheet1", "A1", "A1", mainHeaderStyle)
	if err != nil {
		return err
	}

	tableHeadersStyle, err := xl.NewStyle(&excelize.Style{
		Border:    border,
		Font:      fontBold12,
		Alignment: alignCenter,
	})
	if err != nil {
		return err
	}

	err = xl.SetCellStyle("Sheet1", "A2", "E2", tableHeadersStyle)
	if err != nil {
		return err
	}

	tableRowsStyle, err := xl.NewStyle(&excelize.Style{
		Border:    border,
		Font:      font12,
		Alignment: alignCenter,
	})
	if err != nil {
		return err
	}

	err = xl.SetCellStyle("Sheet1", "A3", "E1014", tableRowsStyle)
	if err != nil {
		return err
	}

	return nil
}

func setColsWidth() error {
	err := xl.SetColWidth("Sheet1", "A", "A", 7)
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetColWidth("Sheet1", "B", "B", 50)
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetColWidth("Sheet1", "C", "C", 9)
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetColWidth("Sheet1", "D", "D", 15)
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetColWidth("Sheet1", "E", "E", 9)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func setRowsHeight() error {
	err := xl.SetRowHeight("Sheet1", 2, 17)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func createHeaders() error {
	err := xl.MergeCell("Sheet1", "A1", "E1")
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetCellValue("Sheet1", "A1", "Movie Ratings List")
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetCellValue("Sheet1", "A2", "№")
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetCellValue("Sheet1", "B2", "Title")
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetCellValue("Sheet1", "C2", "Year")
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetCellValue("Sheet1", "D2", "Date")
	if err != nil {
		fmt.Println(err)
	}

	err = xl.SetCellValue("Sheet1", "E2", "Rating")
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func setMovieRows(movies []*movie) error {
	for index, item := range movies {
		num := index + 1
		offset := strconv.Itoa(num + 2)

		err := xl.SetCellValue("Sheet1", "A"+offset, num)
		if err != nil {
			return err
		}
		err = xl.SetCellValue("Sheet1", "B"+offset, item.title)
		if err != nil {
			return err
		}
		err = xl.SetCellValue("Sheet1", "C"+offset, item.year)
		if err != nil {
			return err
		}
		err = xl.SetCellValue("Sheet1", "D"+offset, item.date)
		if err != nil {
			return err
		}
		err = xl.SetCellValue("Sheet1", "E"+offset, item.rating)
		if err != nil {
			return err
		}
	}

	return nil
}
