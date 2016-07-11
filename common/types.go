package common

import (
	"log"
	"strconv"
	"strings"
)

type Record struct {
	Date     Date
	Open     float32
	Close    float32
	High     float32
	Low      float32
	AdjClose float32
	Volume   int
}

type Symbol string

type Date string

func Str2Date(s string) Date {
	if len(s) != 10 {
		log.Panicf("%s is not a valid date", s)
	}

	return Date(s)
}

func (date Date) GetYearMonthDay() (year int, month int, day int) {
	var err error
	parts := strings.Split(string(date), "-")
	if year, err = strconv.Atoi(parts[0]); err != nil {
		log.Panicf("%s is not a valid date", date)
	}

	if month, err = strconv.Atoi(parts[1]); err != nil {
		log.Panicf("%s is not a valid date", date)
	}

	if day, err = strconv.Atoi(parts[2]); err != nil {
		log.Panicf("%s is not a valid date", date)
	}

	return
}

func (date Date) Year() int {
	y, _, _ := date.GetYearMonthDay()
	return y
}

func (date Date) Month() int {
	_, m, _ := date.GetYearMonthDay()
	return m
}

func (date Date) Day() int {
	_, _, d := date.GetYearMonthDay()
	return d
}
