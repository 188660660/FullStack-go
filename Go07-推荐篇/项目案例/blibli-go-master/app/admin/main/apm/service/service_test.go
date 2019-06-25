package service

import (
	"context"
	"testing"
)

func TestService(t *testing.T) {
	Convey("service", t, func() {
		t.Log("service test")
	})
}

func TestPing(t *testing.T) {
	Convey("ping", t, func() {
		var s = &Service{}
		s.Ping(context.Background())
	})
}
