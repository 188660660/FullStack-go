package archive

import (
	"context"
	"testing"
)

func Test_GetFirstPassByAID(t *testing.T) {
	Convey("test archive", t, WithDao(func(d *Dao) {
		id, err := d.GetFirstPassByAID(context.Background(), 10098814)
		So(err, ShouldBeNil)
		So(id, ShouldNotBeNil)
	}))
}
