package income

import (
	"context"
	"testing"
)

func Test_GetUpWithdraw(t *testing.T) {
	Convey("GetUpWithdraw", t, WithMysql(func(d *Dao) {
		_, err := d.ListUpWithdraw(context.Background(), 0, "", 10)
		So(err, ShouldBeNil)
	}))
}
