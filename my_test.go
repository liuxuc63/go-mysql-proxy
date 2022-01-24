package main

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	println(">> UnixMilli=", now.UnixMilli())
	println(">> UnixMicro=", now.UnixMicro())
	println(">> UnixNano=", now.UnixNano())
	println(">> Nanosecond=", now.Nanosecond())

	dr1, _ := time.ParseDuration("1s")
	now2 := now.Add(dr1)
	println(">> UnixMilli=", now.UnixMilli())
	println(">> 15m=", now2.UnixMilli()-now.UnixMilli())
}

//1642995823146
//1642995842375
