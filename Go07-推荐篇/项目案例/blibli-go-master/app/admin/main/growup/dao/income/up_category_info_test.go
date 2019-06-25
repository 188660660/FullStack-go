package income

import (
	"context"
	"testing"
)

func Test_ListUpInfo(t *testing.T) {
	Convey("ListUpInfo by mids", t, WithMysql(func(d *Dao) {
		var (
			mids = []int64{1, 2}
		)
		_, err := d.ListUpInfo(context.Background(), mids)
		So(err, ShouldBeNil)
	}))
}
