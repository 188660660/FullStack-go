package income

import (
	"context"
	"testing"
)

func Test_GetAvBlackListByAvIds(t *testing.T) {
	Convey("GetAvBlackListByAvIds", t, WithService(func(s *Service) {
		_, err := s.GetAvBlackListByAvIds(context.Background(), nil, 0)
		So(err, ShouldBeNil)
	}))
}
