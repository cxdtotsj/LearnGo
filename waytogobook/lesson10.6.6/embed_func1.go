package main

import "fmt"

type Log struct {
	msg string
}

type Customer struct {
	name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}

func main() {
	// c := new(Customer)
	// c.name = "XDchen"
	// c.log = new(Log)
	// c.log.msg = "YES SUCCESS"
	c := &Customer{"XDchen", &Log{"YES SUCCESS"}}
	c.log.Add("ADD")
	fmt.Println(c.Log())
}
