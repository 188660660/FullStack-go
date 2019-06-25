package service

import (
	"context"
	"testing"
)

func TestTrackVideoup(t *testing.T) {
	Convey("", t, func() {
		err := svr.trackVideoup(context.TODO(), 10114205)
		So(err, ShouldBeNil)
	})
}
