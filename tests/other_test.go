package test

import (
	"testing"
	"time"
	"math/rand"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func Test_String(t *testing.T) {
	t.Log(" start . .  . ")
	for i := 0; i < 100; i++ {
		t.Log(r.Intn(2)+1)
	}
}
