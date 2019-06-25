package income

import (
	"context"
	"testing"
)

func Test_GetUpInfoByAIDs(t *testing.T) {
	Convey("GetUpInfoByAIDs", t, WithService(func(s *Service) {
		_, err := s.GetUpInfoByAIDs(context.Background(), nil)
		So(err, ShouldBeNil)
	}))
}
