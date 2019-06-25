package service

import (
	"context"
	"testing"
)

func Test_ArcsRPC(t *testing.T) {
	Convey("ArcsRPC", t, func() {
		var (
			aids = []int64{123, 345}
		)
		res, err := s.ArcsRPC(context.TODO(), aids)
		t.Logf("res:%+v", res)
		t.Logf("err:%+v", err)
		So(res, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})
}
