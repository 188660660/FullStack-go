package feed

import (
	"context"
	"testing"
)

func Test_banners(t *testing.T) {
	Convey("should get banners", t, func() {
		_, _, err := s.banners(context.Background(), 1, 2, 3, "", "", "", "", "", "", "")
		So(err, ShouldBeNil)
	})
}
