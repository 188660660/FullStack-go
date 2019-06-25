package feed

import (
	"context"
	"testing"
)

func Test_Hots(t *testing.T) {
	Convey("Hots", t, func() {
		d.Hots(context.TODO())
	})
}

func Test_UpRcmdCache(t *testing.T) {
	Convey("UpRcmdCache", t, func() {
		d.UpRcmdCache(context.TODO(), nil)
	})
}
