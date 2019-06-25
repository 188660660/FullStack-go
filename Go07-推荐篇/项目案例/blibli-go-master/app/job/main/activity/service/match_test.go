package service

import (
	"context"
	"testing"
)

func TestService_FinishMatch(t *testing.T) {
	Convey("test finish match", t, WithService(func(s *Service) {
		moID := int64(3)
		err := s.FinishMatch(context.Background(), moID)
		So(err, ShouldBeNil)
	}))
}
