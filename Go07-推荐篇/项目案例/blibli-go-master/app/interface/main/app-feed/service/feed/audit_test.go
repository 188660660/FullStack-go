package feed

import (
	"context"
	"testing"
)

func Test_Audit(t *testing.T) {
	Convey("should get audit", t, func() {
		_, err := s.Audit(context.Background(), "", 1, 2)
		So(err, ShouldBeNil)
	})
}
