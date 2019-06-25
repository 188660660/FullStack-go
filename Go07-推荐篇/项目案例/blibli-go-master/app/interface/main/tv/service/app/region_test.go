package service

import (
	"context"
	"testing"
)

func TestService_Regions(t *testing.T) {
	Convey("test service regions", t, WithService(func(s *Service) {
		res, err := s.Regions(context.Background())
		So(err, ShouldBeNil)
		So(len(res), ShouldBeGreaterThan, 0)
	}))
}
