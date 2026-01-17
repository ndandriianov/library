package library

import "time"

type Book struct {
	Title         string
	Author        string
	NumberOfPages int
	IsFinished    bool
	AddTime       time.Time
	FinishTime    time.Time
}
