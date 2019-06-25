package dao

import (
	"testing"
)

func TestDao_ArchiveAdd(t *testing.T) {
	Convey("TestDao_ArchiveAdd", t, WithDao(func(d *Dao) {
		err := d.NeedImport(123)
		So(err, ShouldBeNil)
	}))
}
