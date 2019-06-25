package vip

import (
	"context"
	"testing"
)

// go test  -test.v -test.run TestServiceManagerInfo
func TestServiceManagerInfo(t *testing.T) {
	Convey("TestServiceManagerInfo", t, func() {
		res, err := s.ManagerInfo(context.TODO())
		t.Logf("res %+v", res)
		So(err, ShouldBeNil)
	})
}
