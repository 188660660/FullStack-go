package manager

import (
	"context"
	"testing"
)

func Test_GetUserRole(t *testing.T) {
	var (
		err error
	)
	Convey("GetUserRole", t, WithDao(func(d *Dao) {
		_, err = d.GetUserRole(context.Background(), 1)
		So(err, ShouldBeNil)
	}))
}
