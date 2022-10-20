package main

import (
	"fmt"
	"testing"
	"time"
)

type person struct {
	name string
}

func (p *person) getName() string {
	fmt.Println(p.name)
	return p.name
}

func TestBlog(t *testing.T) {

	p1 := []*person{{"one"}, {"two"}, {"three"}}
	p2 := []person{{"four"}, {"five"}, {"six"}}

	for _, v := range p1 {
		go v.getName()
	}

	for _, v := range p2 {
		go v.getName()
	}

	time.Sleep(time.Second * 3)

}
