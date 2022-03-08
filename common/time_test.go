package common

import "testing"

func TestTime(t *testing.T) {
	c := "2021-01-01"
	t.Log(Str2Time(c,DATESTR).Weekday())

}
