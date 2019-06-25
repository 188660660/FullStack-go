package job

import (
	"context"
	"testing"
)

func TestDao_Job(t *testing.T) {
	Convey("test job", t, WithDao(func(d *Dao) {
		data, err := d.Jobs(context.TODO())
		So(err, ShouldBeNil)
		Printf("%+v", data)
	}))
}
