package feed

import (
	"context"
	"testing"
	"time"
)

func Test_Menus(t *testing.T) {
	Convey(t.Name(), t, func() {
		m := s.Menus(context.Background(), 1, 2, time.Now())
		So(m, ShouldNotBeNil)
	})
}
