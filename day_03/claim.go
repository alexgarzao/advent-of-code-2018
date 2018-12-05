package main

import (
	"regexp"
	"strconv"
)

type claim struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func lineToClaim(s string) claim {
	re := regexp.MustCompile("[,x :#@]")
	result := re.Split(s, -1)
	c := claim{}
	c.id, _ = strconv.Atoi(result[1])
	c.x, _ = strconv.Atoi(result[4])
	c.y, _ = strconv.Atoi(result[5])
	c.w, _ = strconv.Atoi(result[7])
	c.h, _ = strconv.Atoi(result[8])
	return c
}
