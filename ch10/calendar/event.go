package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	//如果title有超过30个字符，返回错误
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
