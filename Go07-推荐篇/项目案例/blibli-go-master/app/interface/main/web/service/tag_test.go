package service

import (
	"context"
	"testing"
)

func TestService_TagAids(t *testing.T) {
	Convey("should return without err", t, WithService(func(svf *Service) {
		total, res, err := svf.TagAids(context.Background(), 1, 1, 1)
		So(err, ShouldBeNil)
		So(total, ShouldBeGreaterThan, 0)
		So(len(res), ShouldBeGreaterThan, 0)
	}))

}
