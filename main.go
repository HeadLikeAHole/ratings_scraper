package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var xl *excelize.File

type movie struct {
	num    int
	title  string
	year   int
	date   string
	rating int
}

func main() {
	var movies []*movie

	workingDir, _ := os.Getwd()
	filesDir := filepath.Join(workingDir, "ratings")
	baseURL := "file:///"

	t := &http.Transport{}
	t.RegisterProtocol("file", http.NewFileTransport(http.Dir(filesDir)))

	c := colly.NewCollector(
		// MaxDepth is 2, so only the links on the scraped page
		// are visited, and no further links are followed
		colly.MaxDepth(2),
	)
	c.WithTransport(t)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if strings.HasSuffix(link, ".html") {
			err := e.Request.Visit(baseURL + link)
			if err != nil && err != colly.ErrMaxDepth {
				fmt.Println(err)
			}
		}
	})

	c.OnHTML(".profileFilmsList .item", func(e *colly.HTMLElement) {
		numStr := e.ChildText(".num")
		title := e.ChildText(".nameEng")
		titleRus := e.ChildText(".nameRus")
		yearStr := extractYear(titleRus)
		date := extractDate(e.ChildText(".date"))
		ratingStr := e.ChildText(".myVote")

		if title == "" {
			title = titleRus
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println(err)
		}
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			fmt.Println(err)
		}
		rating, err := strconv.Atoi(ratingStr)
		if err != nil {
			fmt.Println(err)
		}

		item := &movie{
			num:    num,
			title:  title,
			year:   year,
			date:   date,
			rating: rating,
		}

		movies = append(movies, item)
	})

	err := c.Visit(baseURL)
	if err != nil {
		fmt.Println(err)
	}

	sort.Sort(byNumDesc(movies))

	xl = excelize.NewFile()
	defer func() {
		if err := xl.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	err = createTable()
	if err != nil {
		fmt.Println(err)
	}

	err = setMovieRows(movies)
	if err != nil {
		fmt.Println(err)
	}

	if err = xl.SaveAs("movie_list.xlsx"); err != nil {
		fmt.Println(err)
	}
}
