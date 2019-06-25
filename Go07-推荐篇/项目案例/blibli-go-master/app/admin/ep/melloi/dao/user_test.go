package dao

import (
	"testing"
)

func Test_User(t *testing.T) {
	Convey("test QueryUser", t, func() {
		_, err := d.QueryUserByUserName("hujianping")
		So(err, ShouldBeNil)
	})
}
