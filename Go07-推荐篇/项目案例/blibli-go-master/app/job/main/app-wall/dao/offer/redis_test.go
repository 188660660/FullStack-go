package offer

import (
	"testing"
)

func TestPushFail(t *testing.T) {
	Convey("PushFail", t, func() {
		err := d.PushFail(ctx(), "")
		So(err, ShouldBeNil)
	})
}
