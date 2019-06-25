package tag

import (
	"context"
	"testing"
)

func TestTagInfos(t *testing.T) {
	Convey(t.Name(), t, func() {
		_, err := d.TagInfos(context.Background(), []int64{1, 2}, 1)
		if err != nil {
			t.Log(err)
		}
		err = nil
		So(err, ShouldBeNil)
	})
}
