package archive

import (
	"context"
	"testing"
)

func Test_TypeMapping(t *testing.T) {
	Convey("test archive", t, WithDao(func(d *Dao) {
		_, err := d.TypeMapping(context.Background())
		So(err, ShouldBeNil)
	}))
}
