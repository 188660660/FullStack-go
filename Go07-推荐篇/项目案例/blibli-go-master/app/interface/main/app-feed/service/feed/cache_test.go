package feed

import (
	"context"
	"testing"
)

func Test_indexCache(t *testing.T) {
	Convey("should get indexCache", t, func() {
		_, err := s.indexCache(context.Background(), 1, 2)
		So(err, ShouldBeNil)
	})
}
