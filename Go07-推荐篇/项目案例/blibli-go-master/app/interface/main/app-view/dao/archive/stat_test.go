package archive

import (
	"context"
	"testing"
)

func TestDao_Stat(t *testing.T) {
	Convey("Stat", t, func() {
		d.Stat(context.TODO(), 2)
	})
}
