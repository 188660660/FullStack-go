package dao

import (
	"context"
	"testing"
)

func Test_MaxAID(t *testing.T) {
	Convey("MaxAID", t, func() {
		d.MaxAID(context.TODO())
	})
}
