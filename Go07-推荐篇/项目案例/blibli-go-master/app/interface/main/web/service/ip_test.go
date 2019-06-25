package service

import (
	"context"
	"testing"
)

func TestService_IPZone(t *testing.T) {
	Convey("test ip IPZone", t, WithService(func(s *Service) {
		res, err := s.IPZone(context.Background())
		So(err, ShouldBeNil)
		So(res, ShouldNotBeNil)
	}))
}
