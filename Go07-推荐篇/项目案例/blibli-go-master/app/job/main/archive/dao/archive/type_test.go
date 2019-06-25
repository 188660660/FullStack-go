package archive

import (
	"context"
	"testing"
)

func Test_TypeMapping(t *testing.T) {
	Convey("TypeMapping", t, func() {
		_, err := d.TypeMapping(context.TODO())
		So(err, ShouldBeNil)
	})
}
