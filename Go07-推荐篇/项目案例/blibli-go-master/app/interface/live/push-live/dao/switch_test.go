package dao

import (
	"context"
	"testing"
)

func Test_switch(t *testing.T) {
	initd()
	Convey("Parse Json To Struct", t, func() {
		target := int64(27515316)
		fs, err := d.GetFansBySwitch(context.TODO(), target)
		t.Logf("the result included(%v) err(%v)", fs, err)

		So(err, ShouldEqual, nil)
	})
}
