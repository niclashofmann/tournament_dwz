package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func parsePointsAndRatings(txt string) (p string, rs []string) {
	r := csv.NewReader(strings.NewReader(txt))
	c := colly.NewCollector()
	ps := float64(0)

	c.OnHTML("body > div.content > div > table > tbody > tr:nth-child(3) > td:nth-child(2)", func(h *colly.HTMLElement) {
		rs = append(rs, h.Text)
	})

	for {
		o, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		ps += mustParseFloat64(o[1])
		c.Visit("http://www.schachbezirk4.de/dwz/player.php?pkz=" + o[0])
	}

	p = fmt.Sprintf("%.1f", ps)

	return
}

func requestCalculation(d, by, p string, rs []string) (res string) {
	c := colly.NewCollector()
	c.OnHTML("body > form > fieldset:nth-child(4) > dl > dd:nth-child(8)", func(h *colly.HTMLElement) {
		res = h.Text
	})

	form := url.Values{}
	form.Add("dwz_name", d)
	form.Add("geb", by)
	form.Add("punkte", p)
	form.Add("Los", "DWZ berechnen!")
	form.Add("name", "")
	form.Add("kennwort", "")
	for _, r := range rs {
		form.Add("dwz[]", r)
	}
	header := http.Header{}
	header.Add("Content-Type", "application/x-www-form-urlencoded")

	err := c.Request(
		http.MethodPost,
		"http://www.isewase.de/dwz/",
		strings.NewReader(form.Encode()),
		nil,
		header,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: ./%s TOURNAMENT_FILE.CSV DWZ YEAR_OF_BIRTH\n", os.Args[0])
		os.Exit(1)
	}

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	p, rs := parsePointsAndRatings(string(content))
	r := requestCalculation(
		os.Args[2],
		os.Args[3],
		p,
		rs,
	)
	fmt.Println(r)
}

func mustParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}
