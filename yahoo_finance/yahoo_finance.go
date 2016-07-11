package yahoo_finance

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	. "github.com/xiaonanln/yahoo-finance-prefetch/common"
)

func Fetch(symbol Symbol, startDate Date, stopDate Date) ([]*Record, error) {
	c, a, b := startDate.GetYearMonthDay()
	f, d, e := stopDate.GetYearMonthDay()

	a = a - 1
	d = d - 1

	url := fmt.Sprintf("http://real-chart.finance.yahoo.com/table.csv?s=%s&d=%d&e=%d&f=%d&g=d&a=%d&b=%d&c=%d&ignore=.csv", symbol, d, e, f, a, b, c)
	log.Printf("GET %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	body := string(data)
	lines := strings.Split(body, "\n")
	if lines[0] != "Date,Open,High,Low,Close,Volume,Adj Close" {
		log.Panicf("invalid yahoo-finance output: %s", lines[0])
	}

	records := make([]*Record, 0, len(lines)-1)
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		record := interpretLine(line)
		records = append(records, record)
	}

	return nil, nil
}

func interpretLine(line string) *Record {
	sp := strings.Split(line, ",")

	// item['Date'], item['Open'], item['High'], item['Low'], item['Close'], item['Volume'], item['AdjClose'] = line
	r := Record{
		Date:     Str2Date(sp[0]),
		Open:     atof(sp[1]),
		High:     atof(sp[2]),
		Low:      atof(sp[3]),
		Close:    atof(sp[4]),
		Volume:   atoi(sp[5]),
		AdjClose: atof(sp[6]),
	}
	log.Printf("Converted %v to %v\n", sp, r)
	return &r
}

func atof(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Panicf("%s is not a valid float", s)
	}
	return float32(f)
}

func atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("%s is not a valid int", s)
	}
	return v
}
