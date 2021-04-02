package main

import (
	"testing"
	"time"
)

func Test_Sleep(t *testing.T) {
	for i := 0; i < 3; i++ {
		Debug("begin", time.Now().Format("2006-01-02_15:04:05"))
		Debug("Do something 1s")
		time.Sleep(time.Second * 1)
		Debug("end", time.Now().Format("2006-01-02_15:04:05"))
		time.Sleep(time.Second * 5)
	}
}

func Test_Tick(t *testing.T) {
	t1 := time.NewTicker(5 * time.Second)(5 * time.Second)
	for {
		select {
		case <-t1.C:
			Debug("begin", time.Now().Format("2006-01-02_15:04:05"))
			Debug("Do something 1s")
			time.Sleep(time.Second * 1)
			Debug("end", time.Now().Format("2006-01-02_15:04:05"))
		}
	}
}
