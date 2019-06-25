package monitor

import (
	"context"
	"testing"
)

func Test_Send(t *testing.T) {
	Convey("Send", t, func() {
		err := d.Send(context.TODO(), "报警短信test")
		So(err, ShouldBeNil)
	})
}
